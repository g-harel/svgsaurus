package main

import (
	"encoding/base64"
	"html/template"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// Template string used to render the output.
var t = `
<svg xmlns="http://www.w3.org/2000/svg"
	 height="{{.Height}}"
	 viewBox="0 0 {{.Width}} {{.Height}}">
	<style>
		{{if .FontData -}}
			<![CDATA[
				@font-face {
					font-family: '{{.Font}}';
					src: url('data:application/x-font-ttf;base64,{{.FontData}}');
				}
			]]>
		{{- end}}
		text {
			font-family: {{.Font}}, sans-serif;
			font-size: {{.Size}}px;
			{{if .Bold -}}
				font-weight: bold;
			{{- end}}
			{{if .Italic -}}
				font-style: italic;
			{{- end}}
			{{if .Underline -}}
				text-decoration: underline;
			{{- end}}
		}
	</style>
	{{if .Background -}}
		<rect width="{{.Width}}" height="{{.Height}}" style="fill:#{{.Background}};" />
	{{- end}}
	<text x="{{.X}}" y="{{.Y}}" fill="#{{.Color}}">{{.Text}}</text>
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

	Font       string
	Color      string
	Background string
	Bold       bool
	Italic     bool
	Underline  bool

	FontData string
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

	c.Text = fallback("t", "svgsaurus!")

	c.Size = fallback("s", "55")
	c.Width = fallback("w", "340")
	c.Height = fallback("h", "60")
	c.X = fallback("x", "5")
	c.Y = fallback("y", "43")

	c.Font = fallback("f", "courier_new")
	c.Color = fallback("c", "0dd1e0")
	c.Background = fallback("b", "1d0d33")
	options := fallback("o", "bi")
	c.Bold = strings.Index(options, "b") >= 0
	c.Italic = strings.Index(options, "i") >= 0
	c.Underline = strings.Index(options, "u") >= 0

	// Use the replacement pattern to insert spaces
	replace := fallback("r", "_")
	c.Text = strings.Replace(c.Text, replace, " ", -1)
	c.Font = strings.Replace(c.Font, replace, " ", -1)

	return c
}

// LoadFontData attempts to load font data into the config.
func (c *Config) LoadFontData() *Config {
	parts := []string{c.Font}
	if c.Bold {
		parts = append(parts, "bold")
	}
	if c.Italic {
		parts = append(parts, "italic")
	}
	name := strings.Join(parts, "-")

	dir, err := os.Getwd()
	if err != nil {
		return c
	}

	preserveTransforms := false
	raw, err := ioutil.ReadFile(path.Join(dir, "fonts", name+".ttf"))
	if err != nil {
		// Do not retry if there were no transforms
		if c.Font == name {
			return c
		}
		raw, err = ioutil.ReadFile(path.Join(dir, "fonts", c.Font+".ttf"))
		if err != nil {
			return c
		}
		preserveTransforms = true
	}

	c.FontData = base64.StdEncoding.EncodeToString(raw)

	// Font is already transformed in ttf file.
	if !preserveTransforms {
		c.Bold = false
		c.Italic = false
	}

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
