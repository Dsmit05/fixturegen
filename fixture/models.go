package fixture

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"

	"github.com/pkg/errors"
)

type fixtureValue struct {
	StructName string
	FieldName  string
	FieldType  string
}

type fixtureGenerator struct {
	typeSpec   *ast.TypeSpec
	structType *ast.StructType
}

// Generate add in file template
func (f fixtureGenerator) Generate(outFile *ast.File) error {
	var buf bytes.Buffer

	params := struct {
		StructName string
	}{StructName: f.typeSpec.Name.Name}

	err := builderTemplate.Execute(&buf, params)
	if err != nil {
		return errors.Wrap(err, "builderTemplate.Execute() err")
	}

	for _, field := range f.structType.Fields.List {
		fieldTypeName, err := f.castInString(field.Type)
		if err != nil {
			continue
		}

		var fieldName string
		if len(field.Names) > 0 {
			fieldName = field.Names[0].Name
		} else {
			fieldName = fieldTypeName
		}

		templ := fixtureValue{
			StructName: f.typeSpec.Name.Name,
			FieldName:  fieldName,
			FieldType:  fieldTypeName,
		}

		err = fixtureFieldTemplate.Execute(&buf, templ)
		if err != nil {
			return errors.Wrap(err, "builderTemplate.Execute() err")
		}

	}

	templateAst, err := parser.ParseFile(token.NewFileSet(), "", buf.Bytes(), parser.ParseComments)
	if err != nil {
		return errors.Wrap(err, "parser.ParseFile() err")
	}

	outFile.Decls = append(outFile.Decls, templateAst.Decls...)

	return nil
}

func (f *fixtureGenerator) castInString(expr ast.Expr) (string, error) {
	var buf bytes.Buffer
	err := printer.Fprint(&buf, token.NewFileSet(), expr)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
