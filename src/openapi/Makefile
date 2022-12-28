.PHONY: imaging-study
imaging-study:
	oapi-codegen -generate "server" -package openapi imaging-study.yaml > ../services/imaging-study/openapi/server.gen.go
	oapi-codegen -generate "types" -package openapi imaging-study.yaml > ../services/imaging-study/openapi/type.gen.go
	oapi-codegen -generate "spec" -package openapi imaging-study.yaml > ../services/imaging-study/openapi/spec.gen.go
