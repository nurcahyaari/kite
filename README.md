# What it is

Kite is a Golang Code Generation inspired by Clean Code Architecture

# Installation

Clone this repo
```bash
git clone https://github.com/nurcahyaari/Kite.git
```

then open the directory in your terminal app.
```bash
go build kite.go
```
The app was built as binary. then move the binary app into your $PATH environment

# How to use
This following bash script is the example to generate project using Kite


<table border="1">
    <tr>
        <th>Command</th>
        <th>Function</th>
    </tr>
    <tr>
        <td>
            kite new --name <your app name>
        </td>
        <td>Create new application/ project/ service</td>
    </tr>
    <tr>
        <td>main module --name <your module name></td>
        <td>creating new modules such as, product modules, user modules, auth modules, etc</td>
    </tr>
</table>

## creating new applications

after you have done to setup your binary of kite in your environtment $PATH you can use kite to create your new applications

command
```bash
kite new --args1 value --args2 value
```

<table border="1">
    <tr>
        <th>args</th>
        <th>function</th>
    </tr>
    <tr>
        <td>name</td>
        <td>your application name</td>
    </tr>
    <tr>
        <td>path</td>
        <td>your application path</td>
    </tr>
</table>

example, I created new Backend service for my apps called as `Todo-Backend`. and I want to save this service at my Documents package `/usr/Documents`

1. First I can use command `cd` to go to `/usr/Documents`
2. Or I can use the path args in kite generation
command
```bash
kite new --name Todo-Backend --path /usr/Documents
```

After I've created my Backend I must create my Modules first


## creating your new modules

because when you create your new apps, currently the generator will not create your first module. you must to create your modules to determines that your service working

command
```bash
kite module --args1 value --args2 value
```

<table border="1">
    <tr>
        <th>args</th>
        <th>function</th>
    </tr>
    <tr>
        <td>name</td>
        <td>your module name</td>
    </tr>
    <tr>
        <td>path</td>
        <td>your application path</td>
    </tr>
</table>

The module name is `Test` and it is under of my `Todo-Backend` so I will use this command
```bash
kite module --name Test --path --value /usr/Documents/Todo-Backend
```
because, Kite still doesn't support to auto adding modules into wire.go. after adding the new modules you must manually inject your modules and any others dependencies into wire.go manually.

after you create your new project, and your new module then you must:
1. add your handlers into internal/protocols/http/router/route.go
2. add package src/handlers/http into protocols/http/http.go, it should be look like this
```go
package router

import (
	"github.com/go-chi/chi/v5"
	"Todo-Backend/src/handlers/http"
)

type HttpRouterImpl struct {
	httpHandlers http.HttpHandler
}

func NewHttpRouter(httpHandlers http.HttpHandler) *HttpRouterImpl {
	return &HttpRouterImpl{
		httpHandlers: httpHandlers,
	}
}
func (h *HttpRouterImpl) Router(r *chi.Mux) {
	h.httpHandlers.Router(r)
}

```

4. add your new modules into wire.go

first you must add into your import package

add your test module's repo `testrepo "Todo-Backend/src/modules/test/repository"`

add your test module's service `testsvc "Todo-Backend/src/modules/test/service"`

```go
var testRepo = wire.NewSet(
	testrepo.NewTestRepository,
	wire.Bind(
		new(testrepo.TestRepository),
		new(*testrepo.TestRepositoryImpl),
	),
)

var testSvc = wire.NewSet(
	testsvc.NewTestService,
	wire.Bind(
		new(testsvc.TestService),
		new(*testsvc.TestServiceImpl),
	),
)
```
and add into domain variables in wire.go
```go
var domain = wire.NewSet(
	testSvc,
)
```

5. run `go generate ./...`
6. copy your .env.example into .env
7. if you've got an error, install the package that needed in your project
8. run your project with this command `go run main.go wire_gen.go`