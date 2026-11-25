package resources

// ENG-3000 / UE-203: interface ACLs on zedcloud_application_instance.
//
// On projects with interface ordering enabled, the controller (zedcloud
// #8548) stores an interface's ACL list as sent; an empty list takes the
// app bundle's ACLs. On projects with ordering disabled it ignores ACL
// changes after create and appends the bundle ACLs at create. The interface
// diff suppressor lets ACL changes through (schemas/schema_helpers.go); this
// file clears them again where the controller would ignore them, and warns
// when the controller stored something other than what was sent.
//
// Design: docs/design/eng-3000-app-instance-acl-defaults.md.

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	project "github.com/zededa/terraform-provider-zedcloud/v2/client/projects"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

const (
	interfaceOrderingOn  = "enabled"
	interfaceOrderingOff = "disabled"
)

// interfaceOrderingLookups caches each project's interface-ordering setting
// per provider process, so a plan over many instances in one project reads
// the project once.
var interfaceOrderingLookups = &projectCache{entries: map[string]*projectEntry{}}

// projectInterfaceOrdering reports whether the project has interface ordering
// enabled. A variable so unit tests can stub it.
var projectInterfaceOrdering = func(ctx context.Context, m interface{}, projectID string) (bool, error) {
	client, ok := m.(*api_client.ZedcloudAPI)
	if !ok || client == nil {
		return false, errors.New("no zedcloud client")
	}
	v, err := interfaceOrderingLookups.get(ctx, projectID, func() (string, error) {
		params := project.GetByIDParams().WithContext(ctx)
		params.ID = projectID
		resp, err := client.Project.GetByID(params, nil)
		if err != nil {
			return "", err
		}
		p := resp.GetPayload()
		if p != nil && p.TagLevelSettings != nil && p.TagLevelSettings.InterfaceOrdering != nil &&
			*p.TagLevelSettings.InterfaceOrdering == models.InterfaceOrderingINTERFACEORDERINGENABLED {
			return interfaceOrderingOn, nil
		}
		return interfaceOrderingOff, nil
	})
	if err != nil {
		return false, err
	}
	return v == interfaceOrderingOn, nil
}

// gateInterfaceACLDiffs clears planned ACL changes of existing app instances
// on projects with interface ordering disabled, where the controller ignores
// them. On those projects create also appends the bundle ACLs, so the stored
// list can differ from the configuration for good; without this, every plan
// would show that difference. The project is only looked up when an ACL
// change is planned. If it can't be determined, the change is kept: applying
// it is harmless where the controller ignores it.
func gateInterfaceACLDiffs(ctx context.Context, d *schema.ResourceDiff, m interface{}) error {
	if d.Id() == "" {
		return nil
	}
	// Same rule as the diff suppressor: only interfaces at the same index
	// under the same name can carry a planned ACL change. HasChange alone
	// would also report the suppressed changes of added or renamed ones.
	o, n := d.GetChange("interfaces")
	oldIntfs, _ := o.([]interface{})
	newIntfs, _ := n.([]interface{})
	if len(oldIntfs) != len(newIntfs) {
		return nil
	}
	var changed []int
	for i := range newIntfs {
		if interfaceName(oldIntfs[i]) == "" || interfaceName(oldIntfs[i]) != interfaceName(newIntfs[i]) {
			continue
		}
		if d.HasChange(fmt.Sprintf("interfaces.%d.acls", i)) {
			changed = append(changed, i)
		}
	}
	if len(changed) == 0 {
		return nil
	}
	projectID, _ := d.Get("project_id").(string)
	if projectID == "" {
		log.Printf("[WARN] edge application instance %q: no project_id, keeping the planned ACL changes", d.Get("name"))
		return nil
	}
	enabled, err := projectInterfaceOrdering(ctx, m, projectID)
	if err != nil {
		log.Printf("[WARN] edge application instance %q: reading interface ordering of project %s: %s; "+
			"keeping the planned ACL changes", d.Get("name"), projectID, err)
		return nil
	}
	if enabled {
		return nil
	}
	for _, i := range changed {
		key := fmt.Sprintf("interfaces.%d.acls", i)
		log.Printf("[DEBUG] edge application instance %q: interface ordering is disabled in project %s, "+
			"ignoring the change to %s", d.Get("name"), projectID, key)
		if err := d.Clear(key); err != nil {
			return err
		}
	}
	return nil
}

// aclReadBackDiags warns when the controller stored an interface ACL list
// other than the one the provider sent, on a project with interface ordering
// enabled. That happens with controllers older than zedcloud #8548, which
// append the bundle ACLs behind the configured ones; the next plan would show
// the same ACL diff again after every apply. sent is d.Get("interfaces")
// taken before the write; d holds the read-back state.
func aclReadBackDiags(ctx context.Context, d *schema.ResourceData, m interface{}, sent []interface{}) diag.Diagnostics {
	got, _ := d.Get("interfaces").([]interface{})
	var mismatched []string
	for i := 0; i < len(sent) && i < len(got); i++ {
		s, _ := sent[i].(map[string]interface{})
		g, _ := got[i].(map[string]interface{})
		if s == nil || g == nil || s["intfname"] != g["intfname"] {
			continue
		}
		want := aclRules(s["acls"])
		if len(want) == 0 {
			continue // omitted: the controller applied the bundle defaults
		}
		if !reflect.DeepEqual(want, aclRules(g["acls"])) {
			name, _ := s["intfname"].(string)
			mismatched = append(mismatched, name)
		}
	}
	if len(mismatched) == 0 {
		return nil
	}
	projectID, _ := d.Get("project_id").(string)
	if projectID == "" {
		return nil
	}
	enabled, err := projectInterfaceOrdering(ctx, m, projectID)
	if err != nil || !enabled {
		return nil
	}
	return diag.Diagnostics{{
		Severity: diag.Warning,
		Summary:  "Interface ACLs were not stored as configured",
		Detail: fmt.Sprintf("Edge application instance %q: the controller stored different ACLs than configured "+
			"for interface(s) %s. Controllers without zedcloud #8548 (ENG-3000) add the app bundle's ACLs behind "+
			"the configured ones, so the next plan will show this ACL change again.",
			d.Get("name"), strings.Join(mismatched, ", ")),
	}}
}

func interfaceName(v interface{}) string {
	m, _ := v.(map[string]interface{})
	name, _ := m["intfname"].(string)
	return name
}

// aclRules returns the rule content of an acls list as read from
// ResourceData: matches and actions only. The id is the controller's, and
// the name is not part of what makes two rules equal.
func aclRules(v interface{}) []interface{} {
	list, _ := v.([]interface{})
	out := make([]interface{}, 0, len(list))
	for _, e := range list {
		ace, _ := e.(map[string]interface{})
		if ace == nil {
			continue
		}
		out = append(out, map[string]interface{}{"matches": ace["matches"], "actions": ace["actions"]})
	}
	return out
}
