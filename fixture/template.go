package fixture

import "text/template"

var builderTemplate = template.Must(template.New("").Parse(
	`
package fixtures

type {{ .StructName }}Builder struct {
	s *{{ .StructName }}
}

func {{ .StructName }}BuilderNew() *{{ .StructName }}Builder {
	return &{{ .StructName }}Builder{
		s: &{{ .StructName }}{},
	}
}

func (b *{{ .StructName }}Builder) P() *{{ .StructName }} {
	return b.s
}

func (b *{{ .StructName }}Builder) V() {{ .StructName }} {
	return *b.s
}

`))

var fixtureFieldTemplate = template.Must(template.New("").Parse(`

func (b *{{ .StructName }}Builder) {{ .FieldName }}(v {{ .FieldType }}) *{{ .StructName }}Builder {
	b.s.{{ .FieldName }} = v
	return b
}

`))
