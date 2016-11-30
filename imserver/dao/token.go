package dao

import (
	"im/imserver/bean"
	"log"
	"sso/mysql"
)

func InsertToken(token *bean.Token) (int64, error) {

	engine := mysql.GetXormEngine()

	count, err := engine.Insert(token)

	return count, err
}

func GetToken(token *bean.Token) (bool, error) {
	engine := mysql.GetXormEngine()

	has, err := engine.Get(token)

	return has, err
}

func RemoveToken(token *bean.Token) (int64, error) {

	engine := mysql.GetXormEngine()

	count, err := engine.Delete(token)
	if err != nil {
		log.Println(err.Error())
	}

	return count, err
}
