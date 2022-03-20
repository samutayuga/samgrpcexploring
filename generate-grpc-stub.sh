
protoc -I . --grpc-gateway_out . \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
    blog/blogpb/blog_service.proto

protoc -I. --swagger_out=logtostderr=true blog/blogpb/*.proto