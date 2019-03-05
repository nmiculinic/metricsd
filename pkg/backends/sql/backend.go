package sql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/nmiculinic/metricsd/pkg/metricsd"
	"time"
)

var _ metricsd.MetricsServiceServer = (*Metricsd)(nil)

type Metricsd struct {
	DB *sql.DB
}

func (s *Metricsd) ReportProcessMeasurement(ctx context.Context, m *metricsd.ProcessMeasurement) (*metricsd.Empty, error) {
	if m.ProcessName == "" {
		return nil, fmt.Errorf("empty process name")
	}
	if m.Nodename == "" {
		return nil, fmt.Errorf("empty node name")
	}
	query := `
		INSERT INTO process_measurement
		  (time, nodename, process_name, timeslice, cpu, mem)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	if _, err := s.DB.ExecContext(ctx, query, time.Now(), m.Nodename, m.ProcessName, m.Timeslice, m.Cpu, m.Mem); err != nil {
		return nil, err
	}
	return &metricsd.Empty{}, nil
}

func (s *Metricsd) ReportNodeMeasurement(ctx context.Context, m *metricsd.NodeMeasurement) (*metricsd.Empty, error) {
	if m.Nodename == "" {
		return nil, fmt.Errorf("empty node name")
	}

	query := `
		INSERT INTO node_measurement
		  (time, nodename, timeslice, cpu, mem)
		VALUES ($1, $2, $3, $4, $5)
	`

	if _, err := s.DB.ExecContext(ctx, query, time.Now(), m.Nodename, m.Timeslice, m.Cpu, m.Mem); err != nil {
		return nil, err
	}
	return &metricsd.Empty{}, nil
}

func (*Metricsd) NodeAverages(context.Context, *metricsd.NodeAvgReq) (*metricsd.NodeAvgResp, error) {
	panic("implement me")
}

func (*Metricsd) ProcessesAverages(context.Context, *metricsd.ProcessesAvgReq) (*metricsd.ProcessesAvgResp, error) {
	panic("implement me")
}

func (*Metricsd) ProcessAverage(context.Context, *metricsd.ProcessAvgReq) (*metricsd.ProcessAvgResp, error) {
	panic("implement me")
}
