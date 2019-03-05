package main

import (
	_ "github.com/lib/pq"
	"github.com/nmiculinic/metricsd/pkg/backends/sql"
	"github.com/nmiculinic/metricsd/pkg/metricsd"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/xo/dburl"
	"google.golang.org/grpc"
	"net"
)

func main() {
	addr := pflag.String("addr", "[::]:5555", "service listening address")
	dbURL := pflag.String("dburl", "", "backend database url. See https://github.com/xo/dburl for details")
	pflag.Parse()

	u, err := dburl.Parse(*dbURL)
	if err != nil {
		logrus.WithError(err).Fatal("cannot parse database string")
	}
	logrus.
		WithField("database", u.Path).
		WithField("host", u.Host).
		WithField("username", u.User.Username()).
		Infoln("Trying to connect to")
	db, err := dburl.Open(*dbURL)
	if err != nil {
		logrus.WithError(err).Fatal("cannot open database")
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		logrus.WithError(err).Fatal("cannot ping database")
	}

	logrus.WithField("database", u.Path).WithField("host", u.Host).Infoln("successfully connected to database")

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		logrus.WithError(err).Fatalf("failed to listen")
	}
	logrus.Infof("listening on %v", lis.Addr())

	s := grpc.NewServer()
	metricsd.RegisterMetricsServiceServer(s, &sql.Metricsd{DB: db})
	if err := s.Serve(lis); err != nil {
		logrus.WithError(err).Panicln("failed to serve")
	}
}
