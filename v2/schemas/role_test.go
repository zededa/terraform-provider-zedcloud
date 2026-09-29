package schemas

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func roleData(attrs map[string]string) *schema.ResourceData {
	r := &schema.Resource{Schema: RoleSchema()}
	return r.Data(&terraform.InstanceState{
		ID:         "CCGFABAEqnH4je5PHZTXSmHOs-ZE",
		Attributes: attrs,
	})
}

// NFR-165 §4.2. The API documents the role name as unchangeable once created,
// so a rename must plan a replacement instead of an update the server rejects.
func TestRoleSchema_ImmutableName(t *testing.T) {
	s := RoleSchema()

	if !s["name"].ForceNew {
		t.Error("name must be ForceNew: the API cannot rename a role in place")
	}
	if s["title"].ForceNew {
		t.Error("title must not be ForceNew: the API allows it to change at any time")
	}
}

// NFR-165 §3. What Read writes into state after an import must survive being
// turned back into a request body, scopes included — a role's scopes are its
// entire reason to exist.
func TestSetRoleResourceData_RoundTrip(t *testing.T) {
	name := "readonly-operators"
	title := "Read-only operators"

	from := &models.Role{
		ID:          "CCGFABAEqnH4je5PHZTXSmHOs-ZE",
		Name:        &name,
		Title:       &title,
		Description: "Operators who may look but not touch",
		Type:        models.NewUserRole(models.UserRoleUSERROLEUSERDEFINED),
		State:       models.NewRoleState(models.RoleStateROLESTATEACTIVE),
		ProjectTags: map[string]string{"env": "prod"},
		Scopes: []*models.Scope{
			{
				AccessApp:         "PermissionAccessRead",
				AccessAppInstance: "PermissionAccessRead",
				AccessDevice:      "PermissionAccessRead",
				AccessEdgeApp:     "PermissionAccessRead",
				AccessEnterprise:  "PermissionAccessRead",
				AccessStorage:     "PermissionAccessRead",
				AccessUser:        "PermissionAccessNone",
				EnterpriseFilter:  []string{"srAll"},
				ProjectFilter:     []string{"srAll"},
			},
		},
	}

	d := roleData(map[string]string{})
	SetRoleResourceData(d, from)

	back := RoleModel(d)

	if back.Name == nil || *back.Name != name {
		t.Errorf("Name = %v, want %q", back.Name, name)
	}
	if back.Title == nil || *back.Title != title {
		t.Errorf("Title = %v, want %q", back.Title, title)
	}
	if back.Description != from.Description {
		t.Errorf("Description = %q, want %q", back.Description, from.Description)
	}
	if back.Type == nil || *back.Type != *from.Type {
		t.Errorf("Type = %v, want %v", back.Type, *from.Type)
	}
	if back.State == nil || *back.State != *from.State {
		t.Errorf("State = %v, want %v", back.State, *from.State)
	}
	if back.ProjectTags["env"] != "prod" {
		t.Errorf("ProjectTags = %v, want env=prod", back.ProjectTags)
	}

	if len(back.Scopes) != 1 {
		t.Fatalf("Scopes = %v, want exactly 1", back.Scopes)
	}
	got := back.Scopes[0]
	want := from.Scopes[0]
	if got.AccessApp != want.AccessApp ||
		got.AccessAppInstance != want.AccessAppInstance ||
		got.AccessDevice != want.AccessDevice ||
		got.AccessEdgeApp != want.AccessEdgeApp ||
		got.AccessEnterprise != want.AccessEnterprise ||
		got.AccessStorage != want.AccessStorage ||
		got.AccessUser != want.AccessUser {
		t.Errorf("scope permissions did not survive the round trip:\n got %+v\nwant %+v", got, want)
	}
	if len(got.EnterpriseFilter) != 1 || got.EnterpriseFilter[0] != "srAll" {
		t.Errorf("EnterpriseFilter = %v, want [srAll]", got.EnterpriseFilter)
	}
	if len(got.ProjectFilter) != 1 || got.ProjectFilter[0] != "srAll" {
		t.Errorf("ProjectFilter = %v, want [srAll]", got.ProjectFilter)
	}
}
