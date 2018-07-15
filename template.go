package svgsaurus

import (
	"html/template"
)

var svgTemplate *template.Template

func init() {
	var err error
	svgTemplate, err = template.New("svg").Parse(svgTemplateString)
	if err != nil {
		panic(err)
	}
}

var svgTemplateString = `
	<svg xmlns="http://www.w3.org/2000/svg" height="{{.Height}}" viewBox="0 0 {{.Width}} {{.Height}}">
		<style>
			text {
				font-family: {{.Font}}, sans-serif;
				font-size: {{.Size}}px;

				{{- if .Bold -}}
					font-weight: bold;
				{{- end}}

				{{- if .Italic -}}
					font-style: italic;
				{{- end}}

				{{- if .Underline -}}
					text-decoration: underline;
				{{- end}}
			}
		</style>

		{{if .Background -}}
			<rect
				width="{{.Width}}"
				height="{{.Height}}"
				style="fill:#{{.Background}};"
			/>
		{{- end}}

		<text
			x="{{.X}}"
			y="{{.Y}}"
			fill="#{{.Color}}"
		>[[Text]]</text>
	</svg>
`
