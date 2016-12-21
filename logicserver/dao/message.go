package dao

import (
	//"database/sql"
	"im/logicserver/bean"
	"im/logicserver/mysql"
	"log"
	"time"
)

func MessageInsert(message *bean.Message) (int64, error) {

	var err error
	for i := 0; i < 1000; i++ {
		index, err := MessageMaxIndexByUserId(message.UserId)
		if err != nil {
			return 0, err
		}

		message.Index = index + 1
		_, err = NewDao().Insert(message)
		if err == nil {
			break
		}
		time.Sleep((time.Duration)(i) * time.Millisecond)
	}
	return 0, err
}

func MessageMaxIndexByUserId(userId string) (uint64, error) {

	engine := mysql.GetXormEngine()
	sqlQuery := "select max(t_message_index) from t_message where t_message_user_id = ?"

	var index interface{}
	err := engine.DB().QueryRow(sqlQuery, userId).Scan(&index)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	log.Println(index)
	ret, ok := index.(int64)
	if ok {
		log.Println(ret)
		return (uint64)(ret), nil
	}
	return 0, err
}
