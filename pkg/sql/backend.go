package sql

import (
	"context"
	"database/sql"
	"github.com/sirupsen/logrus"
	"gitlab.com/neven-miculinic/metricsd/pkg/metricsd"
)

type SQLBackend struct {
	DB *sql.DB
}

func (*SQLBackend) ReportProcessMeasurement(ctx context.Context, m *metricsd.ProcessMeasurement) (*metricsd.Empty, error) {
	logrus.Infof("Called with %v", m)
	return &metricsd.Empty{}, nil
}

func (*SQLBackend) ReportNodeMeasurement(ctx context.Context, m *metricsd.NodeMeasurement) (*metricsd.Empty, error) {
	logrus.Infof("Called with %v", m)
	return &metricsd.Empty{}, nil
}
