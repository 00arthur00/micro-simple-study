cd service/proto
protoc --micro_out=.. --go_out=.. test.proto
protoc-go-inject-tag --input=../test.pb.go

protoc -I. --go_out=plugins=grpc:../../servicegw test.proto
protoc -I. --grpc-gateway_out=logtostderr=true:../../servicegw test.proto