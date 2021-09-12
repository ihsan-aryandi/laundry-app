package database

import (
	"database/sql"
	"github.com/gobuffalo/packr/v2"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
	"log"
	"os"
	"path/filepath"
	"laundry/util"
)

type database struct{}

func NewDatabase() *database {
	return &database{}
}

/*
	OpenDefaultConnection is to make a new connection to database, it will get the configuration in the .env file
*/
func (*database) OpenDefaultConnection() (*sql.DB, error) {
	driverName := os.Getenv("DB_DRIVER_NAME")
	dsn := getDefaultDSN()

	db, err := sql.Open(driverName, dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

/*
	OpenConnection is to make a new connection to database
*/
func (*database) OpenConnection(driverName, dsn string) (*sql.DB, error) {
	db, err := sql.Open(driverName, dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

/*
	CloseConnection is to close a specific database connection
*/
func (*database) CloseConnection(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

/*
	Migrate will automatically extract sql files inside /forum/migrations folder then migrate the sql queries into sql server.
	The second argument (dialect) can be empty. if it's empty, the dialect will automatically filled by driver name in .env file.
	The third argument (direction) must filled by MigrationDirection type which you can get it in github.com/rubenv/sql-migrate package.
	This function returns the numbers of applied migrations and the error if something wrong happens.
*/
func (*database) Migrate(db *sql.DB, dialect string, direction migrate.MigrationDirection) (numApplied int, err error) {
	if dialect == "" {
		dialect = os.Getenv("DB_DRIVER_NAME")
	}

	path, err := filepath.Abs("./migrations")
	if err != nil {
		return
	}

	migrations := &migrate.PackrMigrationSource{
		Box: packr.New("migrations", path),
	}
	return migrate.Exec(db, dialect, migrations, direction)
}

func getDefaultDSN() string {
	dsn := util.NewDSNCreator()
	dsn.DBName(os.Getenv("DB_NAME"))
	dsn.User(os.Getenv("DB_USER"))
	dsn.Password(os.Getenv("DB_PASSWORD"))
	dsn.Host(os.Getenv("DB_HOST"))
	dsn.SSLMode(os.Getenv("DB_SSL_MODE"))
	dsn.Schema(os.Getenv("DB_SCHEMA"))

	return dsn.Create()
}
