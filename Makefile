gen:
	protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=. --grpc-gateway_out=.
server:
	go run cmd/server/main.go
client:
	go run cmd/client/main.go -address 0.0.0.0:8085