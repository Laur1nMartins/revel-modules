package gorp

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-gorp/gorp"

	// mysql package.
	_ "github.com/jinzhu/gorm/dialects/mysql"

	// postgres package.
	_ "github.com/jinzhu/gorm/dialects/postgres"

	// mysql package.
	"github.com/Laur1nMartins/revel"
	"github.com/Laur1nMartins/revel/logger"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	// The database map to use to populate data.
	Db           = &DbGorp{}
	moduleLogger logger.MultiLogger
)

func init() {
	revel.RegisterModuleInit(func(module *revel.Module) {
		moduleLogger = module.Log
		moduleLogger.Debug("Assigned Logger")
	})
}

func (dbGorp *DbGorp) InitDb(open bool) (err error) {
	dbInfo := dbGorp.Info

	switch dbInfo.DbDriver {
	default:
		dbGorp.SqlStatementBuilder = sq.StatementBuilder.PlaceholderFormat(sq.Question)
		dbInfo.Dialect = gorp.SqliteDialect{}
		if len(dbInfo.DbConnection) == 0 {
			dbInfo.DbConnection = fmt.Sprintf(dbInfo.DbHost)
		}
	case "ql":
		fallthrough
	case "ql-mem":
		fallthrough
	case "postgres":
		dbGorp.SqlStatementBuilder = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
		dbInfo.Dialect = gorp.PostgresDialect{}
		if len(dbInfo.DbConnection) == 0 {
			dbInfo.DbConnection = fmt.Sprintf("host=%s port=8500 user=%s dbname=%s sslmode=disable password=%s", dbInfo.DbHost, dbInfo.DbUser, dbInfo.DbName, dbInfo.DbPassword)
		}
	case "mysql":
		dbGorp.SqlStatementBuilder = sq.StatementBuilder.PlaceholderFormat(sq.Question)
		dbInfo.Dialect = gorp.MySQLDialect{}
		if len(dbInfo.DbConnection) == 0 {
			dbInfo.DbConnection = fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local", dbInfo.DbUser, dbInfo.DbPassword, dbInfo.DbHost, dbInfo.DbName)
		}
	}

	if open {
		err = dbGorp.OpenDb()
	}
	return
}

// Initialize the database from revel.Config.
func InitDb(dbGorp *DbGorp) error {
	params := DbInfo{}
	params.DbDriver = revel.Config.StringDefault("db.driver", "sqlite3")
	params.DbHost = revel.Config.StringDefault("db.host", "localhost")
	if params.DbDriver == "sqlite3" && params.DbHost == "localhost" {
		params.DbHost = "/tmp/app.db"
	}
	params.DbUser = revel.Config.StringDefault("db.user", "default")
	params.DbPassword = revel.Config.StringDefault("db.password", "")
	params.DbName = revel.Config.StringDefault("db.name", "default")
	params.DbConnection = revel.Config.StringDefault("db.connection", "")
	params.DbSchema = revel.Config.StringDefault("db.schema", "")
	dbGorp.Info = &params

	return dbGorp.InitDb(true)
}
