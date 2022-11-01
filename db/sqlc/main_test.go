package db

import (
	"database/sql"
	"github.com/UnplugCharger/small_bank/utils"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("Unable to load config file", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Unable to connect to the database", err)
	}
	testQueries = New(testDB)

	os.Exit(m.Run())
}
