package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	pb "gitlab.com/neven-miculinic/metricsd/pkg/metricsd"
	"gitlab.com/neven-miculinic/metricsd/pkg/sql"
	"google.golang.org/grpc"
	"net"
)

func main() {
	addr := pflag.String("addr", "[::]:8080", "service listening address")
	pflag.Parse()
	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		logrus.WithError(err).Fatalf("failed to listen")
	}
	logrus.Infof("listening on %v", lis.Addr())

	s := grpc.NewServer()
	pb.RegisterMetricsServiceServer(s, &sql.SQLBackend{})
	if err := s.Serve(lis); err != nil {
		logrus.WithError(err).Fatalln("failed to serve")
	}
}
