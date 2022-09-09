module github.com/nekitvand/to_do_service

go 1.19

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.11.3
	google.golang.org/genproto v0.0.0-20220908141613-51c1cc9bc6d0
	google.golang.org/grpc v1.49.0
	google.golang.org/protobuf v1.28.1
)

require gopkg.in/yaml.v3 v3.0.1 // indirect

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/pkg/errors v0.9.1
	golang.org/x/net v0.0.0-20220907135653-1e95f45603a7 // indirect
	golang.org/x/sys v0.0.0-20220908150016-7ac13a9a928d // indirect
	golang.org/x/text v0.3.7 // indirect
)

replace github.com/nekitvand/to_do_service => ./pkg/to_do_service
