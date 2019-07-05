package director

import (
	"bytes"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/kyma-incubator/compass/components/director/pkg/graphql"
	"github.com/pkg/errors"
)

// graphqlizer is responsible for converting Go objects to input arguments in graphql format
type graphqlizer struct{}

func (g *graphqlizer) ApplicationInputToGQL(in graphql.ApplicationInput) (string, error) {
	return g.genericToGQL(in, `{
		name: "{{.Name}}",
		{{- if .Description }}
		description: "{{.Description}}",
		{{- end }}
        {{- if .Labels }}
		labels: {{ LabelsToGQL .Labels}},
		{{- end }}
		{{- if .Annotations }}
		annotations: {{ AnnotationsToGQL .Annotations}},
		{{- end}}
		{{- if .Webhooks }}
		webhooks: [
			{{- range $i, $e := .Webhooks }} 
				{{- if $i}}, {{- end}} {{ ApplicationWebhookInputToGQL $e }}
			{{- end }} ],
		{{- end}}
		{{- if .HealthCheckURL }}
		healthCheckURL: "{{ .HealthCheckURL }}"
		{{- end }}
		{{- if .Apis }}
		apis: [
			{{- range $i, $e := .Apis }}
			{{- if $i}}, {{- end}} {{ APIDefinitionInputToGQL $e }}
			{{- end }}]
		{{- end }}
		{{- if .EventAPIs }}
		eventAPIs: [
			{{- range $i, $e := .EventAPIs }}
			{{- if $i}}, {{- end}} {{ EventAPIDefinitionInputToGQL $e }}
			{{- end }}]
		{{- end }}
		{{- if .Documents }} 
		documents: [
			{{- range $i, $e := .Documents }} 
				{{- if $i}}, {{- end}} {{- DocumentInputToGQL $e }}
			{{- end }} ]
		{{- end }}
	}`)
}

func (g *graphqlizer) DocumentInputToGQL(in *graphql.DocumentInput) (string, error) {
	return g.genericToGQL(in, `{
		title: "{{.Title}}",
		displayName: "{{.DisplayName}}",
		description: "{{.Description}}",
		format: {{.Format}},
		{{- if .Kind }}
		kind: "{{.Kind}}",
		{{- end}}
		{{- if .Data }}
		data: "{{.Data}}",
		{{- end}}
		{{- if .FetchRequest }}
		fetchRequest: {{- FetchRequesstInputToGQL .FetchRequest }} 
		{{- end}}
		
}`)
}

func (g *graphqlizer) FetchRequestInputToGQL(in *graphql.FetchRequestInput) (string, error) {
	return g.genericToGQL(in, `{
		url: "{{.URL}}",
		{{- if .Auth }}
		auth: {{- AuthInputToGQL .Auth }}
		{{- end }}
		{{- if .Mode }}
		mode: {{.Mode}},
		{{- end}}
		{{- if .Filter}}
		filter: "{{.Filter}}",
		{{- end}}
	}`)
}

func (g *graphqlizer) AuthInputToGQL(in *graphql.AuthInput) (string, error) {
	return g.genericToGQL(in, `{
		credential:{
			{{- if .Credential.Basic }}
			basic: {
				username: "{{ .Credential.Basic.Username}}",
				password: "{{ .Credential.Basic.Password}}",
			}
			{{- end }}
			{{- if .Credential.Oauth }}
			oauth: {
				clientId: "{{ .Credential.Oauth.ClientID}}",
				clientSecret: "{{ .Credential.Oauth.ClientSecret }}",
				url: "{{ .Credential.Oauth.URL }}",
			}
			{{- end }}
		},
		{{- if .AdditionalHeaders }}
		additionalHeaders: {{ HTTPHeadersToGQL .AdditionalHeaders }},
		{{- end }}
		{{- if .AdditionalQueryParams }}
		additionalQueryParams: {{ QueryParamsToGQL .AdditionalQueryParams}},
		{{- end }}
		{{- if .RequestAuth }}

		{{- end }}
	}`)
}

func (g *graphqlizer) LabelsToGQL(in graphql.Labels) (string, error) {
	return g.genericToGQL(in, `{
		{{- range $k,$v := . }}
			{{$k}}: [
				{{- range $i,$j := $v }}
					{{- if $i}},{{- end}}"{{$j}}"
				{{- end }} ]
		{{- end}}
	}
		`)
}

func (g *graphqlizer) HTTPHeadersToGQL(in graphql.HttpHeaders) (string, error) {
	return g.genericToGQL(in, `{
		{{- range $k,$v := . }}
			{{$k}}: [
				{{- range $i,$j := $v }}
					{{- if $i}},{{- end}}"{{$j}}"
				{{- end }} ],
		{{- end}}
	}
		`)
}

func (g *graphqlizer) QueryParamsToGQL(in graphql.QueryParams) (string, error) {
	return g.genericToGQL(in, `{
		{{- range $k,$v := . }}
			{{$k}}: [
				{{- range $i,$j := $v }}
					{{- if $i}},{{- end}}"{{$j}}"
				{{- end }} ],
		{{- end}}
	}
		`)
}

func (g *graphqlizer) AnnotationsToGQL(in graphql.Annotations) (string, error) {
	return g.genericToGQL(in, `{
		{{- range $k,$v := . }}
			{{$k}}: "{{$v}}"
		{{- end}}
	}
		`)
}

func (g *graphqlizer) ApplicationWebhookInputToGQL(in *graphql.ApplicationWebhookInput) (string, error) {
	return g.genericToGQL(in, `{
		type: {{.Type}},
		url: "{{.URL }}",
		{{- if .Auth }} 
		auth: {{- AuthInputToGQL .Auth }}
		{{- end }}

	}`)
}

func (g *graphqlizer) APIDefinitionInputToGQL(in graphql.APIDefinitionInput) (string, error) {
	return g.genericToGQL(in, `{
		name: "{{ .Name}}",
		{{- if .Description }}
		description: "{{.Description}}",
		{{- end}}
		targetURL: "{{.TargetURL}}",
		{{- if .Group }}
		group: "{{.Group}}",
		{{- end }}
		{{- if .Spec }}
		spec: {{- ApiSpecInputToGQL .Spec }},
		{{- end }}
		{{- if .Version }}
		version: {{- VersionInputToGQL .Version }},
		{{- end}}
		{{- if .DefaultAuth }}
		defaultAuth: {{- AuthInputToGQL .DefaultAuth}},
		{{- end}}
	}`)
}

func (g *graphqlizer) EventAPIDefinitionInputToGQL(in graphql.EventAPIDefinitionInput) (string, error) {
	return g.genericToGQL(in, `{
		name: "{{.Name}}",
		{{- if .Description }}
		description: "{{.Description}}",
		{{- end }}
		spec: {{ EventAPISpecInputToGQL .Spec }},
		{{- if .Group }}
		group: "{{.Group}}", 
		{{- end }}
		{{- if .Version }}
		version: {{- VersionInputToGQL .Version }},
		{{- end}}
	}`)
}

func (g *graphqlizer) EventAPISpecInputToGQL(in graphql.EventAPISpecInput) (string, error) {
	return g.genericToGQL(in, `{
		{{- if .Data }}
		data: "{{.Data}}",
		{{- end }}
		eventSpecType: {{.EventSpecType}},
		{{- if .FetchRequest }}
		fetchRequest: {{- FetchRequesstInputToGQL .FetchRequest }},
		{{- end }}
		format: {{.Format}}
	}`)
}

func (g *graphqlizer) ApiSpecInputToGQL(in graphql.APISpecInput) (string, error) {
	return g.genericToGQL(in, `{
		{{- if .Data}}
		data: "{{.Data}}",
		{{- end}}	
		type: {{.Type}},
		format: {{.Format}},
		{{- if .FetchRequest }}
		fetchRequest: {{- FetchRequesstInputToGQL .FetchRequest }},
		{{- end }}
	}`)
}

func (g *graphqlizer) VersionInputToGQL(in graphql.VersionInput) (string, error) {
	return g.genericToGQL(in, `{
		value: "{{.Value}}",
		{{- if .Deprecated }}
		deprecated: {{.Deprecated}},
		{{- end}}
		{{- if .DeprecatedSince }}
		deprecatedSince: "{{.DeprecatedSince}}",
		{{- end}}
		{{- if .ForRemoval }}
		forRemoval: {{.ForRemoval }}
		{{- end }}
	}`)
}

func (g *graphqlizer) RuntimeInputToGQL(in graphql.RuntimeInput) (string, error) {
	return g.genericToGQL(in, `{
		name: "{{.Name}}",
		{{- if .Description }}
		description: "{{.Description}}",
		{{- end }}
		{{- if .Labels }}
		labels: {{ LabelsToGQL .Labels}},
		{{- end }}
		{{- if .Annotations }}
		annotations: {{AnnotationsToGQL .Annotations}}
		{{- end }}
	}`)
}

func (g *graphqlizer) genericToGQL(obj interface{}, tmpl string) (string, error) {
	fm := sprig.TxtFuncMap()
	fm["DocumentInputToGQL"] = g.DocumentInputToGQL
	fm["FetchRequesstInputToGQL"] = g.FetchRequestInputToGQL
	fm["AuthInputToGQL"] = g.AuthInputToGQL
	fm["LabelsToGQL"] = g.LabelsToGQL
	fm["AnnotationsToGQL"] = g.AnnotationsToGQL
	fm["ApplicationWebhookInputToGQL"] = g.ApplicationWebhookInputToGQL
	fm["APIDefinitionInputToGQL"] = g.APIDefinitionInputToGQL
	fm["EventAPIDefinitionInputToGQL"] = g.EventAPIDefinitionInputToGQL
	fm["ApiSpecInputToGQL"] = g.ApiSpecInputToGQL
	fm["VersionInputToGQL"] = g.VersionInputToGQL
	fm["HTTPHeadersToGQL"] = g.HTTPHeadersToGQL
	fm["QueryParamsToGQL"] = g.QueryParamsToGQL
	fm["EventAPISpecInputToGQL"] = g.EventAPISpecInputToGQL

	t, err := template.New("tmpl").Funcs(fm).Parse(tmpl)
	if err != nil {
		return "", errors.Wrapf(err, "while parsing template")
	}

	var b bytes.Buffer

	if err := t.Execute(&b, obj); err != nil {
		return "", errors.Wrap(err, "while executing template")
	}
	return b.String(), nil
}
