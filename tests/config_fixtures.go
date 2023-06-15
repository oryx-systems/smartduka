package tests

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"text/template"
	"time"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/oryx-systems/smartduka/pkg/smartduka/application/common/testutils"
	"github.com/oryx-systems/smartduka/pkg/smartduka/application/utils"
	"github.com/oryx-systems/smartduka/pkg/smartduka/infrastructure/datastore/db/gorm"
)

var (
	fixtures  *testfixtures.Loader
	testingDB *gorm.PGInstance
	db        *sql.DB

	// Pin variables
	salt, encryptedPin string
	userID             = "6ecbbc80-24c8-421a-9f1a-e14e12678ee0"
	testPhone          = "+254722000000"
)

func setupFixtures() {
	isLocalDB := testutils.CheckIfCurrentDBIsLocal()
	if !isLocalDB {
		fmt.Println("Cannot run tests. The current database is not a local database.")
		os.Exit(1)
	}

	log.Println("setting up test database")
	var err error

	testingDB, err = gorm.NewPGInstance()
	if err != nil {
		fmt.Println("failed to initialize db:", err)
		os.Exit(1)
	}
	db, err = testingDB.DB.DB()
	if err != nil {
		fmt.Println("failed to initialize db:", err)
		os.Exit(1)
	}

	// setup test variables
	salt, encryptedPin = utils.EncryptPIN("0000", nil)

	fixtures, err = testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("postgres"),
		testfixtures.Template(),
		testfixtures.TemplateData(template.FuncMap{
			"salt":         salt,
			"hash":         encryptedPin,
			"valid_to":     time.Now().Add(500).String(),
			"test_user_id": userID,
			"test_phone":   "\"" + testPhone + "\"",
		}),
		// this is the directory containing the YAML files.
		// The file name should be the same as the table name
		// order of inserting values matter to avoid foreign key constraint errors
		testfixtures.Paths(
			"../fixtures/smartduka_user.yml",
			"../fixtures/smartduka_contact.yml",
		),
		// uncomment when running tests locally, if your db is not a test db
		// Ensure the testing db in the ci is named `test`
		// !!Warning!!: this can corrupt data, do not turn on or run tests while in non-test db
		testfixtures.DangerousSkipTestDatabaseCheck(),
	)
	if err != nil {
		fmt.Println("failed to create fixtures:", err)
		os.Exit(1)

	}

	err = prepareTestDatabase()
	if err != nil {
		fmt.Println("failed to prepare test database:", err)
		os.Exit(1)
	}

}

func prepareTestDatabase() error {
	if err := fixtures.Load(); err != nil {
		return err
	}
	return nil
}
