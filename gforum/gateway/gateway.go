package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	gw "gforum/grpc/gen/proto/forum/v1"
)

var (
	echoEndpoint = flag.String("endpoint", "localhost:50051", "endpoint of YourService")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	ropts := []runtime.ServeMuxOption{
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}),
	}

	mux := runtime.NewServeMux(ropts...)
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// users
	err := gw.RegisterUsersHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
	if err != nil {
		return err
	}

	// ariticles
	err = gw.RegisterArticlesHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
	if err != nil {
		return err
	}

	log.Println("starting gateway server on port 3000")
	return http.ListenAndServe(":3000", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
