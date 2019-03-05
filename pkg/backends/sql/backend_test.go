package sql

import (
	"database/sql"
	"fmt"
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
	db *sql.DB
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

func TestIntegration(t *testing.T) {
	if _, ok := os.LookupEnv(dbURLEnv); !ok {
		t.Skipf("database url missing, env var %s", dbURLEnv)
		return
	}
	suite.Run(t, new(End2EndSuite))
}
