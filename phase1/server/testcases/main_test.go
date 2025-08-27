package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	. "home/phase1/server/handler"
	. "home/phase1/server/service"

	"github.com/cockroachdb/cockroach-go/v2/testserver"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type testContext struct {
	t  *testing.T
	ts testserver.TestServer
	db *sql.DB
}

func (c *testContext) beforeEach() {
	log.Println("Setupping database for testcases goes here.")
	var err error
	//1. Initialize the TestServer
	c.ts, err = testserver.NewTestServer()
	if err != nil {
		c.t.Fatalf("failed to create test server: %v", err)
	}
	// 2. Open a SQL connection
	c.db, err = sql.Open("postgres", c.ts.PGURL().String())
	if err != nil {
		c.t.Fatalf("failed to open database connection: %v", err)
	}
	//Ping the database to ensure connection is established
	if err := c.db.Ping(); err != nil {
		c.t.Fatalf("failed to ping databse: %v", err)
	}
	createMessageTable(c.t, c.db)
}

func (c *testContext) afterEach() {
	log.Println("Tearing down database for testcases goes here.")
	c.ts.Stop()  // Ensure the server is stopped after the test
	c.db.Close() //Ensure the database connection is closed
}

func testCase(test func(t *testing.T, c *testContext)) func(*testing.T) {
	return func(t *testing.T) {
		ctx := &testContext{t: t}
		ctx.beforeEach()
		defer ctx.afterEach()
		test(t, ctx)
	}
}

func TestMainLifecycle(t *testing.T) {
	t.Run("TestCheckCockroachdb", testCase(func(t *testing.T, c *testContext) {
		log.Println("TestCheckCockroachdb")
		checkCockroachdbServer(t, c)
	}))
	t.Run("testCreateAndRecheckMessageInsertedToDB", testCase(func(t *testing.T, c *testContext) {
		log.Println("testCreateAndRecheckMessageInsertedToDB")
		testCreateMessageToDB(t, c)
		testNewMessageInsertedToDB(t, c, 1)
	}))
	t.Run("testDefaultRequest", testCase(func(t *testing.T, c *testContext) {
		log.Println("testDefaultRequest")
		testNewMessageInsertedToDB(t, c, 0)
	}))
}

func testCreateMessageToDB(t *testing.T, testContext *testContext) {
	// Create a new Echo instance
	e := echo.New()
	// Create a new http.Request and http.ResponseRecorder
	req := httptest.NewRequest(http.MethodPost, "/send", strings.NewReader(`{"value":"Hello, Docker!"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	//Create an echo context from the request and recorder
	c := e.NewContext(req, rec)
	// Call the handler function directly
	messageService := NewCockroachDBMessageService(testContext.db)
	messageHandler := NewMessageHandler(messageService)
	err := messageHandler.CreateMessage(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func testNewMessageInsertedToDB(t *testing.T, testContext *testContext, expectedRecords int) {
	// Create a new Echo instance
	e := echo.New()
	// Create a new http.Request and http.ResponseRecorder
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	//Create an echo context from the request and recorder
	c := e.NewContext(req, rec)
	// Call the handler function directly
	messageService := NewCockroachDBMessageService(testContext.db)
	messageHandler := NewMessageHandler(messageService)
	err := messageHandler.CountMessages(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, fmt.Sprintf("Hello, Docker! (%d)", expectedRecords), strings.TrimSpace(rec.Body.String()))
}

func createMessageTable(t *testing.T, db *sql.DB) {
	if _, err := db.Exec(
		"create table if not exists message (value text primary key)"); err != nil {
		t.Fatalf("failed to create table: %v", err)
	}
}

func checkCockroachdbServer(t *testing.T, c *testContext) {
	//3. Create a table
	createTableSQL := `
		create table if not exists items (
			id int primary key,
			name string
		);
	`
	_, err := c.db.Exec(createTableSQL)
	if err != nil {
		t.Fatalf("failed to create table: %v", err)
	}
	// 4. Insert data
	insertItemSQL := `insert into items (id, name) values ($1, $2);`
	_, err = c.db.Exec(insertItemSQL, 1, "Test Item")
	if err != nil {
		t.Fatalf("failed to insert item: %v", err)
	}
	log.Println("Item successfully added to the test server database.")
	//Optional: Verify the item was added
	var itemName string
	err = c.db.QueryRow("select name from items where id = $1", 1).Scan(&itemName)
	if err != nil {
		t.Fatalf("failed to query item: %v", err)
	}
	if itemName != "Test Item" {
		t.Errorf("expected item name 'Test Item', got '%s'", itemName)
	}
}
