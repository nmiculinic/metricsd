package sql

import (
	"context"
	"database/sql"
	"fmt"
	svc "github.com/nmiculinic/metricsd/pkg/backends/sql"
	"github.com/nmiculinic/metricsd/pkg/metricsd"
	"github.com/stretchr/testify/suite"
	"github.com/xo/dburl"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const dbURLEnv = "TEST_DBURL"

type End2EndSuite struct {
	suite.Suite
	db  *sql.DB
	svc svc.Metricsd
}

func (s *End2EndSuite) SetupTest() {

	for _, tbl := range []string{
		"node_measurement",
		"process_measurement",
	} {
		if _, err := s.db.Exec(fmt.Sprintf("TRUNCATE TABLE %s", tbl)); err != nil {
			log.Fatalf("cannot truncate table %s %v", tbl, err)
		}
	}
}

func (s *End2EndSuite) SetupSuite() {
	url, _ := os.LookupEnv(dbURLEnv)
	u, err := dburl.Open(url)
	if err != nil {
		log.Fatalf("cannot open db, %v", err)
	}
	if err := u.Ping(); err != nil {
		log.Fatalf("cannot ping db, %v", err)
	}
	s.db = u
	s.svc.DB = u
}

func (s *End2EndSuite) TearDownSuite() {
	if err := s.db.Close(); err != nil {
		log.Fatalf("cannot close db conn %v", err)
	}
}

func (suite *End2EndSuite) TestDbPing() {
	if err := suite.db.Ping(); err != nil {
		suite.Fail("cannot ping db")
	}
}

func (suite *End2EndSuite) TestInsertNodeMeasurement() {
	if _, err := suite.svc.ReportNodeMeasurement(context.Background(), &metricsd.NodeMeasurement{
		Timeslice: 10,
		Cpu:       2,
		Mem:       2,
		Nodename:  "test",
	}); err != nil {
		suite.Failf("cannot insert", "%v", err)
	}
}

func (suite *End2EndSuite) TestInsertProcesssMeasurement() {
	if _, err := suite.svc.ReportProcessMeasurement(context.Background(), &metricsd.ProcessMeasurement{
		Timeslice:   10,
		Cpu:         2,
		Mem:         2,
		Nodename:    "test",
		ProcessName: "test",
	}); err != nil {
		suite.Failf("cannot insert", "%v", err)
	}
}

func TestIntegration(t *testing.T) {
	if _, ok := os.LookupEnv(dbURLEnv); !ok {
		t.Skipf("database url missing, env var %s", dbURLEnv)
		return
	}
	suite.Run(t, new(End2EndSuite))
}
