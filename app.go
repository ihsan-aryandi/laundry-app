package main

import (
	"fmt"
	"github.com/bandungrhapsody/rhaplogger"
	migrate "github.com/rubenv/sql-migrate"
	"laundry/core/database"
	"laundry/core/env"
	"laundry/core/router"
	"laundry/util"
	"laundry/warehouse"
	"log"
	"os"
)

var (
	db        = database.NewDatabase()
	envReader = env.NewEnv()
	rhRouter  = router.NewRhRouter()
	argsUtil  = util.NewArgsFinder()
)

/*
	StartApplication is an application starter function
*/
func StartApplication() {
	var err error

	/*
		Read env
	*/
	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(mydir)
	err = envReader.ReadEnv()
	if err != nil {
		log.Fatal(err)
	}

	/*
		Logger
	*/
	config := &rhaplogger.LogConfig{
		Filename: os.Getenv("LOG_FILE"),
		Stdout:   true,
	}
	logger, err := rhaplogger.NewRhapLogger(config)
	if err != nil {
		log.Fatal(err)
	}
	warehouse.Log = logger

	/*
		Connect to main DB
	*/
	warehouse.Conn.DB, err = db.OpenDefaultConnection()
	if err != nil {
		errLog := warehouse.Log.NewLogError()
		errLog.Message = err.Error()
		errLog.StatusCode = 500
		errLog.Print()

		log.Fatal(err)
	}
	info := warehouse.Log.NewLogInfo()
	info.Message = "Connected to database"
	info.StatusCode = 200
	info.Print()

	/*
		DB migrations
	*/
	numApplied, err := migrateDb()
	if err != nil {
		errLog := warehouse.Log.NewLogError()
		errLog.Message = err.Error()
		errLog.StatusCode = 500
		errLog.Print()

		log.Fatal(err)
	}
	info = warehouse.Log.NewLogInfo()
	info.Message = fmt.Sprintf("Applied %d migrations", numApplied)
	info.StatusCode = 200
	info.Print()

	/*
		Setup & start Router
	*/
	rhRouter.SetupRoutes()
	log.Fatal(rhRouter.Listen(":" + os.Getenv("PORT")))
}

func migrateDb() (int, error) {
	migrateDirection := migrate.Up
	if argsUtil.Has("migrate-down") {
		migrateDirection = migrate.Down
	}

	return db.Migrate(warehouse.Conn.DB, "", migrateDirection)
}
