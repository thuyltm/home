package main

import (
	"context"
	"log"
	"testing"

	crdbpgx "github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgxv5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/cockroachdb"
)

func initTable(ctx context.Context, tx pgx.Tx) error {
	// Dropping existing table if it exists
	log.Println("Drop existing accounts table if necessary.")
	if _, err := tx.Exec(ctx, "DROP TABLE IF EXISTS accounts"); err != nil {
		return err
	}

	// Create the accounts table
	log.Println("Creating accounts table.")
	if _, err := tx.Exec(ctx,
		"CREATE TABLE accounts (id UUID PRIMARY KEY DEFAULT gen_random_uuid(), balance INT8)"); err != nil {
		return err
	}
	return nil
}

func insertRows(ctx context.Context, tx pgx.Tx, accts [4]uuid.UUID) error {
	// Insert four rows into the "accounts" table.
	log.Println("Creating new rows...")
	if _, err := tx.Exec(ctx,
		"INSERT INTO accounts (id, balance) VALUES ($1, $2), ($3, $4), ($5, $6), ($7, $8)", accts[0], 250, accts[1], 100, accts[2], 500, accts[3], 300); err != nil {
		return err
	}
	return nil
}
func TestRunCockroachContainer(t *testing.T) {
	ctx := context.Background()
	req := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Cmd: []string{
				"--insecure",
			},
		},
	}

	cockroachdbContainer, err := cockroachdb.Run(ctx, "cockroachdb/cockroach:latest",
		testcontainers.CustomizeRequest(
			req,
		),
	)

	defer func() {
		if err := testcontainers.TerminateContainer(cockroachdbContainer); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	}()
	if err != nil {
		t.Fatalf("failed to start container: %s", err)
		return
	}
	connectString, _ := cockroachdbContainer.ConnectionConfig(ctx)

	// Read in connection string
	config, err := pgx.ParseConfig(connectString.ConnString())
	if err != nil {
		log.Fatal(err)
	}
	config.RuntimeParams["application_name"] = "$ docs_simplecrud_gopgx"
	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	// Set up table
	if err := crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		return initTable(context.Background(), tx)
	}); err != nil {
		t.Fatalf("failed to create table: %v", err)
	} else {
		assert.NoError(t, err)
	}

	// Insert initial rows
	var accounts [4]uuid.UUID
	for i := 0; i < len(accounts); i++ {
		accounts[i] = uuid.New()
	}

	err = crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		return insertRows(context.Background(), tx, accounts)
	})
	if err == nil {
		assert.NoError(t, err)
	} else {
		log.Fatal("error: ", err)
	}
}
