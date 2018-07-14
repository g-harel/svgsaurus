package main

import (
	"html/template"
	"os"
	"strings"
)

// Template string used to render the output.
var t = `
<svg xmlns="http://www.w3.org/2000/svg"
	 height="{{.Height}}"
	 viewBox="-{{.X}} -{{.Y}} {{.Width}} {{.Height}}">
	<style>
		text {
			font-weight: {{.Boldness}};
			font-family: {{.Font}};
			font-size: {{.Size}}px;
			{{if .Italics -}}
				font-style:italic;
			{{- end}}
			{{if .Underline -}}
				text-decoration: underline;
			{{- end}}
		}
	</style>
	<text fill="#{{.Color}}">{{.Text}}</text>
</svg>
`

// Config contains the data rendered by the template.
type Config struct {
	Text string

	Size   string
	Width  string
	Height string
	X      string
	Y      string

	Font      string
	Color     string
	Boldness  string
	Italics   bool
	Underline bool
}

// FromMap populates the Config's values from a query string type.
func (c *Config) FromMap(m map[string][]string) *Config {
	fallback := func(key string, def string) string {
		list := m[key]
		if list == nil || len(list) < 1 {
			return def
		}
		return list[0]
	}

	c.Text = fallback("t", "svgsaurus")

	// Use the replacement pattern to insert spaces
	replace := fallback("r", "_")
	c.Text = strings.Replace(c.Text, replace, " ", -1)

	c.Size = fallback("s", "55")
	c.Width = fallback("w", "290")
	c.Height = fallback("h", "60")
	c.X = fallback("x", "5")
	c.Y = fallback("y", "45")

	c.Font = fallback("f", "sans-serif")
	c.Color = fallback("c", "ff1493")
	c.Boldness = fallback("b", "600")
	options := fallback("o", "i")
	c.Italics = strings.Index(options, "i") >= 0
	c.Underline = strings.Index(options, "u") >= 0

	return c
}

func main() {
	v := map[string][]string{}

	c := new(Config).FromMap(v)

	tmpl, err := template.New("test").Parse(t)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, c)
	if err != nil {
		panic(err)
	}
}
