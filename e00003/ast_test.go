package e00003

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"testing"
)

func Test_CanParseAST(t *testing.T) {
	fset := token.NewFileSet()

	f, err := os.OpenFile("./../file.go", os.O_RDONLY, 0222)
	if err != nil {
		log.Fatal("unable to open file:", err)
	}

	parsedFile, err := parser.ParseFile(fset, "./../file.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	ast.Inspect(parsedFile, func(n ast.Node) bool {
		it, ok := n.(*ast.CallExpr)
		if ok {
			f.Seek(int64(it.Fun.Pos()-1), 0)
			buf := make([]byte, it.Fun.End()-it.Fun.Pos())
			f.Read(buf)
			fmt.Println(string(buf))
			fname := string(buf)
			if fname != "hello" {
				return true
			}

			for i, v := range it.Args {
				fmt.Printf("%d => %v\n", i, v)
			}

			return false
		}
		return true
	})
}
