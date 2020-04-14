package main

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"main.yapo.fun/servicegw"
)

func main() {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	endpoint := "127.0.0.1:9090"
	servicegw.RegisterTestServiceHandlerFromEndpoint(context.TODO(), mux, endpoint, opts)
	http.ListenAndServe(":9002", mux)
}
