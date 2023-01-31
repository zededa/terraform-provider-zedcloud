package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func DriveModel(d *schema.ResourceData) *models.Drive {
	cleartext, _ := d.Get("cleartext").(bool)
	drvtype, _ := d.Get("drvtype").(string)
	ignorepurge, _ := d.Get("ignorepurge").(bool)
	imagename, _ := d.Get("imagename").(string)
	imvolname, _ := d.Get("imvolname").(string)
	maxsize, _ := d.Get("maxsize").(uint64)
	mountpath, _ := d.Get("mountpath").(string)
	mvolname, _ := d.Get("mvolname").(string)
	preserve, _ := d.Get("preserve").(bool)
	readonly, _ := d.Get("readonly").(bool)
	target, _ := d.Get("target").(string)
	volumelabel, _ := d.Get("volumelabel").(string)
	return &models.Drive{
		Cleartext:   cleartext,
		Drvtype:     &drvtype, // string true false false
		Ignorepurge: ignorepurge,
		Imagename:   &imagename, // string true false false
		Imvolname:   imvolname,
		Maxsize:     &maxsize, // uint64 true false false
		Mountpath:   mountpath,
		Mvolname:    mvolname,
		Preserve:    &preserve, // bool true false false
		Readonly:    &readonly, // bool true false false
		Target:      &target,   // string true false false
		Volumelabel: volumelabel,
	}
}

func DriveModelFromMap(m map[string]interface{}) *models.Drive {
	cleartext := m["cleartext"].(bool)
	drvtype := m["drvtype"].(string)
	ignorepurge := m["ignorepurge"].(bool)
	imagename := m["imagename"].(string)
	imvolname := m["imvolname"].(string)
	maxsize := m["maxsize"].(uint64)
	mountpath := m["mountpath"].(string)
	mvolname := m["mvolname"].(string)
	preserve := m["preserve"].(bool)
	readonly := m["readonly"].(bool)
	target := m["target"].(string)
	volumelabel := m["volumelabel"].(string)
	return &models.Drive{
		Cleartext:   cleartext,
		Drvtype:     &drvtype,
		Ignorepurge: ignorepurge,
		Imagename:   &imagename,
		Imvolname:   imvolname,
		Maxsize:     &maxsize,
		Mountpath:   mountpath,
		Mvolname:    mvolname,
		Preserve:    &preserve,
		Readonly:    &readonly,
		Target:      &target,
		Volumelabel: volumelabel,
	}
}

// Update the underlying Drive resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDriveResourceData(d *schema.ResourceData, m *models.Drive) {
	d.Set("cleartext", m.Cleartext)
	d.Set("drvtype", m.Drvtype)
	d.Set("ignorepurge", m.Ignorepurge)
	d.Set("imagename", m.Imagename)
	d.Set("imvolname", m.Imvolname)
	d.Set("maxsize", m.Maxsize)
	d.Set("mountpath", m.Mountpath)
	d.Set("mvolname", m.Mvolname)
	d.Set("preserve", m.Preserve)
	d.Set("readonly", m.Readonly)
	d.Set("target", m.Target)
	d.Set("volumelabel", m.Volumelabel)
}

// Iterate through and update the Drive resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDriveSubResourceData(m []*models.Drive) (d []*map[string]interface{}) {
	for _, DriveModel := range m {
		if DriveModel != nil {
			properties := make(map[string]interface{})
			properties["cleartext"] = DriveModel.Cleartext
			properties["drvtype"] = DriveModel.Drvtype
			properties["ignorepurge"] = DriveModel.Ignorepurge
			properties["imagename"] = DriveModel.Imagename
			properties["imvolname"] = DriveModel.Imvolname
			properties["maxsize"] = DriveModel.Maxsize
			properties["mountpath"] = DriveModel.Mountpath
			properties["mvolname"] = DriveModel.Mvolname
			properties["preserve"] = DriveModel.Preserve
			properties["readonly"] = DriveModel.Readonly
			properties["target"] = DriveModel.Target
			properties["volumelabel"] = DriveModel.Volumelabel
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the Drive resource defined in the Terraform configuration
func DriveSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cleartext": {
			Description: `If set, don't encrypt the drive. Store in cleartext`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"drvtype": {
			Description: `Type of Drive. Valid Values: UNSPECIFIED, CDROM, HDD, NET, HDD_EMPTY. HDD_EMPTY - is to allocate the empty disk of maxsizebytes specified`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"ignorepurge": {
			Description: `don't purge this drive as part of purge command for mutable volumes`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"imagename": {
			Description: `Name of Image Object used for the drive.`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"imvolname": {
			Description: `immutable Volume for this drive. Only one of imvolname and mvolname must be specified.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"maxsize": {
			Description: `Drive maximum size`,
			Type:        schema.TypeInt,
			Required:    true,
		},

		"mountpath": {
			Description: `Mount Path for the drive in the App Instance`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"mvolname": {
			Description: `mutable Volume for this drive. Only one of imvolname and mvolname must be specified.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"preserve": {
			Description: `Preserve the drive even when all app instances using it are deleted`,
			Type:        schema.TypeBool,
			Required:    true,
		},

		"readonly": {
			Description: `Read only flag. If set, drive is mounted as readonly by app instance.`,
			Type:        schema.TypeBool,
			Required:    true,
		},

		"target": {
			Description: `type of target. Valid values: Disk, Kernel, Initrd, RamDisk`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"volumelabel": {
			Description: `User defined volume to use for this drive `,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the Drive resource
func GetDrivePropertyFields() (t []string) {
	return []string{
		"cleartext",
		"drvtype",
		"ignorepurge",
		"imagename",
		"imvolname",
		"maxsize",
		"mountpath",
		"mvolname",
		"preserve",
		"readonly",
		"target",
		"volumelabel",
	}
}
