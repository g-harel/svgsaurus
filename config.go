package svgsaurus

import (
	"bytes"
	"regexp"
	"strings"
)

var whitespace = regexp.MustCompile("\\s+")

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
}

// FromQuery populates the Config's values from a url query.
func (c *Config) FromQuery(q map[string]string) *Config {
	fallback := func(key string, def string) string {
		val, ok := q[key]
		if !ok {
			return def
		}
		return val
	}

	c.Text = fallback("t", "svgsaurus")

	c.Size = fallback("s", "55")
	c.Width = fallback("w", "265")
	c.Height = fallback("h", "60")
	c.X = fallback("x", "5")
	c.Y = fallback("y", "46")

	c.Font = fallback("f", "arial")
	c.Color = fallback("c", "000000")
	c.Background = fallback("b", "")
	options := fallback("o", "")
	c.Bold = strings.Index(options, "b") >= 0
	c.Italic = strings.Index(options, "i") >= 0
	c.Underline = strings.Index(options, "u") >= 0

	// Use the replacement pattern to insert spaces.
	replace := fallback("r", "_")
	c.Text = strings.Replace(c.Text, replace, " ", -1)
	c.Font = strings.Replace(c.Font, replace, " ", -1)

	return c
}

// Render generates the output string from a Config.
func (c *Config) Render() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := svgTemplate.Execute(buf, c)
	if err != nil {
		return nil, err
	}

	// Remove extra whitespace without modifying the text.
	b := whitespace.ReplaceAllLiteral(buf.Bytes(), []byte{' '})
	b = bytes.Replace(b, []byte("[[Text]]"), []byte(c.Text), -1)

	return b, nil
}
