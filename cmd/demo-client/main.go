package main

import (
	"github.com/nmiculinic/metricsd/pkg/metricsd"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"
)

func main() {
	addr := pflag.String("addr", "[::]:8080", "service listening address")
	gaddr := pflag.String("gaddr", "localhost:5555", "grpc service address")
	pflag.Parse()
	metricsd.Serve(
		*addr,
		*gaddr,
		metricsd.DefaultHtmlStringer,
		grpc.WithInsecure(),
	)
}
