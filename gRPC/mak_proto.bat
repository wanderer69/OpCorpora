protoc -I. --go_out=plugins=grpc:. opcorpora.proto
rem protoc -I. -I%GOPATH%\src -IC:/Development/Go_pkg/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis -IC:/Development/Go_pkg/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0 --go_out=plugins=grpc:. auth.proto
rem protoc -I. -I%GOPATH%\src -IC:/Development/Go_pkg/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis -IC:/Development/Go_pkg/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0 --grpc-gateway_out=logtostderr=true:. bookstore.proto
rem protoc -I. -I%GOPATH%\src -IC:/Development/Go_pkg/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis -IC:/Development/Go_pkg/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0 -I%GOPATH%\src\github.com\grpc-ecosystem\grpc-gateway --swagger_out=logtostderr=true:. http_bookstore.proto
