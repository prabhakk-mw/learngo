module github.com/prabhakk-mw/learngo/mw/gateway

go 1.24.1

replace github.com/prabhakk-mw/learngo/mw/services/capitalize => ../services/capitalize

replace github.com/prabhakk-mw/learngo/mw/common => ../common

require (
	github.com/prabhakk-mw/learngo/mw/common v0.0.0-00010101000000-000000000000
	github.com/prabhakk-mw/learngo/mw/services/capitalize v0.0.0-00010101000000-000000000000
	github.com/swaggo/http-swagger/v2 v2.0.2
	github.com/swaggo/swag v1.16.4
	google.golang.org/grpc v1.73.0
)

require (
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/go-openapi/jsonpointer v0.21.1 // indirect
	github.com/go-openapi/jsonreference v0.21.0 // indirect
	github.com/go-openapi/spec v0.21.0 // indirect
	github.com/go-openapi/swag v0.23.1 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.9.0 // indirect
	github.com/swaggo/files v1.0.1 // indirect
	github.com/swaggo/files/v2 v2.0.0 // indirect
	github.com/swaggo/http-swagger v1.3.4 // indirect
	golang.org/x/net v0.41.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.26.0 // indirect
	golang.org/x/tools v0.34.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250324211829-b45e905df463 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
