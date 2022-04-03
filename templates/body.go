package templates

import _ "embed"

type TemplateBody string

//go:embed package_template.tmpl
var PackageTemplate TemplateBody

//go:embed config.tmpl
var ConfigTemplate TemplateBody

//go:embed routes.tmpl
var RouteTemplate TemplateBody

//go:embed env.tmpl
var EnvTemplate TemplateBody

//go:embed wire.tmpl
var WireTemplate TemplateBody

//go:embed gitignore.tmpl
var GitignoreTemplate TemplateBody

// Infra

var (
	//go:embed infra/db.tmpl
	DBTemplate TemplateBody
)

// Utils

// encpr
var (
	//go:embed utils/encrp/rsa.tmpl
	RSABaseTemplate TemplateBody
	//go:embed utils/encrp/read_private_key.tmpl
	ReadPrivateKeyTemplate TemplateBody
	//go:embed utils/encrp/read_public_key.tmpl
	ReadPublicKeyTemplate TemplateBody
)

// Logger
var (
	//go:embed utils/logger/log.tmpl
	Loggertemplate TemplateBody
)

//go:embed dependency.tmpl
var DependencyTemplate TemplateBody

// Protocols HTTP using Chi
var (
	//go:embed protocols/http/chi/err.tmpl
	ProtocolHttpChiErrorTemplate TemplateBody
	//go:embed protocols/http/chi/response.tmpl
	ProtocolHttpChiResponseTemplate TemplateBody
	//go:embed protocols/http/chi/middleware.tmpl
	ProtocolHttpChiMiddlewareTemplate TemplateBody
)
