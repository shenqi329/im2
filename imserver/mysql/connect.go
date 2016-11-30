package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"runtime"
)

var engine *xorm.Engine

func GetXormEngine() *xorm.Engine {
	if engine == nil {
		if runtime.GOOS == "windows" {
			eng, err := xorm.NewEngine("mysql", "user_connect:user_connect@tcp(localhost:3306)/db_im?charset=utf8")
			if err != nil {
				log.Println(err.Error())
				return nil
			}
			engine = eng
			engine.ShowSQL(true)
		} else {
			eng, err := xorm.NewEngine("mysql", "user_connect:user_connect@tcp(172.17.0.2:3306)/db_im?charset=utf8")
			if err != nil {
				log.Println(err.Error())
				return nil
			}
			engine = eng
			engine.ShowSQL(true)
		}

		//engine.ShowExecTime(true)
	}
	return engine
}
