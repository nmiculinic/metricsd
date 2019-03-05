package sql

import (
	"context"
	"database/sql"
	"github.com/nmiculinic/metricsd/pkg/metricsd"
)

var _ metricsd.MetricsServiceServer = (*SQLBackend)(nil)

type SQLBackend struct {
	DB *sql.DB
}

func (*SQLBackend) ReportProcessMeasurement(context.Context, *metricsd.ProcessMeasurement) (*metricsd.Empty, error) {
	panic("implement me")
}

func (*SQLBackend) ReportNodeMeasurement(context.Context, *metricsd.NodeMeasurement) (*metricsd.Empty, error) {
	panic("implement me")
}

func (*SQLBackend) NodeAverages(context.Context, *metricsd.NodeAvgReq) (*metricsd.NodeAvgResp, error) {
	panic("implement me")
}

func (*SQLBackend) ProcessesAverages(context.Context, *metricsd.ProcessesAvgReq) (*metricsd.ProcessesAvgResp, error) {
	panic("implement me")
}

func (*SQLBackend) ProcessAverage(context.Context, *metricsd.ProcessAvgReq) (*metricsd.ProcessAvgResp, error) {
	panic("implement me")
}
