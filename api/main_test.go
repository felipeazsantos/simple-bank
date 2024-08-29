package api

import (
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	db "github.com/felipeazsantos/simple_bank/db/sqlc"
	"github.com/felipeazsantos/simple_bank/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

const (
	DBSource = "postgresql://root:secret@localhost:5442/simple_bank?sslmode=disable"
	DBDriver = "postgres"
)

var testQueries *db.Queries
var testDB *sql.DB

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
		DBDriver:            DBDriver,
		DBSource:            DBSource,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)
	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	var err error
	testDB, err = sql.Open(DBDriver, DBSource)
	if err != nil {
		log.Fatal("cannot connect to database", err)
	}

	testQueries = db.New(testDB)
	os.Exit(m.Run())
}
