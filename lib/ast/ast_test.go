package ast_test

import (
	"go/parser"
	"testing"

	libast "github.com/nurcahyaari/kite/lib/ast"
	"github.com/stretchr/testify/assert"
)

func TestAddImport(t *testing.T) {
	t.Run("Test add import 1", func(t *testing.T) {
		code := `
		package test
		import (
			a "path/of/package/a"
			"path/of/package/b"
			usermodel "path/model/user"
		)

		func main() {
			
		}
		`

		abstractCode := libast.NewAbstractCode(code, parser.ParseComments)
		abstractCode.AddImport(libast.ImportSpec{
			Name: "c",
			Path: "\"path/of/package/c\"",
		})
		err := abstractCode.RebuildCode()
		act := abstractCode.GetCode()
		// don't touch the expected code please
		// expect you change the test case please :D
		// building this string obviously difficult
		exp := `package test

import (
	a "path/of/package/a"
	"path/of/package/b"
	usermodel "path/model/user"
	c "path/of/package/c"
)

func main() {

}
`

		assert.NoError(t, err)
		assert.Equal(t, exp, act)
	})

	t.Run("Test add import 2", func(t *testing.T) {
		code := `
		package test
		import (
			a "path/of/package/a"
			"path/of/package/b"
			usermodel "path/model/user"
		)

		const a = "init"

		type b int

		func main() {
			
		}
		`

		abstractCode := libast.NewAbstractCode(code, parser.ParseComments)
		abstractCode.AddImport(libast.ImportSpec{
			Name: "c",
			Path: "\"path/of/package/c\"",
		})
		err := abstractCode.RebuildCode()
		act := abstractCode.GetCode()
		// don't touch the expected code please
		// expect you change the test case please :D
		// building this string obviously difficult
		exp := `package test

import (
	a "path/of/package/a"
	"path/of/package/b"
	usermodel "path/model/user"
	c "path/of/package/c"
)

const a = "init"

type b int

func main() {

}
`

		assert.NoError(t, err)
		assert.Equal(t, exp, act)
	})
}
