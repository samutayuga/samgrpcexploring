
protoc -I . --grpc-gateway_out . \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
    --openapiv2_out . \
    --openapiv2_opt logtostderr=true \
    --openapiv2_opt use_go_templates=true \
    blog/blogpb/blog_service.proto