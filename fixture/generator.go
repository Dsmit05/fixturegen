package fixture

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/tools/go/ast/inspector"
)

const commentCommand = "//fixtureGen"

type Generator struct {
	inPath string
}

func NewGenerator(inPath string) (*Generator, error) {
	if inPath == "" {
		return nil, ErrPathFile
	}
	return &Generator{inPath: inPath}, nil
}

func (g *Generator) Start() error {
	fileSet := token.NewFileSet()
	astInFile, err := parser.ParseFile(fileSet, g.inPath, nil, parser.ParseComments)
	if err != nil {
		return errors.Wrap(err, "parser.ParseFile() err")
	}

	genTasks := g.GenetareTask(astInFile)
	if len(genTasks) == 0 {
		return ErrZeroGenerateTask
	}

	astOutFile := &ast.File{Name: astInFile.Name}

	for _, task := range genTasks {
		err := task.Generate(astOutFile)
		if err != nil {
			return errors.Wrap(err, "task.Generate() err")
		}
	}

	outFile, err := os.Create(strings.TrimSuffix(g.inPath, ".go") + "_fgen.go")
	if err != nil {
		return errors.Wrap(err, "os.Create() err")
	}
	defer outFile.Close()

	err = printer.Fprint(outFile, token.NewFileSet(), astOutFile)
	if err != nil {
		return errors.Wrap(err, "printer.Fprint() err")
	}

	return nil
}

func (g *Generator) GenetareTask(astInFile *ast.File) []fixtureGenerator {
	var genTasks []fixtureGenerator

	ir := inspector.New([]*ast.File{astInFile})
	iFilter := []ast.Node{
		&ast.GenDecl{},
	}

	ir.Nodes(iFilter, func(node ast.Node, push bool) (proceed bool) {
		genDecl := node.(*ast.GenDecl)
		if genDecl.Doc == nil {
			return false
		}

		typeSpec, ok := genDecl.Specs[0].(*ast.TypeSpec)
		if !ok {
			return false
		}

		structType, ok := typeSpec.Type.(*ast.StructType)
		if !ok {
			return false
		}

		for _, comment := range genDecl.Doc.List {
			if comment.Text == commentCommand {
				genTasks = append(genTasks, fixtureGenerator{
					typeSpec:   typeSpec,
					structType: structType,
				})
			}
		}

		return false
	})

	return genTasks
}
