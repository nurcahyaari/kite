# What it is

Kite is a Golang Code Generation inspired by Clean Code Architecture

# Installation

Clone this repo
```bash
git clone https://github.com/nurcahyaari/Kite.git
```

then open the directory in your terminal app.
```bash
go build main.go
```
The app was built as binary. then move the binary app into your $PATH environment

# How To use
This following bash script is the example to generate project using Kite

creating new project
```bash
main new --name <your app name>
```

creating new module
```bash
main module --name <your module name>
```

currently Kite doesn't support auto adding modules into wire.go. after adding the new modules you must manually inject your modules and any others dependencies into wire.go manually.

after you create your new project, and your new module then you must:
1. add your handlers into protocols/http/router/route.go
2. add your router/route.go into protocols/http/http.go
3. run go generate ./...
4. if you've got an error, install the package that needed in your project
5. copy your .env.example into .env
6. run your project again