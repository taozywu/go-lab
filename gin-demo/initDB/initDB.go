package initDB

// db配置
import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "zerotest:#*$#@($&32432432!@@tcp(192.168.1.235:3306)/livechat")
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	Db.SetMaxOpenConns(10) // 最大打开数
	Db.SetMaxIdleConns(10) // 最大闲置数
}
