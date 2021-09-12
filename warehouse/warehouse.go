package warehouse

import (
	"database/sql"
	"github.com/bandungrhapsody/rhaplogger"
)

/*
	Connections
*/
type Connections struct {
	DB *sql.DB
}

var Conn = &Connections{}

/*
	Application Logger
*/
var Log *rhaplogger.RhapLogger
