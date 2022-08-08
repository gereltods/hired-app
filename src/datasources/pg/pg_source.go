package pg

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	pgsqlUsername = "pg_username"
	pgsqlPassword = "pg_password"
	pgsqlHost     = "pg_host"
	pgsqlSchema   = "pg_schema"
)

var (
	PgPool   *pgxpool.Pool
	username = os.Getenv(pgsqlUsername)
	password = os.Getenv(pgsqlPassword)
	host     = os.Getenv(pgsqlHost)
	schema   = os.Getenv(pgsqlSchema)
)

func init() {

	//User ID=tmp;Password=Tmp@123;Host=103.50.205.205;Port=5432;Database=oky;
	databaseUrl := "postgres://tmp:Tmp@123@103.50.205.205:5432/oky"

	var err error
	// this returns connection pool
	PgPool, err = pgxpool.Connect(context.Background(), databaseUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprint(os.Stderr, "Connect to database: \n")
	// to close DB pool
	//defer PgPool.Close()
}
