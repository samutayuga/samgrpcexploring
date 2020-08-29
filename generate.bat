REM sprotoc -I simple\ --go_out=simple\ simple\simple.proto
REM protoc -I enumsample\ --go_out=enumsample\ enumsample\enum_sample.proto

REM protoc -I complex\ --go_out=complex\ complex\complexsample.proto
REM protoc -I addressbook\ --go_out=addressbook\ addressbook\add_book.proto

REM protoc greet\greetpb\greet.proto --go_out=plugins=grpc:.

REM protoc calculator\calculatorpb\calculator.proto --go_out=plugins=grpc:.

REM protoc --proto_path=C:\apps\protoc-3.11.4-win64\include -I. -I%GOPATH%\src\github.com\grpc-ecosystem\grpc-gateway\third_party\googleapis --go_out=plugins=grpc:. blog\blogpb\blog.proto

REM protoc --proto_path=C:\apps\protoc-3.11.4-win64\include -I. -I%GOPATH%\src\github.com\grpc-ecosystem\grpc-gateway\third_party\googleapis --grpc-gateway_out=logtostderr=true,paths=source_relative:. blog\blogpb\blog.proto

protoc -I. -I%GOPATH%\src\github.com\grpc-ecosystem\grpc-gateway\third_party\googleapis --swagger_out=logtostderr=true:. blog\blogpb\blog.proto