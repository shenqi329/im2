package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"runtime"
)

//// ?nul)5E75rdu

var engine *xorm.Engine

func GetXormEngine() *xorm.Engine {
	if engine == nil {
		log.Println("engine")
		if runtime.GOOS == "windows" {
			eng, err := xorm.NewEngine("mysql", "im_connect:im_connect@tcp(localhost:3306)/db_im?charset=utf8")
			if err != nil {
				log.Println(err.Error())
				return nil
			}
			engine = eng
			//engine.ShowSQL(true)
		} else {
			eng, err := xorm.NewEngine("mysql", "im_connect:im_connect@tcp(localhost:3306)/db_im?charset=utf8")
			if err != nil {
				log.Println(err.Error())
				return nil
			}
			engine = eng
			//engine.ShowSQL(true)
		}
		engine.SetMaxIdleConns(200)
		engine.SetMaxOpenConns(1000)
		//engine.ShowExecTime(true)
	}
	return engine
}