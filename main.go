package main

import (
	"html/template"
	"os"
)

// Template string used to render the output.
var t = `
<svg xmlns="http://www.w3.org/2000/svg" height="100" viewBox="0 -75 150 150">
	<style>
		text {
			font: italic 75px sans-serif;
		}
	</style>
	<text x="0" y="0" fill="orange">{{.Text}}</text>
</svg>
`

// Config contains the data used to render the template.
type Config struct {
	Text string
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

	c.Text = fallback("t", "Hello World!")

	return c
}

func main() {
	v := map[string][]string{
		"t": {"✨"},
	}

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
