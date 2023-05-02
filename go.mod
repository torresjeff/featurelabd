module github.com/torresjeff/featurelabd

go 1.20

require (
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/torresjeff/go-feature-lab v0.0.0-20230501214049-3a13b19ff20b
	google.golang.org/grpc v1.54.0
	google.golang.org/protobuf v1.30.0
)

require (
	github.com/go-resty/resty/v2 v2.7.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
)

//replace github.com/torresjeff/go-feature-lab => ../go-feature-lab
