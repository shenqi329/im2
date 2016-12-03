package dao

import (
	"im/imserver/bean"
	"im/imserver/mysql"
	"log"
)

func InsertSession(session *bean.Session) error {

	// engine := mysql.GetXormEngine()
	// session := engine.NewSession()

	// err := session.Begin()

	// _, err = engine.Delete(bean.Token{DeviceId: token.DeviceId})
	// if err != nil {
	// 	session.Rollback()
	// 	return err
	// }

	// _, err = engine.Insert(token)

	// if err != nil {
	// 	log.Println(err.Error())
	// 	session.Rollback()
	// 	return err
	// }

	// err = session.Commit()

	// if err != nil {
	// 	log.Println(err.Error())
	// 	return err
	// }

	return nil
}

func FindSession(beans interface{}, condiBeans ...interface{}) error {

	engine := mysql.GetXormEngine()

	err := engine.Find(beans, condiBeans)

	return err
}

func GetSession(session *bean.Session) (bool, error) {
	engine := mysql.GetXormEngine()

	has, err := engine.Get(session)

	return has, err
}

func RemoveSession(session *bean.Session) (int64, error) {

	engine := mysql.GetXormEngine()

	count, err := engine.Delete(session)
	if err != nil {
		log.Println(err.Error())
	}

	return count, err
}
