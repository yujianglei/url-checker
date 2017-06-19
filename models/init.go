package models

// import (
// 	"github.com/astaxie/beego/orm"
// 	_ "github.com/mattn/go-sqlite3"
// )

// func Init() {

// 	orm.RegisterDriver("sqlite", orm.DRSqlite)
// 	orm.RegisterDataBase("default", "sqlite3", "/data/url/url.db")
// 	orm.Debug = true
// }
import (
	"net/url"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func InitSQL() {

	host := beego.AppConfig.String("db.host")
	user := beego.AppConfig.String("db.user")
	pass := beego.AppConfig.String("db.pass")
	name := beego.AppConfig.String("db.name")
	port := beego.AppConfig.String("db.port")
	timezone := beego.AppConfig.String("db.timezone")
	if port == "" {
		port = "3306"
	}
	dsn := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + name + "?charset=utf8" + "&loc=" + url.QueryEscape(timezone)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.SetMaxIdleConns("default", 30)
	orm.SetMaxOpenConns("default", 30)
	orm.DefaultTimeLoc = time.UTC

	orm.RegisterModel(new(UrlItem))
	// orm.Debug = true
}
