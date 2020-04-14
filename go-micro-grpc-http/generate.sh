cd service/proto
protoc --micro_out=.. --go_out=.. *.proto
protoc-go-inject-tag --input=../models.pb.go
