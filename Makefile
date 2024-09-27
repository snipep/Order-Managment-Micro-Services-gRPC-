# Run the orders service
run-orders:
	@cd services/orders && go run .

# Run the kitchen service
run-kitchen:
	@cd services/kitchen && go run .

# Generate gRPC stubs
gen:
	@protoc \
		--proto_path=protobuf "protobuf/orders.proto" \
		--go_out=services/common/genproto/orders --go_opt=paths=source_relative \
		--go-grpc_out=services/common/genproto/orders --go-grpc_opt=paths=source_relative
