package resources

import (
	"github.com/davecgh/go-spew/spew"
	httptransport "github.com/go-openapi/runtime/client"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	zschema "github.com/zededa/terraform-provider-zedcloud/v2/schemas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

{{- $operationGroup := .Name -}} {{/* operation groups map to OpenAPI tags */}}
/*
{{ pascalize $operationGroup }} {{ if .Summary }}{{ pluralizeFirstWord (humanize .Summary) }}{{ if .Description }}

{{ blockcomment .Description }}{{ end }}{{ else if .Description}}{{ blockcomment .Description }}{{ else }}{{ humanize $operationGroup }} API{{ end }}
*/

func {{ pascalize $operationGroup }}() *schema.Resource {
	return &schema.Resource{
	    /*
		{{- range .Operations }}
			{{- $operation := .Name }}
			{{- if not (stringContains $operation "List") }} {{/* GetList endpoints are used for datasources, not resources */}}
			{{- if eq .Method "GET" }}
		ReadContext: {{ $operation }},
			{{- end }}
			{{- if eq .Method "POST" }}
		CreateContext: {{ $operation }},
			{{- end }}
			{{- if eq .Method "PUT" }} {{/* this template can be easily expanded to use PATCH instead of PUT if your API supports it */}}
		UpdateContext: {{ $operation }},
			{{- end }}
			{{- if eq .Method "DELETE" }}
		DeleteContext: {{ $operation }},
			{{- end }}
			{{- end }}
		{{- end }}
		*/
		Schema: zschema.{{- pascalize $operationGroup -}}Schema(),
	}
}

func DataResource{{ pascalize $operationGroup }}() *schema.Resource {
	return &schema.Resource{
		{{- range .Operations }} {{/* skip everything except GetList opertaions */}}
			{{- $operation := .Name }}
			{{- if stringContains $operation "List" }}
			{{- if eq .Method "GET" }}
		ReadContext: {{ $operation }},
			{{- end }}
			{{- end}}
		{{- end }}
		Schema: zschema.{{- pascalize $operationGroup -}}Schema(),
	}
}

{{ range .Operations }}
	{{- $operation := .Name }}
	{{- if eq .Method "GET" }} {{/* this is either the Get endpoint used by the resource or the GetList endpoint used by the datasource */}}
func {{ $operation }}(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := {{ $operationGroup }}.New{{ pascalize $operation }}Params()
	{{ range .Params }}
		{{- template "generateParams" . -}}
	{{ end }}

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.{{ pascalize $operationGroup }}.{{ pascalize $operation }}(params, nil)
	if err != nil {
        log.Printf("[TRACE] {{ pascalize $operationGroup }}.{{ pascalize $operation }} error: %s", spew.Sdump(err))
        if ds, ok := ZsrvResponderToDiags(err); ok {
            diags = append(diags, ds...)
            return diags
        }

        diags = append(diags, diag.Errorf("{{ pascalize $operationGroup }}.{{ pascalize $operation }} error: %s", err)...)
        return diags
    }

	respModel := resp.GetPayload()
	{{- if stringContains $operation "List" }}
	if len(respModel.List) == 0 {
		return append(diags, diag.Errorf("no devices found")...)
	}
	// limit output to a single result
	result := respModel.List[0]

	// the zedcloud API does not return the partially updated object but a custom response.
	// thus, we need to fetch the object and populate the state.
	d.SetId(result.ID)
	return GetDevice(ctx, d, m)
	{{- else }}
	zschema.Set{{- pascalize $operationGroup -}}ResourceData(d, respModel)

	return diags
	{{- end}}
}
	{{- end }}
{{ end }}

{{ range .Operations }}
	{{- $operation := .Name }}
	{{- if eq .Method "POST" }}
func {{ $operation }}(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	model := zschema.{{ pascalize $operationGroup }}Model(d)
	params := {{ $operationGroup }}.New{{ pascalize $operation }}Params()
	params.SetBody(model)

	client := m.(*api_client.ZedcloudAPI)

	resp, err := client.{{ pascalize $operationGroup }}.{{ pascalize $operation }}(params, nil)
	if err != nil {
        log.Printf("[TRACE] {{ pascalize $operationGroup }}.{{ pascalize $operation }} error: %s", spew.Sdump(err))
        if ds, ok := ZsrvResponderToDiags(err); ok {
            diags = append(diags, ds...)
            return diags
        }

        diags = append(diags, diag.Errorf("{{ pascalize $operationGroup }}.{{ pascalize $operation }} error: %s", err)...)
        return diags
    }

	responseData := resp.GetPayload()
	if responseData != nil && len(responseData.Error) > 0 {
		for _, err := range responseData.Error {
			// FIXME: zedcloud api returns a response that contains and error even in case of success.
			// remove this code once it is fixed on API side.
			if err.ErrorCode != nil && *err.ErrorCode == models.ErrorCodeSuccess {
				continue
			}
			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
		}
		if diags.HasError() {
			return diags
		}
	}

	// note, we need to set the ID in any case, even GetByName endpoint seems to require the ID
    // but doesn't return any error if it's not set.
    d.SetId(responseData.ObjectID)

    // the zedcloud API does not return the partially updated object but a custom response.
	// thus, we need to fetch the object and populate the state.
    if errs := GetDevice(ctx, d, m); errs != nil {
		return append(diags, errs...)
    }

	return diags
}
	{{- end }}
{{ end }}

{{ range .Operations }}
	{{- $operation := .Name }}
	{{- if eq .Method "PUT" }}
func {{ $operation }}(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	d.Partial(true)

	params := {{ $operationGroup }}.New{{ pascalize $operation }}Params()
	{{ range .Params }}
		{{ if eq (camelize .ID) "body" }}
			params.SetBody(zschema.{{pascalize $operationGroup }}Model(d))
			// {{.GoType}}
		{{ else }}
			{{- template "generateParams" . -}}
		{{ end }}
	{{ end }}


	// makes a bulk update for all properties that were changed
	client := m.(*api_client.ZedcloudAPI)
	resp, err := client.{{ pascalize $operationGroup }}.{{ pascalize $operation }}(params, nil)
	if err != nil {
        log.Printf("[TRACE] {{ pascalize $operationGroup }}.{{ pascalize $operation }} error: %s", spew.Sdump(err))
        if ds, ok := ZsrvResponderToDiags(err); ok {
            diags = append(diags, ds...)
            return diags
        }

        diags = append(diags, diag.Errorf("{{ pascalize $operationGroup }}.{{ pascalize $operation }} error: %s", err)...)
        return diags
    }

	responseData := resp.GetPayload()
	if responseData != nil && len(responseData.Error) > 0 {
		for _, err := range responseData.Error {
			// FIXME: zedcloud api returns a response that contains and error even in case of success.
			// remove this code once it is fixed on API side.
			if err.ErrorCode != nil && *err.ErrorCode == models.ErrorCodeSuccess {
				continue
			}
			diags = append(diags, diag.FromErr(errors.New(err.Details))...)
		}
		if diags.HasError() {
			return diags
		}
	}

    // the zedcloud API does not return the partially updated object but a custom response.
	// thus, we need to fetch the object and populate the state.
    if errs := GetDevice(ctx, d, m); errs != nil {
		return append(diags, errs...)
    }

	d.Partial(false)

	return diags
}
	{{- end }}
{{ end }}

{{ range .Operations }}
	{{- $operation := .Name }}
	{{- if eq .Method "DELETE" }}
func {{ $operation }}(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	params := {{ $operationGroup }}.New{{ pascalize $operation }}Params()

	{{ range .Params }}
		{{- template "generateParams" . -}}
	{{ end }}

	client := m.(*api_client.ZedcloudAPI)

	_, err := client.{{ pascalize $operationGroup }}.{{ pascalize $operation }}(params, nil)
	if err != nil {
        log.Printf("[TRACE] {{ pascalize $operationGroup }}.{{ pascalize $operation }} error: %s", spew.Sdump(err))
        if ds, ok := ZsrvResponderToDiags(err); ok {
            diags = append(diags, ds...)
            return diags
        }

        diags = append(diags, diag.Errorf("{{ pascalize $operationGroup }}.{{ pascalize $operation }} error: %s", err)...)
        return diags
    }

	d.SetId("")
	return diags
}
	{{- end }}
{{ end }}

{{ define "generateParams" }}
	{{ camelize .ID }}Val, {{ camelize .ID }}IsSet := d.GetOk("{{ snakize .ID}}")
	if({{ camelize .ID }}IsSet){
		{{- if eq (camelize .ID) "id" }}
			{{ camelize .ID }}, _ := {{ camelize .ID }}Val.(string)
			params.{{ pascalize .ID}} = {{ if and (not .IsArray) (not .IsMap) (not .HasDiscriminator) (not .IsInterface) (not .IsStream) (or .IsNullable  ) }}{{ end }}{{ if not .IsFileParam }}{{ if and (not .IsArray) (not .IsMap) (not .HasDiscriminator) (not .IsInterface) (not .IsStream) (or .IsNullable  ) }}&{{ end }}{{ camelize .ID }}{{ else }}runtime.NamedReadCloser{{- end -}}

		{{- else }}
			params.{{ pascalize .ID}} = {{ camelize .ID }}Val.({{ if and (not .IsArray) (not .IsMap) (not .HasDiscriminator) (not .IsInterface) (not .IsStream) (or .IsNullable  ) }}*{{ end }}{{ .GoType }})
		{{- end }}
	} {{ if .Required }} else {
		diags = append(diags, diag.Errorf("missing client parameter: {{ .Name }}")...)
		return diags
	} {{ end }}
{{ end }}

