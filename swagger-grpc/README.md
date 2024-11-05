# Swagger-GRPC

Welcome to the implementation of OpenAPI(Swagger) with grpc ecosystem

### How to run

1. `go mod tidy`
2. `go run main.go`
3. open [localhost:50051/sw](http://localhost:50051/sw)

### Here is the step-by-step guide on how to make it yourself

1. Install [protoc](https://grpc.io/docs/protoc-installation/); [buf](https://buf.build/docs/installation/) and add it to `$PATH`
2. Run `go mod init {your_module_name}` and [install required packages](https://gist.github.com/dongxuny/964d6181ed06c1b35943311f5b761286#file-buf-requirement-md)
3. Write your proto files in the same format as [here](proto/product/product.proto)/[here](proto/user/user.proto)
4. Create `buf.yaml` & `buf.gen.yaml` files
5. Fill `buf.yaml` & `buf.gen.yaml` in the similar format
6. Run `buf mod update` && `buf generate`. This should've generated go files and single `*.swagger.json`
7. Create `boot.yaml` and fill it with necessary elements
8. Create `main.go` in root directory and write in it in the similar format
9. Run `go run main.go` and open [localhost:50051/sw](http://localhost:50051/sw)
10. Thats it !

### Useful links:
- [Medium article](https://medium.com/@pointgoal/grpc-how-to-add-swagger-ui-on-grpc-466e5fd71097)
- [Adding annotations](https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/adding_annotations/)
- [Customizing OpenAPI output](https://grpc-ecosystem.github.io/grpc-gateway/docs/mapping/customizing_openapi_output/)