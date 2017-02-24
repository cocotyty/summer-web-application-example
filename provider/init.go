package provider

import (
	_ "github.com/go-sql-driver/mysql"
	providers "github.com/cocotyty/summer-providers"
	"github.com/cocotyty/summer"
)

const (
	DBMain = "DB:Main"
)

func init() {
	summer.Add(DBMain, &providers.SqlXDB{}, "databases.main")
}
