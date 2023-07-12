module example.com/example-app

go 1.20

require (
	example.com/server v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.12.0
	google.golang.org/grpc v1.56.2
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/svarlamov/goyhfin v0.0.0-20161220065822-c7565afb5e91 // indirect
	golang.org/x/sys v0.10.0 // indirect
	golang.org/x/text v0.11.0 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
)

replace example.com/server => ./server
