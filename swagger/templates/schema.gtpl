package schemas

{{- $operationGroup := pascalize .Name -}}
{{- $hasID := false -}}

{{/* Some property fields can be characterized as more "complex" objects whose data must be retrieved with specialized helper functions */}}
{{- $isReadOnlyModel := true -}}
{{- $needsUtils := false -}}
{{- range .Properties -}}
	{{- if not .ReadOnly -}} {{- $isReadOnlyModel = false -}} {{- end -}}
	{{- if eq .Name "customProperties" -}} {{- $needsUtils = true -}} {{- end -}}
	{{- if eq .Name "id" -}} {{- $hasID = true -}} {{- end -}}
{{- end }}

import (
	{{- if $needsUtils }}
	"tf-provider-example/logicmonitor/utils"
	{{- end }}
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
    "github.com/zededa/terraform-provider-zedcloud/models"
    "github.com/go-openapi/strfmt"
)

func {{ $operationGroup }}Model(d *schema.ResourceData) *models.{{ $operationGroup }} {
	{{- range .Properties }}
		{{- if or (not .ReadOnly) (and (eq .Name "id") (not $isReadOnlyModel)) }}
			{{- if .IsBase64 }}
	{{ varname .Name }}, _ := d.Get("{{ snakize .Name }}").(strfmt.Base64)
			{{- else if and .IsInterface (not .IsAnonymous) }}
	{{ varname .Name }}, _ := d.Get("{{ snakize .Name }}").(models.{{ .GoType }}) // {{.GoType}}
			{{- else if .IsMap }}
			  {{- if stringContains .GoType "map[string]string" }}
	{{ varname .Name }} := map[string]string{}
	{{ varname .Name }}Interface, {{ varname .Name }}IsSet := d.GetOk("{{ varname .Name }}")
	if {{ varname .Name }}IsSet {
		{{ varname .Name }}Map := {{ varname .Name }}Interface.(map[string]interface{})
		for k, v := range {{ varname .Name }}Map {
			if v == nil {
				continue
			}
			{{ varname .Name }}[k] = v.(string)
		}
	}
              {{else}}
	{{ varname .Name }}, _ := d.Get("{{ snakize .Name }}").({{ .GoType }}) // {{.GoType}}
              {{end}}
			{{- else if .IsComplexObject }}
	var {{ varname .Name }} *models.{{ pascalize .GoType }} // {{ .GoType }}
	{{ .Name }}Interface, {{ .Name }}IsSet := d.GetOk("{{ snakize .Name }}")
	if {{ .Name }}IsSet && {{ .Name }}Interface != nil {
		{{ .Name }}Map := {{ .Name }}Interface.([]interface{})
		if len({{ .Name }}Map) > 0{
		    {{ varname .Name }} = {{ .GoType }}ModelFromMap({{ .Name }}Map[0].(map[string]interface{}))
		}
	}
			{{- else if stringContains .Name "Properties" }}
	{{ varname .Name }} := utils.GetPropertiesFromResource(d, "{{ snakize .Name }}")
	{{ varname .Name }} := d.Get("{{ snakize .Name }}")
			{{- else if eq .GoType "strfmt.DateTime" }}
	{{ varname .Name }}, _ := d.Get("{{ snakize .Name }}").(strfmt.DateTime)
			{{- else if eq .GoType "interface{}" }}
	{{ varname .Name }} := d.Get("{{ snakize .Name }}")
			{{- else if eq .GoType "[]float64" }}
	{{ varname .Name }}, _ := d.Get("{{ snakize .Name }}").([]float64)
			{{- else if or (eq .GoType "int32") ( eq .GoType "int64") }}
	{{ varname .Name }}Int, _ := d.Get("{{ snakize .Name }}").(int)
	{{ varname .Name }} := {{ .GoType }}({{ varname .Name }}Int)
			{{- else if and (not (eq .GoType "string")) (not (eq .GoType "bool")) (not (eq .GoType "int")) (not (eq .GoType "float32")) (not (eq .GoType "float64")) (not (eq .GoType "uint32")) (not (eq .GoType "uint64")) }}
				{{- if hasPrefix .GoType "[]*" }}
	var {{ varname .Name }} []*models.{{ pascalize .GoType }} // {{ .GoType }}
	{{ .Name }}Interface, {{ .Name }}IsSet := d.GetOk("{{ snakize .Name }}")
	if {{ .Name }}IsSet {
		var items []interface{}
		if listItems, isList := {{ .Name }}Interface.([]interface{}); isList {
			items = listItems
		} else {
			items = {{ .Name }}Interface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := {{ pascalize .GoType }}ModelFromMap(v.(map[string]interface{}))
			{{ varname .Name }} = append({{ varname .Name }}, m)
		}
	}
				{{- else if or (eq .GoType "[]string") (eq .GoType "[]strfmt.Base64") }}
	var {{ varname .Name }} []string
	{{ varname .Name }}Interface, {{ varname .Name }}IsSet := d.GetOk("{{ varname .Name }}")
	if {{ varname .Name }}IsSet {
		var items []interface{}
		if listItems, isList := {{ .Name }}Interface.([]interface{}); isList {
			items = listItems
		} else {
			items = {{ .Name }}Interface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			{{ varname .Name }} = append({{ varname .Name }}, v.(string))
		}
	}
				{{- else if hasPrefix .GoType "[]" }}
	var {{ varname .Name }} []models.{{ pascalize .GoType }} // {{ .GoType }}
	{{ .Name }}Interface, {{ .Name }}IsSet := d.GetOk("{{ snakize .Name }}")
	if {{ .Name }}IsSet {
		var items []interface{}
		if listItems, isList := {{ .Name }}Interface.([]interface{}); isList {
			items = listItems
		} else {
			items = {{ .Name }}Interface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := {{ pascalize .GoType }}ModelFromMap(v.(map[string]interface{}))
			{{ varname .Name }} = append({{ varname .Name }}, *m)
		}
	}
				{{- else if .IsAliased }}
	var {{ varname .Name }} *models.{{ pascalize .GoType }} // {{ .GoType }}
	{{ .Name }}Interface, {{ .Name }}IsSet := d.GetOk("{{ snakize .Name }}")
	if {{ .Name }}IsSet {
	    {{ .Name }}Model := {{ .Name }}Interface.(string)
	    {{ varname .Name }} = models.New{{.GoType}}(models.{{.GoType}}({{ .Name }}Model))
	}
				{{- else }}
	{{ varname .Name }}, _ := d.Get("{{ snakize .Name }}").(models.{{ pascalize .GoType }}) // {{ .GoType }}
				{{- end }}
			{{- else }}
	{{ varname .Name }}, _ := d.Get("{{ snakize .Name }}").({{ .GoType }})
			{{- end }}
		{{- end }}
	{{- end }}

	{{- if not .GenSchema.IsComplexObject }}
		{{- if .GenSchema.IsAliased }}
   {{ varname .Name }}, _ := d.Get("{{ snakize .Name }}").(models.{{ $operationGroup }})
	return &{{ varname .Name }}
		{{- else }}
   {{ varname .Name }}, _ := d.Get("{{ snakize .Name }}").(models.{{ $operationGroup }})
	return {{ varname .Name }}
		{{- end }}
	{{- else }}
	return &models.{{ $operationGroup }} {
		{{- if not $isReadOnlyModel }}
		{{- range .Properties }}
			{{- if or (not .ReadOnly) (eq .Name "id") }}
				{{- if and .IsNullable (and (not .IsComplexObject) (not .IsAliased)) }}
		{{ pascalize .Name }}: &{{ varname .Name }}, // {{ .GoType }}
				{{- else if and (and (not .IsMap) .IsNullable (not .IsSuperAlias) (not .IsAliased)) (not .IsComplexObject) }}
		{{ pascalize .Name }}: &{{ varname .Name }},
				{{- else }}
		{{ pascalize .Name }}: {{ varname .Name }},
				{{- end }}
			{{- end }}
		{{- end }}
		{{- end }}
	}
	{{- end }}
}


func {{ $operationGroup }}ModelFromMap(m map[string]interface{}) *models.{{ $operationGroup }} {
	{{- range .Properties }}
		{{- if or (not .ReadOnly) (and (eq .Name "id") (not $isReadOnlyModel)) }}
			{{- if .IsBase64 }}
	{{ varname .Name }} := m["{{ snakize .Name }}"].(strfmt.Base64)
			{{- else if and .IsInterface (not .IsAnonymous)}}
	{{ varname .Name }} := m["{{ snakize .Name }}"].(models.{{ .GoType }}) // {{.GoType}}
			{{- else if .IsMap }}
			  {{- if stringContains .GoType "map[string]string" }}
	{{ varname .Name }} := map[string]string{}
	{{ varname .Name }}Interface, {{ varname .Name }}IsSet := m["{{ snakize .Name }}"]
	if {{ varname .Name }}IsSet {
		{{ varname .Name }}Map := {{ varname .Name }}Interface.(map[string]interface{})
		for k, v := range {{ varname .Name }}Map {
			if v == nil {
				continue
			}
			{{ varname .Name }}[k] = v.(string)
		}
	}
              {{else}}
	{{ varname .Name }} := m["{{ snakize .Name }}"].({{ .GoType }})
              {{end}}
			{{- else if .IsComplexObject }}
	var {{ varname .Name }} *models.{{ pascalize .GoType }} // {{ .GoType }}
	{{ .Name }}Interface, {{ .Name }}IsSet := m["{{ snakize .Name }}"]
	if {{ .Name }}IsSet && {{ .Name }}Interface != nil {
		{{ .Name }}Map := {{ .Name }}Interface.([]interface{})
		if len({{ .Name }}Map) > 0{
		    {{ varname .Name }} = {{ .GoType }}ModelFromMap({{ .Name }}Map[0].(map[string]interface{}))
		}
	}
			// {{- else if stringContains .Name "Properties" }}
	// {{ varname .Name }} := utils.GetPropertiesFromResource(d, "{{ snakize .Name }}")
			{{- else if eq .GoType "interface{}" }}
	{{ varname .Name }} := m["{{ snakize .Name }}"]
			{{- else if eq .GoType "[]float64" }}
	{{ varname .Name }} := m["{{ snakize .Name }}"].([]float64)
			{{- else if eq .GoType "strfmt.DateTime" }}
	{{ varname .Name }} := m["{{ snakize .Name }}"].(strfmt.DateTime)
			{{- else if or (eq .GoType "int32") ( eq .GoType "int64") }}
	{{ varname .Name }} := {{ .GoType }}(m["{{ snakize .Name }}"].(int)) // {{ .GoType }}
			{{- else if and (not (eq .GoType "string")) (not (eq .GoType "bool")) (not (eq .GoType "int")) (not (eq .GoType "float32")) (not (eq .GoType "float64")) (not (eq .GoType "uint32")) (not (eq .GoType "uint64")) }}
				{{- if hasPrefix .GoType "[]*" }}
	var {{ varname .Name }} []*models.{{ pascalize .GoType }} // {{ .GoType }}
	{{ .Name }}Interface, {{ .Name }}IsSet := m["{{ snakize .Name }}"]
	if {{ .Name }}IsSet {
		var items []interface{}
		if listItems, isList := {{ .Name }}Interface.([]interface{}); isList {
			items = listItems
		} else {
			items = {{ .Name }}Interface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := {{ pascalize .GoType }}ModelFromMap(v.(map[string]interface{}))
			{{ varname .Name }} = append({{ varname .Name }}, m)
		}
	}
				{{- else if or (hasPrefix .GoType "[]string") (hasPrefix .GoType "[]strfmt.Base64") }}
	var {{ varname .Name }} []string
	{{ varname .Name }}Interface, {{ varname .Name }}IsSet := m["{{ varname .Name }}"]
	if {{ varname .Name }}IsSet {
		var items []interface{}
		if listItems, isList := {{ .Name }}Interface.([]interface{}); isList {
			items = listItems
		} else {
			items = {{ .Name }}Interface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			{{ varname .Name }} = append({{ varname .Name }}, v.(string))
		}
	}
				{{- else if hasPrefix .GoType "[]" }}
	var {{ varname .Name }} []models.{{ pascalize .GoType }} // {{ .GoType }}
	{{ .Name }}Interface, {{ .Name }}IsSet := m["{{ snakize .Name }}"]
	if {{ .Name }}IsSet {
		var items []interface{}
		if listItems, isList := {{ .Name }}Interface.([]interface{}); isList {
			items = listItems
		} else {
			items = {{ .Name }}Interface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := {{ pascalize .GoType }}ModelFromMap(v.(map[string]interface{}))
			{{ varname .Name }} = append({{ varname .Name }}, *m)
		}
	}
				{{- else if .IsAliased }}
	var {{ varname .Name }} *models.{{ pascalize .GoType }} // {{ .GoType }}
	{{ .Name }}Interface, {{ .Name }}IsSet := m["{{ snakize .Name }}"]
	if {{ .Name }}IsSet {
	    {{ .Name }}Model := {{ .Name }}Interface.(string)
	    {{ varname .Name }} = models.New{{.GoType}}(models.{{.GoType}}({{ .Name }}Model))
	}
				{{- else }}
	{{ varname .Name }} := m["{{ snakize .Name }}"].(models.{{ pascalize .GoType }}) // {{ .GoType }}
				{{- end }}
			{{- else }}
	{{ varname .Name }} := m["{{ snakize .Name }}"].({{ .GoType }})
			{{- end }}
		{{- end }}
	{{- end }}

	{{- if not .GenSchema.IsComplexObject }}
		{{- if .GenSchema.IsAliased }}
   {{ varname .Name }} := m["{{ snakize .Name }}"].(models.{{ $operationGroup }})
	return &{{ varname .Name }}
		{{- else }}
	return m["{{ snakize .Name }}"].(models.{{ $operationGroup }})
		{{- end }}
	{{- else }}
	return &models.{{ $operationGroup }} {
		{{- if not $isReadOnlyModel }}
		{{- range .Properties }}
			{{- if or (eq .Name "adminState") (eq .Name "modelID")}}
			{{- end }}
			{{- if or (not .ReadOnly) (eq .Name "id") }}
				{{- if and .IsNullable (and (not .IsComplexObject) (not .IsAliased)) }}
		{{ pascalize .Name }}: &{{ varname .Name }},
				{{- else if and (and (not .IsMap) .IsNullable (not .IsSuperAlias) (not .IsAliased)) (not .IsComplexObject) }}
		{{ pascalize .Name }}: &{{ varname .Name }},
				{{- else }}
		{{ pascalize .Name }}: {{ varname .Name }},
				{{- end }}
			{{- end }}
		{{- end }}
		{{- end }}
	}
	{{- end }}
}

func Set{{ $operationGroup }}ResourceData(d *schema.ResourceData, m *models.{{ $operationGroup }}) {
	{{- range .Properties }}
		{{- if (eq .Name "id") }}
	d.Set("id", m.ID)
		{{- else if or .IsComplexObject }}
	d.Set("{{ snakize .Name }}", Set{{ pascalize .GoType }}SubResourceData([]*models.{{ pascalize .GoType }}{m.{{ pascalize .Name }}}))
		{{- else if stringContains .GoType "[]*" }}
	d.Set("{{ snakize .Name }}", Set{{ pascalize .Items.GoType }}SubResourceData(m.{{ pascalize .Name }}))
		{{- else if .IsBase64 }}
	d.Set("{{ snakize .Name }}", m.{{ pascalize .Name }}.String())
		{{- else }}
	d.Set("{{ snakize .Name }}", m.{{ pascalize .Name }})
		{{- end }}
	{{- end }}
}

func Set{{ $operationGroup }}SubResourceData(m []*models.{{ $operationGroup }}) (d []*map[string]interface{}) {
	{{- $model := camelize $operationGroup }}
	{{- $model = print $operationGroup "Model" }}
	for _, {{ $model }} := range m {
		if {{ $model }} != nil {
			properties := make(map[string]interface{})
			{{- range .Properties }}
				{{- if or .IsComplexObject }}
			properties["{{snakize .Name}}"] = Set{{ pascalize .GoType }}SubResourceData([]*models.{{ pascalize .GoType }}{ {{- $model }}.{{ pascalize .Name -}} })
				{{- else if stringContains .GoType "[]*" }}
			properties["{{snakize .Name}}"] = Set{{ pascalize .Items.GoType }}SubResourceData({{ $model }}.{{ pascalize .Name }})
			{{- else if eq .GoType "strfmt.DateTime" }}
			properties["{{ snakize .Name }}"] = {{ $model }}.{{ pascalize .Name }}.String()
		        {{- else if .IsBase64 }}
			properties["{{ snakize .Name }}"] = {{ $model }}.{{ pascalize .Name }}.String()
				{{- else }}
			properties["{{ snakize .Name }}"] = {{ $model }}.{{ pascalize .Name }}
				{{- end }}
			{{- end }}
			d = append(d, &properties)
		}
	}
	return
}

{{/* Only resources have respective data sources that need to be mapped to an appropriate schema */}}
{{- if eq .Example "isResource" }}
func DataSource{{ $operationGroup }}Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		{{- range .Properties }}
		"{{ snakize .Name }}": {
			Description: `{{ .Description }}`,
				{{- if (eq .Name "id") }}
			Type: schema.TypeString,
			Computed: true,
				{{- else }}
					{{- if eq .GoType "string" }}
			Type: schema.TypeString,
						{{- if .Default }}
			Default: "{{ .Default }}",
						{{- end }}
					{{- else if eq .GoType "bool" }}
			Type: schema.TypeBool,
					{{- if .Default }}
			Default: {{ .Default }},
						{{- end }}
					{{- else if eq .GoType "int64" }}
			Type: schema.TypeInt,
					{{- else if eq .GoType "int32" }}
			Type: schema.TypeInt,
					{{- else if eq .GoType "uint64" }}
			Type: schema.TypeInt,
					{{- else if eq .GoType "uint32" }}
			Type: schema.TypeInt,
					{{- else if eq .GoType "float64" }}
			Type: schema.TypeFloat,
					{{- else if eq .GoType "float32" }}
			Type: schema.TypeFloat,
					{{- else if eq .GoType "strfmt.DateTime" }}
            Type:         schema.TypeString,
            ValidateFunc: validation.IsRFC3339Time,
					{{- else if eq .GoType "strfmt.Base64" }}
            Type:         schema.TypeString,
					{{- else if eq .GoType "map[string]bool" }}
			Type: schema.TypeMap, //GoType: {{ .GoType }}
			Elem: &schema.Schema{
				Type: schema.TypeBool,
			},
					{{- else if eq .GoType "map[string]int64" }}
			Type: schema.TypeMap, //GoType: {{ .GoType }}
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
					{{- else if eq .GoType "map[string]string" }}
			Type: schema.TypeMap, //GoType: {{ .GoType }}
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
					{{- else if eq .GoType "interface{}" }}
			Type: schema.TypeMap, //GoType: {{ .GoType }}
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
					{{- else }}
			Type: schema.TypeList, //GoType: {{ .GoType }}
						{{- if stringContains .GoType "[]*" }}
			Elem: &schema.Resource{
				Schema: {{ .Items.GoType }}Schema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
					    {{- else if eq .GoType "[]float64" }}
			Elem: &schema.Schema{
				Type: schema.TypeFloat,
			},
					    {{- else if eq .GoType "[]string" }}
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
						{{- else if stringContains .GoType "[]" }}
			Elem: &schema.Resource{
				Schema: {{ .Items.GoType }}Schema(),
			},
						{{- else }}
			Elem: &schema.Resource{
				Schema: {{ .GoType }}Schema(),
			},
						{{- end }}
					{{- end }}
					{{- if and .Required (not .ReadOnly) }}
			Required: true,
					{{- else if .ReadOnly }}
			Computed: true,
					{{- else }}
			Optional: true,
					{{- end }}
				{{- end }}
		},
		{{ end }}
	}
}
{{- end }}

func {{ $operationGroup }}Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		{{- range .Properties }}
		"{{ snakize .Name }}": {
			Description: `{{ .Description }}`,
				{{- if (eq .Name "id") }}
			Type: schema.TypeString,
			Computed: true,
				{{- else }}
					{{- if eq .GoType "string" }}
			Type: schema.TypeString,
						{{- if .Default }}
			Default: "{{ .Default }}",
						{{- end }}
					{{- else if eq .GoType "bool" }}
			Type: schema.TypeBool,
					{{- if .Default }}
			Default: {{ .Default }},
						{{- end }}
					{{- else if eq .GoType "int64" }}
			Type: schema.TypeInt,
					{{- else if eq .GoType "int32" }}
			Type: schema.TypeInt,
					{{- else if eq .GoType "uint64" }}
			Type: schema.TypeInt,
					{{- else if eq .GoType "uint32" }}
			Type: schema.TypeInt,
					{{- else if eq .GoType "float64" }}
			Type: schema.TypeFloat,
					{{- else if eq .GoType "float32" }}
			Type: schema.TypeFloat,
					{{- else if eq .GoType "strfmt.DateTime" }}
            Type:         schema.TypeString,
            ValidateFunc: validation.IsRFC3339Time,
					{{- else if eq .GoType "strfmt.Base64" }}
            Type:         schema.TypeString,
					{{- else if eq .GoType "map[string]bool" }}
			Type: schema.TypeMap, //GoType: {{ .GoType }}
			Elem: &schema.Schema{
				Type: schema.TypeBool,
			},
					{{- else if eq .GoType "map[string]int64" }}
			Type: schema.TypeMap, //GoType: {{ .GoType }}
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
					{{- else if eq .GoType "map[string]string" }}
			Type: schema.TypeMap, //GoType: {{ .GoType }}
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
					{{- else if eq .GoType "interface{}" }}
			Type: schema.TypeMap, //GoType: {{ .GoType }}
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
					{{- else if stringContains .GoType "[]" }}
			Type: schema.TypeList, //GoType: {{ .GoType }}
						{{- if stringContains .GoType "[]*" }}
			Elem: &schema.Resource{
				Schema: {{ .Items.GoType }}Schema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
					    {{- else if eq .GoType "[]float64" }}
			Elem: &schema.Schema{
				Type: schema.TypeFloat,
			},
					    {{- else if eq .GoType "[]string" }}
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
						{{- else if stringContains .GoType "[]" }}
			Elem: &schema.Resource{
				Schema: {{ .Items.GoType }}Schema(),
			},
						{{- else }}
			Elem: &schema.Resource{
				Schema: {{ .GoType }}Schema(),
			},
						{{- end }}

			        {{- else if .IsComplexObject }}
			Type: schema.TypeList, //GoType: {{ .GoType }}
			Elem: &schema.Resource{
				Schema: {{ .GoType }}Schema(),
			},
					{{- else }}
			Type: schema.TypeString,
					{{- end }}
					{{- if and .Required (not .ReadOnly) }}
			Required: true,
					{{- else if .ReadOnly }}
			Computed: true,
					{{- else }}
			Optional: true,
					{{- end }}
				{{- end }}
		},
		{{ end }}
	}
}


func Get{{ $operationGroup }}PropertyFields() (t []string) {
	return []string{
		{{- range .Properties }}
			{{- if or (not .ReadOnly) (eq .Name "id") }}
		"{{ snakize .Name }}",
			{{- end }}
		{{- end }}
	}
}
