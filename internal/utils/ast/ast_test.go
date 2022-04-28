package ast_test

import (
	"go/parser"
	"testing"

	"github.com/nurcahyaari/kite/internal/utils/ast"
	libast "github.com/nurcahyaari/kite/internal/utils/ast"
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

	t.Run("Test add import 3", func(t *testing.T) {
		code := `
			package test
			import "path/of/package/a"

			const a = "init"

			type b int

			func main() {

			}
			`

		abstractCode := libast.NewAbstractCode(code, parser.ParseComments)
		abstractCode.AddImport(libast.ImportSpec{
			Name: "b",
			Path: "\"path/of/package/b\"",
		})
		err := abstractCode.RebuildCode()
		act := abstractCode.GetCode()
		// don't touch the expected code please
		// expect you change the test case please :D
		// building this string obviously difficult
		exp := `package test

import (
	"path/of/package/a"
	b "path/of/package/b"
)

const a = "init"

type b int

func main() {

}
`

		assert.NoError(t, err)
		assert.Equal(t, exp, act)
	})

	t.Run("Test add import 4", func(t *testing.T) {
		code := `
		package test

		const a = "init"

		type b int

		func main() {
			
		}
		`

		abstractCode := libast.NewAbstractCode(code, parser.ParseComments)
		abstractCode.AddImport(libast.ImportSpec{
			Name: "b",
			Path: "\"path/of/package/b\"",
		})
		err := abstractCode.RebuildCode()
		act := abstractCode.GetCode()
		// don't touch the expected code please
		// expect you change the test case please :D
		// building this string obviously difficult
		exp := `package test

import b "path/of/package/b"

const a = "init"

type b int

func main() {

}
`

		assert.NoError(t, err)
		assert.Equal(t, exp, act)
	})

	t.Run("Test add import 5", func(t *testing.T) {
		code := `
		package test
		`

		abstractCode := libast.NewAbstractCode(code, parser.ParseComments)
		abstractCode.AddImport(libast.ImportSpec{
			Path: "\"path/of/package/b\"",
		})
		err := abstractCode.RebuildCode()
		act := abstractCode.GetCode()
		// don't touch the expected code please
		// expect you change the test case please :D
		// building this string obviously difficult
		exp := `package test

import  "path/of/package/b"
`

		assert.NoError(t, err)
		assert.Equal(t, exp, act)
	})
}

func TestAddInterface(t *testing.T) {
	t.Run("Test add interfaces 1", func(t *testing.T) {
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
		abstractCode.AddInterfaces(libast.InterfaceSpecList{
			&libast.InterfaceSpec{
				Name: "Test",
			},
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
)

type Test interface {
}

func main() {

}
`

		assert.NoError(t, err)
		assert.Equal(t, exp, act)
	})

	t.Run("Test add interfaces 2", func(t *testing.T) {
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
		abstractCode.AddInterfaces(libast.InterfaceSpecList{
			&libast.InterfaceSpec{
				Name: "Test",
			},
			&libast.InterfaceSpec{
				Name: "Test2",
			},
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
)

type Test2 interface {
}
type Test interface {
}

func main() {

}
`

		assert.NoError(t, err)
		assert.Equal(t, exp, act)
	})

	t.Run("Test add interfaces 3", func(t *testing.T) {
		code := `
		package test

		func main() {
			
		}
		`

		abstractCode := libast.NewAbstractCode(code, parser.ParseComments)
		abstractCode.AddInterfaces(libast.InterfaceSpecList{
			&libast.InterfaceSpec{
				Name: "Test",
			},
			&libast.InterfaceSpec{
				Name: "Test2",
			},
		})
		err := abstractCode.RebuildCode()
		act := abstractCode.GetCode()
		// don't touch the expected code please
		// expect you change the test case please :D
		// building this string obviously difficult
		exp := `package test

type Test2 interface {
}
type Test interface {
}

func main() {

}
`

		assert.NoError(t, err)
		assert.Equal(t, exp, act)
	})

	t.Run("Test add interfaces 4", func(t *testing.T) {
		code := `
		package test

		type TestImpl struct {
		}

		func main() {
			
		}
		`

		abstractCode := libast.NewAbstractCode(code, parser.ParseComments)
		abstractCode.AddInterfaces(libast.InterfaceSpecList{
			&libast.InterfaceSpec{
				Name:       "Test",
				StructName: "TestImpl",
			},
		})
		err := abstractCode.RebuildCode()
		act := abstractCode.GetCode()
		// don't touch the expected code please
		// expect you change the test case please :D
		// building this string obviously difficult
		exp := `package test

type Test interface {
}

type TestImpl struct {
}

func main() {

}
`
		assert.NoError(t, err)
		assert.Equal(t, exp, act)
	})
}

func TestAddStructs(t *testing.T) {
	t.Run("Test add structs 1", func(t *testing.T) {
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
		abstractCode.AddStructs(libast.StructSpecList{
			&libast.StructSpec{
				Name: "Test",
			},
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
)

type Test struct {
}

func main() {

}
`

		assert.NoError(t, err)
		assert.Equal(t, exp, act)
	})

	t.Run("Test add structs 2", func(t *testing.T) {
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
		abstractCode.AddStructs(libast.StructSpecList{
			&libast.StructSpec{
				Name: "Test",
			},
			&libast.StructSpec{
				Name: "Test2",
			},
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
)

type Test2 struct {
}
type Test struct {
}

func main() {

}
`

		assert.NoError(t, err)
		assert.Equal(t, exp, act)
	})

	t.Run("Test add structs 3", func(t *testing.T) {
		code := `
			package test

			func main() {

			}
			`

		abstractCode := libast.NewAbstractCode(code, parser.ParseComments)
		abstractCode.AddStructs(libast.StructSpecList{
			&libast.StructSpec{
				Name: "Test",
			},
			&libast.StructSpec{
				Name: "Test2",
			},
		})
		err := abstractCode.RebuildCode()
		act := abstractCode.GetCode()
		// don't touch the expected code please
		// expect you change the test case please :D
		// building this string obviously difficult
		exp := `package test

type Test2 struct {
}
type Test struct {
}

func main() {

}
`

		assert.NoError(t, err)
		assert.Equal(t, exp, act)
	})

	t.Run("Test add structs 4", func(t *testing.T) {
		code := `
			package test

			type Test interface {
			}

			func main() {

			}
			`

		abstractCode := libast.NewAbstractCode(code, parser.ParseComments)
		abstractCode.AddStructs(libast.StructSpecList{
			&libast.StructSpec{
				Name:          "TestImpl",
				InterfaceName: "Test",
			},
		})
		err := abstractCode.RebuildCode()
		act := abstractCode.GetCode()
		// don't touch the expected code please
		// expect you change the test case please :D
		// building this string obviously difficult
		exp := `package test

type Test interface {
}
type TestImpl struct {
}

func main() {

}
`

		assert.NoError(t, err)
		assert.Equal(t, exp, act)
	})
}

func TestAddStructVarDecl(t *testing.T) {
	t.Run("Test add struct var decl 1", func(t *testing.T) {
		code := `
		package test
		import (
			a "path/of/package/a"
			"path/of/package/b"
			usermodel "path/model/user"
		)

		type Http struct {}

		func main() {
			
		}
		`

		abstractCode := libast.NewAbstractCode(code, parser.ParseComments)
		abstractCode.AddStructVarDecl(libast.StructArgList{
			&libast.StructArg{
				StructName: "Http",
				Name:       "HttpRouter",
				DataType: libast.StructDtypes{
					LibName:  "router",
					TypeName: "HttpRouterImpl",
				},
			},
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
)

type Http struct{ HttpRouter router.HttpRouterImpl }

func main() {

}
`

		assert.NoError(t, err)
		assert.Equal(t, exp, act)
	})

	t.Run("Test add struct var decl 2", func(t *testing.T) {
		code := `
		package test
		import (
			a "path/of/package/a"
			"path/of/package/b"
			usermodel "path/model/user"
		)

		type Http struct {}

		func main() {
			
		}
		`

		abstractCode := libast.NewAbstractCode(code, parser.ParseComments)
		abstractCode.AddStructVarDecl(libast.StructArgList{
			&libast.StructArg{
				StructName: "Http",
				Name:       "HttpRouter",
				DataType: libast.StructDtypes{
					LibName:  "router",
					TypeName: "HttpRouterImpl",
				},
			},
			&libast.StructArg{
				StructName: "Http",
				Name:       "HttpRouter2",
				DataType: libast.StructDtypes{
					LibName:  "router",
					TypeName: "HttpRouterImpl",
				},
			},
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
)

type Http struct {
	HttpRouter	router.HttpRouterImpl
	HttpRouter2	router.HttpRouterImpl
}

func main() {

}
`

		assert.NoError(t, err)
		assert.Equal(t, exp, act)
	})

	t.Run("Test add struct var decl 3", func(t *testing.T) {
		code := `
		package test
		import (
			a "path/of/package/a"
			"path/of/package/b"
			usermodel "path/model/user"
		)

		type Http struct {}

		func main() {
			
		}
		`

		abstractCode := libast.NewAbstractCode(code, parser.ParseComments)
		abstractCode.AddStructVarDecl(libast.StructArgList{
			&libast.StructArg{
				StructName: "Http",
				Name:       "HttpRouter",
				DataType: libast.StructDtypes{
					LibName:  "router",
					TypeName: "HttpRouterImpl",
				},
				IsPointer: true,
			},
			&libast.StructArg{
				StructName: "Http",
				Name:       "HttpRouter2",
				DataType: libast.StructDtypes{
					LibName:  "router",
					TypeName: "HttpRouterImpl",
				},
			},
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
)

type Http struct {
	HttpRouter	*router.HttpRouterImpl
	HttpRouter2	router.HttpRouterImpl
}

func main() {

}
`

		assert.NoError(t, err)
		assert.Equal(t, exp, act)
	})

	t.Run("Test add struct var decl 4", func(t *testing.T) {
		code := `
		package test
		import (
			a "path/of/package/a"
			"path/of/package/b"
			usermodel "path/model/user"
		)

		type Http struct {}

		func main() {
			
		}
		`

		abstractCode := libast.NewAbstractCode(code, parser.ParseComments)
		abstractCode.AddStructVarDecl(libast.StructArgList{
			&libast.StructArg{
				StructName: "Http",
				DataType: libast.StructDtypes{
					LibName:  "router",
					TypeName: "HttpRouterImpl",
				},
				IsPointer: true,
			},
			&libast.StructArg{
				StructName: "Http",
				Name:       "HttpRouter2",
				DataType: libast.StructDtypes{
					LibName:  "router",
					TypeName: "HttpRouterImpl",
				},
			},
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
)

type Http struct {
	*router.HttpRouterImpl
	HttpRouter2	router.HttpRouterImpl
}

func main() {

}
`

		assert.NoError(t, err)
		assert.Equal(t, exp, act)
	})
}

func TestAddFunction(t *testing.T) {
	t.Run("Test add function 1", func(t *testing.T) {
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
		abstractCode.AddFunction(libast.FunctionSpecList{
			&libast.FunctionSpec{
				Name: "Test",
			},
			&libast.FunctionSpec{
				Name: "Test2",
			},
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
)

func Test2() {
}
func Test() {
}

func main() {

}
`

		assert.NoError(t, err)
		assert.Equal(t, exp, act)
	})

	t.Run("Test add function 2", func(t *testing.T) {
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
		abstractCode.AddFunction(libast.FunctionSpecList{
			&libast.FunctionSpec{
				Name: "Test",
				Args: libast.FunctionArgList{
					&libast.FunctionArg{
						Name:     "age",
						DataType: "int",
					},
					&libast.FunctionArg{
						Name:      "name",
						DataType:  "string",
						IsPointer: true,
					},
				},
			},
			&libast.FunctionSpec{
				Name: "Test2",
			},
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
)

func Test2() {
}
func Test(age int, name *string) {
}

func main() {

}
`

		assert.NoError(t, err)
		assert.Equal(t, exp, act)
	})

	t.Run("Test add function 3", func(t *testing.T) {
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
		abstractCode.AddFunction(libast.FunctionSpecList{
			&libast.FunctionSpec{
				Name: "Test",
				Args: libast.FunctionArgList{
					&libast.FunctionArg{
						Name:     "age",
						DataType: "int",
					},
					&libast.FunctionArg{
						Name:      "name",
						DataType:  "string",
						IsPointer: true,
					},
				},
			},
			&libast.FunctionSpec{
				Name: "Test2",
				Returns: &libast.FunctionReturnSpecList{
					&libast.FunctionReturnSpec{
						IsPointer: true,
						LibName:   "http",
						DataType:  "HttpInterface",
						Return:    "HttpInterface",
					},
				},
			},
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
)

func Test2() *http.HttpInterface {
	return &http.HttpInterface{}
}
func Test(age int, name *string) {
}

func main() {

}
`

		assert.NoError(t, err)
		assert.Equal(t, exp, act)
	})

	t.Run("Test add function 4", func(t *testing.T) {
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
		abstractCode.AddFunction(libast.FunctionSpecList{
			&libast.FunctionSpec{
				Name: "Test",
				Args: libast.FunctionArgList{
					&libast.FunctionArg{
						Name:     "age",
						DataType: "int",
					},
					&libast.FunctionArg{
						Name:      "name",
						DataType:  "string",
						IsPointer: true,
					},
				},
			},
			&libast.FunctionSpec{
				Name: "Test2",
				Returns: &libast.FunctionReturnSpecList{
					&libast.FunctionReturnSpec{
						IsPointer: true,
						LibName:   "http",
						DataType:  "HttpInterface",
						Return:    "HttpInterface",
					},
					&libast.FunctionReturnSpec{
						DataType: "error",
						Return:   "nil",
					},
				},
			},
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
)

func Test2() (*http.HttpInterface, error) {
	return &http.HttpInterface{}, nil
}
func Test(age int, name *string) {
}

func main() {

}
`

		assert.NoError(t, err)
		assert.Equal(t, exp, act)
	})

	t.Run("Test add function 5", func(t *testing.T) {
		code := `
		package test
		import (
			a "path/of/package/a"
			"path/of/package/b"
			usermodel "path/model/user"
		)

		type HttpImpl struct{}

		func main() {
			
		}
		`

		abstractCode := libast.NewAbstractCode(code, parser.ParseComments)
		abstractCode.AddFunction(libast.FunctionSpecList{
			&libast.FunctionSpec{
				Name: "Test",
				Args: libast.FunctionArgList{
					&libast.FunctionArg{
						Name:     "age",
						DataType: "int",
					},
					&libast.FunctionArg{
						Name:      "name",
						DataType:  "string",
						IsPointer: true,
					},
				},
			},
			&libast.FunctionSpec{
				Name: "Test2",
				StructSpec: &libast.FunctionStructSpec{
					IsConstruct: true,
					DataTypes:   "HttpImpl",
				},
				Returns: &libast.FunctionReturnSpecList{
					&libast.FunctionReturnSpec{
						IsPointer: true,
						LibName:   "http",
						DataType:  "HttpInterface",
						Return:    "HttpInterface",
					},
				},
			},
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
)

func Test(age int, name *string) {
}

type HttpImpl struct{}

func Test2() *http.HttpInterface {
	return &http.HttpInterface{}
}

func main() {

}
`

		assert.NoError(t, err)
		assert.Equal(t, exp, act)
	})

	t.Run("Test add function 6", func(t *testing.T) {
		code := `
		package test
		import (
			a "path/of/package/a"
			"path/of/package/b"
			usermodel "path/model/user"
		)

		type HttpImpl struct{}

		func main() {
			
		}
		`

		abstractCode := libast.NewAbstractCode(code, parser.ParseComments)
		abstractCode.AddFunction(libast.FunctionSpecList{
			&libast.FunctionSpec{
				Name: "Test",
				Args: libast.FunctionArgList{
					&libast.FunctionArg{
						Name:     "age",
						DataType: "int",
					},
					&libast.FunctionArg{
						Name:      "name",
						DataType:  "string",
						IsPointer: true,
					},
				},
			},
			&libast.FunctionSpec{
				Name: "Test2",
				StructSpec: &libast.FunctionStructSpec{
					Name:      "h",
					DataTypes: "HttpImpl",
				},
			},
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
)

func Test(age int, name *string) {
}

type HttpImpl struct{}

func (h HttpImpl) Test2() {
}

func main() {

}
`

		assert.NoError(t, err)
		assert.Equal(t, exp, act)
	})

	t.Run("Test add function 7", func(t *testing.T) {
		code := `
		package test
		import (
			a "path/of/package/a"
			"path/of/package/b"
			usermodel "path/model/user"
		)

		type HttpImpl struct{}

		func main() {
			
		}
		`

		abstractCode := libast.NewAbstractCode(code, parser.ParseComments)
		abstractCode.AddFunction(libast.FunctionSpecList{
			&libast.FunctionSpec{
				Name: "Test",
				Args: libast.FunctionArgList{
					&libast.FunctionArg{
						Name:     "age",
						DataType: "int",
					},
					&libast.FunctionArg{
						Name:      "name",
						DataType:  "string",
						IsPointer: true,
					},
				},
			},
			&libast.FunctionSpec{
				Name: "Test2",
				StructSpec: &libast.FunctionStructSpec{
					Name:      "h",
					DataTypes: "HttpImpl",
					IsPointer: true,
				},
			},
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
)

func Test(age int, name *string) {
}

type HttpImpl struct{}

func (h *HttpImpl) Test2() {
}

func main() {

}
`

		assert.NoError(t, err)
		assert.Equal(t, exp, act)
	})

	t.Run("Test add function 8", func(t *testing.T) {
		code := `
		package test
		import (
			a "path/of/package/a"
			"path/of/package/b"
			usermodel "path/model/user"
		)

		type HttpImpl struct{}

		func main() {
			
		}
		`

		abstractCode := libast.NewAbstractCode(code, parser.ParseComments)
		abstractCode.AddFunction(libast.FunctionSpecList{
			&libast.FunctionSpec{
				Name: "Test",
				Args: libast.FunctionArgList{
					&libast.FunctionArg{
						Name:     "age",
						DataType: "int",
					},
					&libast.FunctionArg{
						Name:      "name",
						DataType:  "string",
						IsPointer: true,
					},
				},
			},
			&libast.FunctionSpec{
				Name: "Test2",
				StructSpec: &libast.FunctionStructSpec{
					Name:      "h",
					DataTypes: "HttpImpl",
					IsPointer: true,
				},
			},
			&libast.FunctionSpec{
				Name: "Test3",
				StructSpec: &libast.FunctionStructSpec{
					Name:      "h",
					DataTypes: "HttpImpl",
					IsPointer: true,
				},
			},
			&libast.FunctionSpec{
				Name: "Test4",
				StructSpec: &libast.FunctionStructSpec{
					Name:      "h",
					DataTypes: "HttpImpl",
					IsPointer: true,
				},
			},
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
)

func Test(age int, name *string) {
}

type HttpImpl struct{}

func (h *HttpImpl) Test2() {
}
func (h *HttpImpl) Test3() {
}
func (h *HttpImpl) Test4() {
}

func main() {

}
`
		assert.NoError(t, err)
		assert.Equal(t, exp, act)
	})

	t.Run("Test add function 9", func(t *testing.T) {
		code := `
		package test
		import (
			a "path/of/package/a"
			"path/of/package/b"
			usermodel "path/model/user"
		)

		type HttpImpl struct{}

		func main() {
			
		}
		`

		abstractCode := libast.NewAbstractCode(code, parser.ParseComments)
		abstractCode.AddFunction(libast.FunctionSpecList{
			&libast.FunctionSpec{
				Name: "Test",
				Args: libast.FunctionArgList{
					&libast.FunctionArg{
						Name:     "age",
						DataType: "int",
					},
					&libast.FunctionArg{
						Name:      "name",
						DataType:  "string",
						IsPointer: true,
					},
				},
			},
			&libast.FunctionSpec{
				Name: "NewHttpImpl",
				StructSpec: &libast.FunctionStructSpec{
					DataTypes: "HttpImpl",
				},
			},
			&libast.FunctionSpec{
				Name: "Test2",
				StructSpec: &libast.FunctionStructSpec{
					Name:      "h",
					DataTypes: "HttpImpl",
					IsPointer: true,
				},
			},
			&libast.FunctionSpec{
				Name: "Test3",
				StructSpec: &libast.FunctionStructSpec{
					Name:      "h",
					DataTypes: "HttpImpl",
					IsPointer: true,
				},
			},
			&libast.FunctionSpec{
				Name: "Test4",
				StructSpec: &libast.FunctionStructSpec{
					Name:      "h",
					DataTypes: "HttpImpl",
					IsPointer: true,
				},
			},
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
)

func Test(age int, name *string) {
}

type HttpImpl struct{}

func NewHttpImpl() {
}
func (h *HttpImpl) Test2() {
}
func (h *HttpImpl) Test3() {
}
func (h *HttpImpl) Test4() {
}

func main() {

}
`
		assert.NoError(t, err)
		assert.Equal(t, exp, act)
	})

	t.Run("Test add function 10", func(t *testing.T) {
		code := `
		package test
		import (
			a "path/of/package/a"
			"path/of/package/b"
			usermodel "path/model/user"
		)

		type HttpImpl struct{}

		func main() {
			
		}
		`

		abstractCode := libast.NewAbstractCode(code, parser.ParseComments)
		abstractCode.AddFunction(libast.FunctionSpecList{
			&libast.FunctionSpec{
				Name: "Test",
				Args: libast.FunctionArgList{
					&libast.FunctionArg{
						Name:     "age",
						DataType: "int",
					},
					&libast.FunctionArg{
						Name:      "name",
						DataType:  "string",
						IsPointer: true,
					},
				},
			},
			&libast.FunctionSpec{
				Name: "NewHttpImpl",
				StructSpec: &libast.FunctionStructSpec{
					DataTypes: "HttpImpl",
				},
				Returns: &libast.FunctionReturnSpecList{
					&libast.FunctionReturnSpec{
						IsStruct: true,
						DataType: "HttpImpl",
						Return:   "HttpImpl",
					},
				},
			},
			&libast.FunctionSpec{
				Name: "Test2",
				StructSpec: &libast.FunctionStructSpec{
					Name:      "h",
					DataTypes: "HttpImpl",
					IsPointer: true,
				},
			},
			&libast.FunctionSpec{
				Name: "Test3",
				StructSpec: &libast.FunctionStructSpec{
					Name:      "h",
					DataTypes: "HttpImpl",
					IsPointer: true,
				},
			},
			&libast.FunctionSpec{
				Name: "Test4",
				StructSpec: &libast.FunctionStructSpec{
					Name:      "h",
					DataTypes: "HttpImpl",
					IsPointer: true,
				},
			},
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
)

func Test(age int, name *string) {
}

type HttpImpl struct{}

func NewHttpImpl() HttpImpl {
	return HttpImpl{}
}
func (h *HttpImpl) Test2() {
}
func (h *HttpImpl) Test3() {
}
func (h *HttpImpl) Test4() {
}

func main() {

}
`
		assert.NoError(t, err)
		assert.Equal(t, exp, act)
	})

	t.Run("Test add function 11", func(t *testing.T) {
		code := `
		package test
		import (
			a "path/of/package/a"
			"path/of/package/b"
			usermodel "path/model/user"
		)

		type HttpImpl struct{}

		func main() {
			
		}
		`

		abstractCode := libast.NewAbstractCode(code, parser.ParseComments)
		abstractCode.AddFunction(libast.FunctionSpecList{
			&libast.FunctionSpec{
				Name: "Test",
				Args: libast.FunctionArgList{
					&libast.FunctionArg{
						Name:     "age",
						DataType: "int",
					},
					&libast.FunctionArg{
						Name:      "name",
						DataType:  "string",
						IsPointer: true,
					},
				},
			},
			&libast.FunctionSpec{
				Name: "NewHttpImpl",
				StructSpec: &libast.FunctionStructSpec{
					DataTypes: "HttpImpl",
				},
				Returns: &libast.FunctionReturnSpecList{
					&libast.FunctionReturnSpec{
						IsPointer: true,
						IsStruct:  true,
						DataType:  "HttpImpl",
						Return:    "HttpImpl",
					},
				},
			},
			&libast.FunctionSpec{
				Name: "Test2",
				StructSpec: &libast.FunctionStructSpec{
					Name:      "h",
					DataTypes: "HttpImpl",
					IsPointer: true,
				},
			},
			&libast.FunctionSpec{
				Name: "Test3",
				StructSpec: &libast.FunctionStructSpec{
					Name:      "h",
					DataTypes: "HttpImpl",
					IsPointer: true,
				},
			},
			&libast.FunctionSpec{
				Name: "Test4",
				StructSpec: &libast.FunctionStructSpec{
					Name:      "h",
					DataTypes: "HttpImpl",
					IsPointer: true,
				},
			},
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
)

func Test(age int, name *string) {
}

type HttpImpl struct{}

func NewHttpImpl() *HttpImpl {
	return &HttpImpl{}
}
func (h *HttpImpl) Test2() {
}
func (h *HttpImpl) Test3() {
}
func (h *HttpImpl) Test4() {
}

func main() {

}
`
		assert.NoError(t, err)
		assert.Equal(t, exp, act)
	})
}

func TestAddFunctionArgs(t *testing.T) {
	t.Run("Test add function args 1", func(t *testing.T) {
		code := `
		package test

		type Http struct{}

		func NewHttp() {}

		func main() {
			
		}
		`

		abstractCode := libast.NewAbstractCode(code, parser.ParseComments)
		abstractCode.AddFunctionArgs(libast.FunctionSpec{
			Name: "NewHttp",
			Args: libast.FunctionArgList{
				&libast.FunctionArg{
					Name:     "usersvc",
					LibName:  "user",
					DataType: "UserService",
				},
			},
		})
		err := abstractCode.RebuildCode()
		act := abstractCode.GetCode()
		// don't touch the expected code please
		// expect you change the test case please :D
		// building this string obviously difficult
		exp := `package test

type Http struct{}

func NewHttp(usersvc user.UserService,)	{}

func main() {

}
`
		assert.NoError(t, err)
		assert.Equal(t, exp, act)
	})

	t.Run("Test add function args 2", func(t *testing.T) {
		code := `
		package test

		type Http struct{
			UserSvc usersvc.UserService
		}

		func main() {
			
		}
		`

		abstractCode := libast.NewAbstractCode(code, parser.ParseComments)
		abstractCode.AddFunction(libast.FunctionSpecList{
			&libast.FunctionSpec{
				Name: "NewHttp",
				StructSpec: &libast.FunctionStructSpec{
					DataTypes:   "Http",
					IsConstruct: true,
				},
				Returns: &libast.FunctionReturnSpecList{
					&libast.FunctionReturnSpec{
						IsPointer: true,
						IsStruct:  true,
						LibName:   "",
						DataType:  "Http",
						Return:    "Http",
					},
				},
			},
		})
		abstractCode.AddFunctionArgs(libast.FunctionSpec{
			Name: "NewHttp",
			Args: libast.FunctionArgList{
				&libast.FunctionArg{
					Name:     "UserSvc",
					LibName:  "usersvc",
					DataType: "UserService",
				},
			},
		})
		err := abstractCode.RebuildCode()
		act := abstractCode.GetCode()
		// don't touch the expected code please
		// expect you change the test case please :D
		// building this string obviously difficult
		exp := `package test

type Http struct {
	UserSvc usersvc.UserService
}

func NewHttp(UserSvc usersvc.UserService) *Http {
	return &Http{}
}

func main() {

}
`
		assert.NoError(t, err)
		assert.Equal(t, exp, act)
	})

	t.Run("Test add function args 3", func(t *testing.T) {
		code := `
		package test

		type Http struct{
			UserSvc usersvc.UserService
		}

		func main() {
			
		}
		`

		abstractCode := libast.NewAbstractCode(code, parser.ParseComments)
		abstractCode.AddFunction(libast.FunctionSpecList{
			&libast.FunctionSpec{
				Name: "NewHttp",
				StructSpec: &libast.FunctionStructSpec{
					DataTypes:   "Http",
					IsConstruct: true,
				},
				Returns: &libast.FunctionReturnSpecList{
					&libast.FunctionReturnSpec{
						IsPointer: true,
						IsStruct:  true,
						LibName:   "",
						DataType:  "Http",
						Return:    "Http",
					},
				},
			},
		})
		abstractCode.AddFunctionArgs(libast.FunctionSpec{
			Name: "NewHttp",
			Args: libast.FunctionArgList{
				&libast.FunctionArg{
					Name:     "UserSvc",
					LibName:  "usersvc",
					DataType: "UserService",
				},
			},
		})
		err := abstractCode.RebuildCode()
		assert.NoError(t, err)
		act := abstractCode.GetCode()

		abstractCode = libast.NewAbstractCode(act, parser.ParseComments)
		abstractCode.AddFunctionArgs(libast.FunctionSpec{
			Name: "NewHttp",
			Args: libast.FunctionArgList{
				&libast.FunctionArg{
					Name:     "ProductSvc",
					LibName:  "productsvc",
					DataType: "ProductService",
				},
			},
		})
		err = abstractCode.RebuildCode()
		act = abstractCode.GetCode()

		// don't touch the expected code please
		// expect you change the test case please :D
		// building this string obviously difficult
		exp := `package test

type Http struct {
	UserSvc usersvc.UserService
}

func NewHttp(UserSvc usersvc.UserService, ProductSvc productsvc.ProductService,) *Http {
	return &Http{}
}

func main() {

}
`
		assert.NoError(t, err)
		assert.Equal(t, exp, act)
	})
}

func TestAddFunctionReturns(t *testing.T) {
	t.Run("Test add function return 1", func(t *testing.T) {
		code := `
		package test

		type Http struct{
			UserSvc usersvc.UserService
		}

		func main() {
			
		}
		`

		abstractCode := libast.NewAbstractCode(code, parser.ParseComments)
		abstractCode.AddFunction(libast.FunctionSpecList{
			&libast.FunctionSpec{
				Name: "NewHttp",
				StructSpec: &libast.FunctionStructSpec{
					DataTypes:   "Http",
					IsConstruct: true,
				},
				Returns: &libast.FunctionReturnSpecList{
					&libast.FunctionReturnSpec{
						IsPointer: true,
						IsStruct:  true,
						LibName:   "",
						DataType:  "Http",
						Return:    "Http",
					},
				},
			},
		})
		abstractCode.AddFunctionArgs(libast.FunctionSpec{
			Name: "NewHttp",
			Args: libast.FunctionArgList{
				&libast.FunctionArg{
					Name:     "UserSvc",
					LibName:  "usersvc",
					DataType: "UserService",
				},
			},
		})

		err := abstractCode.RebuildCode()
		assert.NoError(t, err)
		act := abstractCode.GetCode()
		abstractCode = libast.NewAbstractCode(act, parser.ParseComments)

		abstractCode.AddFunctionArgsToReturn(libast.FunctionReturnArgsSpec{
			FuncName:      "NewHttp",
			ReturnName:    "Http",
			DataTypeKey:   "UserSvc",
			DataTypeValue: "UserSvc",
		})
		err = abstractCode.RebuildCode()
		assert.NoError(t, err)
		act = abstractCode.GetCode()
		// don't touch the expected code please
		// expect you change the test case please :D
		// building this string obviously difficult
		exp := `package test

type Http struct {
	UserSvc usersvc.UserService
}

func NewHttp(UserSvc usersvc.UserService) *Http {
	return &Http{UserSvc: UserSvc}
}

func main() {

}
`
		assert.Equal(t, exp, act)
	})

	t.Run("Test add function return 2", func(t *testing.T) {
		code := `
		package test

		type Http struct{
			UserSvc usersvc.UserService
		}

		func main() {
			
		}
		`

		abstractCode := libast.NewAbstractCode(code, parser.ParseComments)
		abstractCode.AddFunction(libast.FunctionSpecList{
			&libast.FunctionSpec{
				Name: "NewHttp",
				StructSpec: &libast.FunctionStructSpec{
					DataTypes:   "Http",
					IsConstruct: true,
				},
				Returns: &libast.FunctionReturnSpecList{
					&libast.FunctionReturnSpec{
						IsPointer: true,
						IsStruct:  true,
						LibName:   "",
						DataType:  "Http",
						Return:    "Http",
					},
				},
			},
		})
		abstractCode.AddFunctionArgs(libast.FunctionSpec{
			Name: "NewHttp",
			Args: libast.FunctionArgList{
				&libast.FunctionArg{
					Name:     "UserSvc",
					LibName:  "usersvc",
					DataType: "UserService",
				},
			},
		})

		err := abstractCode.RebuildCode()
		assert.NoError(t, err)
		act := abstractCode.GetCode()

		abstractCode = libast.NewAbstractCode(act, parser.ParseComments)

		abstractCode.AddFunctionArgsToReturn(libast.FunctionReturnArgsSpec{
			FuncName:      "NewHttp",
			ReturnName:    "Http",
			DataTypeKey:   "UserSvc",
			DataTypeValue: "UserSvc",
		})
		err = abstractCode.RebuildCode()
		assert.NoError(t, err)
		act = abstractCode.GetCode()

		abstractCode = libast.NewAbstractCode(act, parser.ParseComments)
		abstractCode.AddFunctionArgsToReturn(libast.FunctionReturnArgsSpec{
			FuncName:      "NewHttp",
			ReturnName:    "Http",
			DataTypeKey:   "AccountSvc",
			DataTypeValue: "AccountSvc",
		})
		err = abstractCode.RebuildCode()
		assert.NoError(t, err)
		act = abstractCode.GetCode()

		// don't touch the expected code please
		// expect you change the test case please :D
		// building this string obviously difficult
		exp := `package test

type Http struct {
	UserSvc usersvc.UserService
}

func NewHttp(UserSvc usersvc.UserService) *Http {
	return &Http{UserSvc: UserSvc, AccountSvc: AccountSvc}
}

func main() {

}
`
		assert.Equal(t, exp, act)
	})
}

func TestAddFunctionCaller(t *testing.T) {
	t.Run("Test add wire build from AddFunctionCaller", func(t *testing.T) {
		code := `
		package test
		`

		abstractCode := libast.NewAbstractCode(code, parser.ParseComments)
		abstractCode.AddFunction(libast.FunctionSpecList{
			&libast.FunctionSpec{
				Name: "InitHttpProtocol",
				Returns: &libast.FunctionReturnSpecList{
					&libast.FunctionReturnSpec{
						IsPointer: true,
						IsStruct:  true,
						LibName:   "http",
						DataType:  "HttpImpl",
						Return:    "HttpImpl",
					},
				},
			},
		})
		err := abstractCode.RebuildCode()
		assert.NoError(t, err)
		act := abstractCode.GetCode()

		abstractCode = libast.NewAbstractCode(act, parser.ParseComments)
		abstractCode.AddFunctionCaller("InitHttpProtocol", libast.CallerSpec{
			Func: libast.CallerFunc{
				Name: libast.CallerSelecterExpr{
					Name: "wire",
				},
				Selector: "Build",
			},
		})
		err = abstractCode.RebuildCode()
		assert.NoError(t, err)
		act = abstractCode.GetCode()
		// don't touch the expected code please
		// expect you change the test case please :D
		// building this string obviously difficult
		exp := `package test

func InitHttpProtocol() *http.HttpImpl {
	wire.Build()
	return &http.HttpImpl{}
}
`

		assert.Equal(t, exp, act)
	})
}

func TestAddArgsToCallExpr(t *testing.T) {
	t.Run("Test add args to callexpr 1", func(t *testing.T) {
		code := `
		package test
		`

		abstractCode := libast.NewAbstractCode(code, parser.ParseComments)
		abstractCode.AddFunction(libast.FunctionSpecList{
			&libast.FunctionSpec{
				Name: "InitHttpProtocol",
				Returns: &libast.FunctionReturnSpecList{
					&libast.FunctionReturnSpec{
						IsPointer: true,
						IsStruct:  true,
						LibName:   "http",
						DataType:  "HttpImpl",
						Return:    "HttpImpl",
					},
				},
			},
		})
		err := abstractCode.RebuildCode()
		assert.NoError(t, err)
		act := abstractCode.GetCode()

		abstractCode = libast.NewAbstractCode(act, parser.ParseComments)
		abstractCode.AddFunctionCaller("InitHttpProtocol", libast.CallerSpec{
			Func: libast.CallerFunc{
				Name: libast.CallerSelecterExpr{
					Name: "wire",
				},
				Selector: "Build",
			},
		})
		err = abstractCode.RebuildCode()
		assert.NoError(t, err)
		act = abstractCode.GetCode()

		abstractCode = libast.NewAbstractCode(act, parser.ParseComments)
		abstractCode.AddArgsToCallExpr(libast.CallerSpec{
			Func: libast.CallerFunc{
				Name: libast.CallerSelecterExpr{
					Name: "wire",
				},
				Selector: "Build",
			},
			Args: libast.CallerArgList{
				&libast.CallerArg{
					SelectorStmt: &libast.CallerArgSelectorStmt{
						LibName:  "db",
						DataType: "NewMysqlClient",
					},
				},

				&libast.CallerArg{
					Ident: &libast.CallerArgIdent{
						Name: "userSvc",
					},
				},
				&libast.CallerArg{
					Ident: &libast.CallerArgIdent{
						Name: "userRepo",
					},
				},
				&libast.CallerArg{
					SelectorStmt: &libast.CallerArgSelectorStmt{
						LibName:  "http",
						DataType: "NewHttpProtocol",
					},
				},
			},
		})
		err = abstractCode.RebuildCode()
		assert.NoError(t, err)
		act = abstractCode.GetCode()
		// don't touch the expected code please
		// expect you change the test case please :D
		// building this string obviously difficult
		exp := `package test

func InitHttpProtocol() *http.HttpImpl {
	wire.Build(db.NewMysqlClient, userSvc, userRepo, http.NewHttpProtocol)
	return &http.HttpImpl{}
}
`

		assert.Equal(t, exp, act)
	})
}

func TestAddWireDependencyInjection(t *testing.T) {
	t.Run("Test add wire dependency injection 1", func(t *testing.T) {
		code := `
		package test
		`

		abstractCode := libast.NewAbstractCode(code, parser.ParseComments)
		abstractCode.AddFunction(libast.FunctionSpecList{
			&libast.FunctionSpec{
				Name: "InitHttpProtocol",
				Returns: &libast.FunctionReturnSpecList{
					&libast.FunctionReturnSpec{
						IsPointer: true,
						IsStruct:  true,
						LibName:   "http",
						DataType:  "HttpImpl",
						Return:    "HttpImpl",
					},
				},
			},
		})
		err := abstractCode.RebuildCode()
		assert.NoError(t, err)
		act := abstractCode.GetCode()

		abstractCode = libast.NewAbstractCode(act, parser.ParseComments)
		abstractCode.AddFunctionCaller("InitHttpProtocol", libast.CallerSpec{
			Func: libast.CallerFunc{
				Name: libast.CallerSelecterExpr{
					Name: "wire",
				},
				Selector: "Build",
			},
		})
		err = abstractCode.RebuildCode()
		assert.NoError(t, err)
		act = abstractCode.GetCode()

		abstractCode = libast.NewAbstractCode(act, parser.ParseComments)
		abstractCode.AddWireDependencyInjection(libast.WireDependencyInjection{
			VarName:                   "productSvc",
			TargetInjectName:          "productsvc",
			TargetInjectConstructName: "NewProductService",
			InterfaceLib:              "productsvc",
			InterfaceName:             "ProductService",
			StructLib:                 "productsvc",
			StructName:                "ProductServiceImpl",
		})
		err = abstractCode.RebuildCode()
		assert.NoError(t, err)
		act = abstractCode.GetCode()

		// don't touch the expected code please
		// expect you change the test case please :D
		// building this string obviously difficult
		exp := `package test

var productSvc = wire.NewSet(productsvc.NewProductService, wire.Bind(new(productsvc.ProductService), new(*productsvc.ProductServiceImpl)))

func InitHttpProtocol() *http.HttpImpl {
	wire.Build()
	return &http.HttpImpl{}
}
`
		assert.Equal(t, exp, act)
	})

	t.Run("Test add wire dependency injection 2", func(t *testing.T) {
		code := `//+build wireinject
		package test
		`

		abstractCode := libast.NewAbstractCode(code, parser.ParseComments)
		abstractCode.AddFunction(libast.FunctionSpecList{
			&libast.FunctionSpec{
				Name: "InitHttpProtocol",
				Returns: &libast.FunctionReturnSpecList{
					&libast.FunctionReturnSpec{
						IsPointer: true,
						IsStruct:  true,
						LibName:   "http",
						DataType:  "HttpImpl",
						Return:    "HttpImpl",
					},
				},
			},
		})
		err := abstractCode.RebuildCode()
		assert.NoError(t, err)
		act := abstractCode.GetCode()

		abstractCode = libast.NewAbstractCode(act, parser.ParseComments)
		abstractCode.AddFunctionCaller("InitHttpProtocol", libast.CallerSpec{
			Func: libast.CallerFunc{
				Name: libast.CallerSelecterExpr{
					Name: "wire",
				},
				Selector: "Build",
			},
		})
		err = abstractCode.RebuildCode()
		assert.NoError(t, err)
		act = abstractCode.GetCode()

		abstractCode = libast.NewAbstractCode(act, parser.ParseComments)
		abstractCode.AddWireDependencyInjection(libast.WireDependencyInjection{
			VarName:                   "productSvc",
			TargetInjectName:          "productsvc",
			TargetInjectConstructName: "NewProductService",
			InterfaceLib:              "productsvc",
			InterfaceName:             "ProductService",
			StructLib:                 "productsvc",
			StructName:                "ProductServiceImpl",
		})
		err = abstractCode.RebuildCode()
		assert.NoError(t, err)
		act = abstractCode.GetCode()

		abstractCode = libast.NewAbstractCode(act, parser.ParseComments)
		abstractCode.AddArgsToCallExpr(
			libast.CallerSpec{
				Func: libast.CallerFunc{
					Name: libast.CallerSelecterExpr{
						Name: "wire",
					},
					Selector: "Build",
				},
				Args: libast.CallerArgList{
					&libast.CallerArg{
						Ident: &libast.CallerArgIdent{
							Name: "productSvc",
						},
					},
				},
			},
		)
		err = abstractCode.RebuildCode()
		assert.NoError(t, err)
		act = abstractCode.GetCode()

		// don't touch the expected code please
		// expect you change the test case please :D
		// building this string obviously difficult
		exp := `//+build wireinject
package test

var productSvc = wire.NewSet(productsvc.NewProductService, wire.Bind(new(productsvc.ProductService), new(*productsvc.ProductServiceImpl)))

func InitHttpProtocol() *http.HttpImpl {
	wire.Build(productSvc)
	return &http.HttpImpl{}
}
`
		assert.Equal(t, exp, act)
	})

	t.Run("Test add wire dependency injection 3", func(t *testing.T) {
		code := `
		package test
		`

		abstractCode := libast.NewAbstractCode(code, parser.ParseComments)
		abstractCode.AddFunction(libast.FunctionSpecList{
			&libast.FunctionSpec{
				Name: "InitHttpProtocol",
				Returns: &libast.FunctionReturnSpecList{
					&libast.FunctionReturnSpec{
						IsPointer: true,
						IsStruct:  true,
						LibName:   "http",
						DataType:  "HttpImpl",
						Return:    "HttpImpl",
					},
				},
			},
		})
		err := abstractCode.RebuildCode()
		assert.NoError(t, err)
		act := abstractCode.GetCode()

		abstractCode = libast.NewAbstractCode(act, parser.ParseComments)
		abstractCode.AddFunctionCaller("InitHttpProtocol", libast.CallerSpec{
			Func: libast.CallerFunc{
				Name: libast.CallerSelecterExpr{
					Name: "wire",
				},
				Selector: "Build",
			},
		})
		err = abstractCode.RebuildCode()
		assert.NoError(t, err)
		act = abstractCode.GetCode()

		abstractCode = libast.NewAbstractCode(act, parser.ParseComments)
		abstractCode.AddWireDependencyInjection(libast.WireDependencyInjection{
			VarName:                   "productSvc",
			TargetInjectName:          "productsvc",
			TargetInjectConstructName: "NewProductService",
		})
		err = abstractCode.RebuildCode()
		assert.NoError(t, err)
		act = abstractCode.GetCode()

		abstractCode = libast.NewAbstractCode(act, parser.ParseComments)
		abstractCode.AddArgsToCallExpr(
			libast.CallerSpec{
				Func: libast.CallerFunc{
					Name: libast.CallerSelecterExpr{
						Name: "wire",
					},
					Selector: "Build",
				},
				Args: libast.CallerArgList{
					&libast.CallerArg{
						Ident: &libast.CallerArgIdent{
							Name: "productSvc",
						},
					},
				},
			},
		)
		err = abstractCode.RebuildCode()
		assert.NoError(t, err)
		act = abstractCode.GetCode()

		// don't touch the expected code please
		// expect you change the test case please :D
		// building this string obviously difficult
		exp := `package test

var productSvc = wire.NewSet(productsvc.NewProductService)

func InitHttpProtocol() *http.HttpImpl {
	wire.Build(productSvc)
	return &http.HttpImpl{}
}
`
		assert.Equal(t, exp, act)
	})

	t.Run("Test add wire dependency injection 4", func(t *testing.T) {
		code := `
		package test
		`

		abstractCode := libast.NewAbstractCode(code, parser.ParseComments)
		abstractCode.AddFunction(libast.FunctionSpecList{
			&libast.FunctionSpec{
				Name: "InitHttpProtocol",
				Returns: &libast.FunctionReturnSpecList{
					&libast.FunctionReturnSpec{
						IsPointer: true,
						IsStruct:  true,
						LibName:   "http",
						DataType:  "HttpImpl",
						Return:    "HttpImpl",
					},
				},
			},
		})
		err := abstractCode.RebuildCode()
		assert.NoError(t, err)
		act := abstractCode.GetCode()

		abstractCode = libast.NewAbstractCode(act, parser.ParseComments)
		abstractCode.AddFunctionCaller("InitHttpProtocol", libast.CallerSpec{
			Func: libast.CallerFunc{
				Name: libast.CallerSelecterExpr{
					Name: "wire",
				},
				Selector: "Build",
			},
		})
		err = abstractCode.RebuildCode()
		assert.NoError(t, err)
		act = abstractCode.GetCode()

		abstractCode = libast.NewAbstractCode(act, parser.ParseComments)
		abstractCode.AddWireDependencyInjection(libast.WireDependencyInjection{
			VarName:                   "productSvc",
			TargetInjectName:          "productsvc",
			TargetInjectConstructName: "NewProductService",
		})
		err = abstractCode.RebuildCode()
		assert.NoError(t, err)
		act = abstractCode.GetCode()

		abstractCode = libast.NewAbstractCode(act, parser.ParseComments)
		abstractCode.AddArgsToCallExpr(
			libast.CallerSpec{
				Func: libast.CallerFunc{
					Name: libast.CallerSelecterExpr{
						Name: "wire",
					},
					Selector: "Build",
				},
				Args: libast.CallerArgList{
					&libast.CallerArg{
						Ident: &libast.CallerArgIdent{
							Name: "productSvc",
						},
					},
				},
			},
		)
		err = abstractCode.RebuildCode()
		assert.NoError(t, err)
		act = abstractCode.GetCode()

		// append new wire
		abstractCode = libast.NewAbstractCode(act, parser.ParseComments)
		abstractCode.AddWireDependencyInjection(libast.WireDependencyInjection{
			VarName:                   "userSvc",
			TargetInjectName:          "usersvc",
			TargetInjectConstructName: "NewUserService",
		})
		err = abstractCode.RebuildCode()
		assert.NoError(t, err)
		act = abstractCode.GetCode()

		abstractCode = libast.NewAbstractCode(act, parser.ParseComments)
		abstractCode.AddArgsToCallExpr(
			libast.CallerSpec{
				Func: libast.CallerFunc{
					Name: libast.CallerSelecterExpr{
						Name: "wire",
					},
					Selector: "Build",
				},
				Args: libast.CallerArgList{
					&libast.CallerArg{
						Ident: &libast.CallerArgIdent{
							Name: "userSvc",
						},
					},
				},
			},
		)
		err = abstractCode.RebuildCode()
		assert.NoError(t, err)
		act = abstractCode.GetCode()

		// don't touch the expected code please
		// expect you change the test case please :D
		// building this string obviously difficult
		exp := `package test

var userSvc = wire.NewSet(usersvc.NewUserService)

var productSvc = wire.NewSet(productsvc.NewProductService)

func InitHttpProtocol() *http.HttpImpl {
	wire.Build(productSvc, userSvc)
	return &http.HttpImpl{}
}
`
		assert.Equal(t, exp, act)
	})

	t.Run("Test add wire dependency injection 5", func(t *testing.T) {
		code := `//+build wireinject
package main

import (
	db "test/infrastructure/database"
	"test/internal/protocol/http"
	httprouter "test/internal/protocol/http/router"
	httphandler "test/src/handler/http"

	"github.com/google/wire"
)

var storages = wire.NewSet(db.NewMysqlClient)

var httpHandler = wire.NewSet(httphandler.NewHttpHandler)

var httpRouter = wire.NewSet(httprouter.NewHttpRouter)

func InitHttpProtocol() *http.HttpImpl {
	wire.Build(storages, httpHandler, httpRouter, http.NewHttp)
	return &http.HttpImpl{}
}

		`

		abstractCode := libast.NewAbstractCode(code, parser.ParseComments)
		abstractCode.AddWireDependencyInjection(
			libast.WireDependencyInjection{
				VarName:                   "productSvc",
				TargetInjectName:          "productsvc",
				TargetInjectConstructName: "NewProductService",
				InterfaceLib:              "productsvc",
				InterfaceName:             "ProductService",
				StructLib:                 "productsvc",
				StructName:                "ProductServiceImpl",
			},
		)
		abstractCode.AddArgsToCallExpr(
			ast.CallerSpec{
				Func: ast.CallerFunc{
					Name: ast.CallerSelecterExpr{
						Name: "wire",
					},
					Selector: "Build",
				},
				Args: ast.CallerArgList{
					&ast.CallerArg{
						Ident: &ast.CallerArgIdent{
							Name: "productSvc",
						},
					},
				},
			},
		)

		err := abstractCode.RebuildCode()
		assert.NoError(t, err)
		act := abstractCode.GetCode()
		// don't touch the expected code please
		// expect you change the test case please :D
		// building this string obviously difficult
		exp := `//+build wireinject
package main

import (
	db "test/infrastructure/database"
	"test/internal/protocol/http"
	httprouter "test/internal/protocol/http/router"
	httphandler "test/src/handler/http"

	"github.com/google/wire"
)

var productSvc = wire.NewSet(productsvc.NewProductService, wire.Bind(new(productsvc.ProductService), new(*productsvc.ProductServiceImpl)))

var storages = wire.NewSet(db.NewMysqlClient)

var httpHandler = wire.NewSet(httphandler.NewHttpHandler)

var httpRouter = wire.NewSet(httprouter.NewHttpRouter)

func InitHttpProtocol() *http.HttpImpl {
	wire.Build(storages, httpHandler, httpRouter, http.NewHttp, productSvc)
	return &http.HttpImpl{}
}
		
`

		assert.Equal(t, exp, act)
	})

}

func TestAddCommentBeforeFunction(t *testing.T) {
	t.Run("Test add comment before function 1", func(t *testing.T) {
		code := `
		package test

		func main() {}
		`

		abstractCode := libast.NewAbstractCode(code, parser.ParseComments)
		abstractCode.AddCommentOutsideFunction(libast.Comment{
			Value: "//go:generate go run github.com/google/wire/cmd/wire",
		})
		err := abstractCode.RebuildCode()
		assert.NoError(t, err)
		act := abstractCode.GetCode()

		// don't touch the expected code please
		// expect you change the test case please :D
		// building this string obviously difficult
		exp := `//go:generate go run github.com/google/wire/cmd/wire
package test

func main()	{}
`
		assert.Equal(t, exp, act)
	})

	t.Run("Test add comment before function 2", func(t *testing.T) {
		code := `
		package test

		func main() {}
		`

		abstractCode := libast.NewAbstractCode(code, parser.ParseComments)
		abstractCode.AddCommentOutsideFunction(libast.Comment{
			FunctionName: "main",
			Value:        "// Main function",
		})
		err := abstractCode.RebuildCode()
		assert.NoError(t, err)
		act := abstractCode.GetCode()

		// don't touch the expected code please
		// expect you change the test case please :D
		// building this string obviously difficult
		exp := `package test

// Main function
func main()	{}
`
		assert.Equal(t, exp, act)
	})

	t.Run("Test add comment before function 3", func(t *testing.T) {
		code := `
		package test

		import (
			"fmt"
			"strings"
		)

		func main() {
			// test
		}
		`

		abstractCode := libast.NewAbstractCode(code, parser.ParseComments)
		abstractCode.AddCommentOutsideFunction(libast.Comment{
			Value: "//go:generate go run github.com/google/wire/cmd/wire",
		})
		err := abstractCode.RebuildCode()
		assert.NoError(t, err)
		act := abstractCode.GetCode()

		// don't touch the expected code please
		// expect you change the test case please :D
		// building this string obviously difficult
		exp := `//go:generate go run github.com/google/wire/cmd/wire
package test

import (
	"fmt"
	"strings"
)

func main() {
	// test
}
`
		assert.Equal(t, exp, act)
	})

	t.Run("Test add comment before function 4", func(t *testing.T) {
		code := `
		package test

		import (
			"fmt"
			"strings"
		)

		func main() {
			// test
		}
		`

		abstractCode := libast.NewAbstractCode(code, parser.ParseComments)
		abstractCode.AddCommentOutsideFunction(libast.Comment{
			Value: "//go:generate go run github.com/google/wire/cmd/wire",
		})
		abstractCode.AddCommentOutsideFunction(libast.Comment{
			Value: "//go:generate go run github.com/google/wire/cmd/wire init",
		})
		err := abstractCode.RebuildCode()
		assert.NoError(t, err)
		act := abstractCode.GetCode()

		// don't touch the expected code please
		// expect you change the test case please :D
		// building this string obviously difficult
		exp := `//go:generate go run github.com/google/wire/cmd/wire init
//go:generate go run github.com/google/wire/cmd/wire
package test

import (
	"fmt"
	"strings"
)

func main() {
	// test
}
`
		assert.Equal(t, exp, act)
	})

	t.Run("Test add comment before function 5", func(t *testing.T) {
		code := `
		package test

		import (
			"github.com/google/wire"
			db "test1/infrastructure/database"
			"test1/internal/protocol/http"
			httprouter "test1/internal/protocol/http/router"
			httphandler "test1/src/handler/http"
		)
		
		var storages = wire.NewSet(db.NewMysqlClient)
		
		var httpHandler = wire.NewSet(httphandler.NewHttpHandler)
		
		var httpRouter = wire.NewSet(httprouter.NewHttpRouter)
		
		func InitHttpProtocol() *http.HttpImpl {
			wire.Build(storages, httpHandler, httpRouter)
			return &http.HttpImpl{}
		}	
		`

		abstractCode := libast.NewAbstractCode(code, parser.ParseComments)
		abstractCode.AddCommentOutsideFunction(libast.Comment{
			Value: "//+build wireinject",
		})
		err := abstractCode.RebuildCode()
		assert.NoError(t, err)
		act := abstractCode.GetCode()

		// don't touch the expected code please
		// expect you change the test case please :D
		// building this string obviously difficult
		exp := `//+build wireinject
package test

import (
	"github.com/google/wire"
	db "test1/infrastructure/database"
	"test1/internal/protocol/http"
	httprouter "test1/internal/protocol/http/router"
	httphandler "test1/src/handler/http"
)

var storages = wire.NewSet(db.NewMysqlClient)

var httpHandler = wire.NewSet(httphandler.NewHttpHandler)

var httpRouter = wire.NewSet(httprouter.NewHttpRouter)

func InitHttpProtocol() *http.HttpImpl {
	wire.Build(storages, httpHandler, httpRouter)
	return &http.HttpImpl{}
}
`
		assert.Equal(t, exp, act)
	})

}
