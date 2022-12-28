{{- $operationGroup := pascalize .Name -}}
{{- $hasID := false -}}

{{/* Some property fields can be characterized as more "complex" objects whose data must be retrieved with specialized helper functions */}}
{{- $isReadOnlyModel := true -}}
{{- range .Properties -}}
	{{- if not .ReadOnly -}} {{- $isReadOnlyModel = false -}} {{- end -}}
	{{- if eq .Name "customProperties" -}} {{- $needsUtils = true -}} {{- end -}}
	{{- if eq .Name "id" -}} {{- $hasID = true -}} {{- end -}}
{{- end }}
/*
resource "{{ snakize $operationGroup }}"  "{{ snakize $operationGroup }}_test" {
		# computed
		{{- range .Properties }}
				{{- if or .ReadOnly }}
    # {{ snakize .Name }} =
		        {{- end }}
		{{- end }}

		# required
		{{- range .Properties }}
				{{- if .Required }}
		{{ snakize .Name}} 
					{{- if eq .GoType "string" }}
			        {{- if .Default }}"{{ .Default }}"{{- else }} = ""{{- end }}
					{{- else if eq .GoType "bool" }}
					{{- if .Default }}{{ .Default }}{{- else }} = false{{- end }}
					{{- else if eq .GoType "int64" }}
					{{- if .Default }}{{ .Default }}{{- else }} = 0{{- end }}
					{{- else if eq .GoType "int32" }}
					{{- if .Default }}{{ .Default }}{{- else }} = 0{{- end }}
					{{- else if eq .GoType "uint64" }}
					{{- if .Default }}{{ .Default }}{{- else }} = 0{{- end }}
					{{- else if eq .GoType "uint32" }}
					{{- if .Default }}{{ .Default }}{{- else }} = 0{{- end }}
					{{- else if eq .GoType "float64" }}
					{{- if .Default }}{{ .Default }}{{- else }} = 0.0{{- end }}
					{{- else if eq .GoType "float32" }}
					{{- if .Default }}{{ .Default }}{{- else }} = 0.0{{- end }}
					{{- else if eq .GoType "strfmt.DateTime" }}
					{{- if .Default }}{{ .Default }}{{- else }} = 0.0{{- end }}
					{{- else if eq .GoType "strfmt.Base64" }}
					{{- if .Default }}{{ .Default }}{{- else }} = ""{{- end }}
					{{- else if eq .GoType "map[string]bool" }}
					{{- if .Default }}{{ .Default }}{{- else }} {"" = false}{{- end }}
					{{- else if eq .GoType "map[string]int64" }}
					{{- if .Default }}{{ .Default }}{{- else }} {"" = 0}{{- end }}
					{{- else if eq .GoType "map[string]string" }}
					{{- if .Default }}{{ .Default }}{{- else }} {"" = ""}{{- end }}
					{{- else if eq .GoType "interface{}" }}
					{{- if .Default }}{{ .Default }}{{- else }} = {{- end }}
					{{- else if stringContains .GoType "[]" }}
						{{- if stringContains .GoType "[]*" }}
					{{- if .Default }}{{ .Default }}{{- else }} = []{{- end }}
					    {{- else if eq .GoType "[]float64" }}
					{{- if .Default }}{{ .Default }}{{- else }} = [0]{{- end }}
					    {{- else if eq .GoType "[]string" }}
					{{- if .Default }}{{ .Default }}{{- else }} = [""]{{- end }}
					    {{- else}}
					{{- if .Default }}{{ .Default }}{{- else }} = []{{- end }}
						{{- end }}
			        {{- else if .IsComplexObject }}
					{{- if .Default }}{{ .Default }}{{- else }} = # {{.GoType}}{{- end }}
					{{- else }}
			        {{- if .Default }}"{{ .Default }}"{{- else }} = ""{{- end }}
					{{- end }}
		        {{- end}}
		{{- end }}

		# optional
		{{- range .Properties }}
				{{- if and (not .Required) (not .ReadOnly) }}
		{{ snakize .Name}}
					{{- if eq .GoType "string" }}
			        {{- if .Default }}"{{ .Default }}"{{- else }} = ""{{- end }}
					{{- else if eq .GoType "bool" }}
					{{- if .Default }}{{ .Default }}{{- else }} = false{{- end }}
					{{- else if eq .GoType "int64" }}
					{{- if .Default }}{{ .Default }}{{- else }} = 0{{- end }}
					{{- else if eq .GoType "int32" }}
					{{- if .Default }}{{ .Default }}{{- else }} = 0{{- end }}
					{{- else if eq .GoType "uint64" }}
					{{- if .Default }}{{ .Default }}{{- else }} = 0{{- end }}
					{{- else if eq .GoType "uint32" }}
					{{- if .Default }}{{ .Default }}{{- else }} = 0{{- end }}
					{{- else if eq .GoType "float64" }}
					{{- if .Default }}{{ .Default }}{{- else }} = 0.0{{- end }}
					{{- else if eq .GoType "float32" }}
					{{- if .Default }}{{ .Default }}{{- else }} = 0.0{{- end }}
					{{- else if eq .GoType "strfmt.DateTime" }}
					{{- if .Default }}{{ .Default }}{{- else }} = 0.0{{- end }}
					{{- else if eq .GoType "strfmt.Base64" }}
					{{- if .Default }}{{ .Default }}{{- else }} = ""{{- end }}
					{{- else if eq .GoType "map[string]bool" }}
					{{- if .Default }}{{ .Default }}{{- else }} {"" = false}{{- end }}
					{{- else if eq .GoType "map[string]int64" }}
					{{- if .Default }}{{ .Default }}{{- else }} {"" = 0}{{- end }}
					{{- else if eq .GoType "map[string]string" }}
					{{- if .Default }}{{ .Default }}{{- else }} {"" = ""}{{- end }}
					{{- else if eq .GoType "interface{}" }}
					{{- if .Default }}{{ .Default }}{{- else }} = {{- end }}
					{{- else if stringContains .GoType "[]" }}
						{{- if stringContains .GoType "[]*" }}
					{{- if .Default }}{{ .Default }}{{- else }} = []{{- end }}
					    {{- else if eq .GoType "[]float64" }}
					{{- if .Default }}{{ .Default }}{{- else }} = [0]{{- end }}
					    {{- else if eq .GoType "[]string" }}
					{{- if .Default }}{{ .Default }}{{- else }} = [""]{{- end }}
					    {{- else}}
					{{- if .Default }}{{ .Default }}{{- else }} = []{{- end }}
						{{- end }}
			        {{- else if .IsComplexObject }}
					{{- if .Default }}{{ .Default }}{{- else }} = # {{.GoType}}{{- end }}
					{{- else }}
			        {{- if .Default }}"{{ .Default }}"{{- else }} = ""{{- end }}
					{{- end }}
		        {{- end}}
		{{- end }}
}
*/
