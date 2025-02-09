package main

import (
	"log"
	"math"
	"os"
	"strings"
	"text/template"

	"github.com/perpetuatheme/go"
)

//go:generate go run generate.go

func main() {
	f, err := os.Create("../README.md")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	name := "README.tmpl.md"
	t := template.Must(template.New(name).ParseFiles(name))

	err = t.Execute(f, struct {
		PaletteLight string
		PaletteDark  string
	}{
		PaletteLight: generate(perpetua.Light),
		PaletteDark:  generate(perpetua.Dark),
	})
	if err != nil {
		log.Fatal(err)
	}
}

func generate(p perpetua.Palette) string {
	builder := new(strings.Builder)
	t := template.Must(template.New("tmplTableRow").Parse(tmplTableRow))
	for i, c := range colors(p) {
		if c.Index != uint(i) {
			panic("colors out of order")
		}
		r, g, b, _ := c.RGBA()
		h, s, l := c.Okhsl()
		err := t.Execute(builder, struct {
			PaletteName                string
			SwatchName                 string
			Name                       string
			Hex                        string
			Red, Green, Blue           uint8
			Hue, Saturation, Lightness uint
		}{
			PaletteName: p.Name(),
			SwatchName:  swatchName(p.Name(), c.Name),
			Name:        c.Name,
			Hex:         c.Hex(),
			Red:         uint8(r), Green: uint8(g), Blue: uint8(b),
			Hue: approxOkhsl(h), Saturation: approxOkhsl(s), Lightness: approxOkhsl(l),
		})
		if err != nil {
			log.Fatal(err)
		}
	}
	return strings.TrimSuffix(strings.ReplaceAll(builder.String(), "@", "            "), "\n")
}

func colors(p perpetua.Palette) []perpetua.Color {
	return []perpetua.Color{
		p.Red(),
		p.Orange(),
		p.Yellow(),
		p.Lime(),
		p.Green(),
		p.Turquoise(),
		p.Cyan(),
		p.Cerulean(),
		p.Blue(),
		p.Violet(),
		p.Lavender(),
		p.Pink(),
		p.RedBack(),
		p.OrangeBack(),
		p.YellowBack(),
		p.LimeBack(),
		p.GreenBack(),
		p.TurquoiseBack(),
		p.CyanBack(),
		p.CeruleanBack(),
		p.BlueBack(),
		p.VioletBack(),
		p.LavenderBack(),
		p.PinkBack(),
		p.Base0(),
		p.Base1(),
		p.Base2(),
		p.Base3(),
		p.Base4(),
		p.Base5(),
		p.Over0(),
		p.Over1(),
		p.Over2(),
		p.Text0(),
		p.Text1(),
		p.Text2(),
	}
}

func swatchName(palette, color string) string {
	s := strings.ToLower(palette + "_" + color)
	if strings.Contains(s, "back") {
		return strings.ReplaceAll(s, " ", "_")
	}
	return strings.ReplaceAll(s, " ", "")
}

func approxOkhsl(v float64) uint {
	if v-math.Floor(v) >= 0.6 {
		return uint(math.Ceil(v))
	}
	return uint(math.Floor(v))
}

// NOTE: Replace @ to apply indentation.
const tmplTableRow = `@<tr>
@    <td align="center"><img alt="Perpetua {{ .PaletteName }}: {{ .Name }}" height="22" src="./assets/swatch_{{ .SwatchName }}.png"></td>
@    <td>{{ .Name }}</td>
@    <td><code>{{ .Hex }}</code></td>
@    <td><code>rgb({{ .Red }}, {{ .Green }}, {{ .Blue }})</code></td>
@    <td><code>okhsl({{ .Hue }}, {{ .Saturation }}%, {{ .Lightness }}%)</code></td>
@</tr>
`
