REM sprotoc -I simple\ --go_out=simple\ simple\simple.proto
REM protoc -I enumsample\ --go_out=enumsample\ enumsample\enum_sample.proto

REM protoc -I complex\ --go_out=complex\ complex\complexsample.proto
REM protoc -I addressbook\ --go_out=addressbook\ addressbook\add_book.proto

protoc greet\greetpb\greet.proto --go_out=plugins=grpc:.

protoc calculator\calculatorpb\calculator.proto --go_out=plugins=grpc:.

protoc blog\blogpb\blog.proto --go_out=plugins=grpc:.