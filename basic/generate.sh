cd api/prod/proto
protoc --micro_out=.. --go_out=.. prods.proto
protoc-go-inject-tag --input=../prods.pb.go
