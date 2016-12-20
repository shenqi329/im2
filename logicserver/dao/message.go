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

	// sql := "INSERT INTO t_message(t_message_id,t_message_index,t_message_session_id,t_message_user_id,t_message_type,t_message_content) select ?,(select max(t_message_index) from t_message where t_message_user_id = ? and t_message_session_id = ?)+1,?,?,?,?"

	// engine := mysql.GetXormEngine()

	// _, err := engine.Exec(sql, message.Id, message.UserId, message.SessionId, message.SessionId, message.UserId, message.Type, message.Content)

	// if err != nil {
	// 	log.Println(err)
	// }

	// return 1, err

	// sqlQuery := "select t_message_index from t_message where t_message_user_id=? order by t_message_index desc limit 1 for update"
	// engine := mysql.GetXormEngine()
	// session := engine.NewSession()
	// defer session.Close()

	// err := session.Begin()
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return 0, err
	// }

	// var index int64 = 0
	// err = session.Tx.QueryRow(sqlQuery, message.UserId).Scan(&index)
	// if err != nil && err != sql.ErrNoRows {
	// 	session.Rollback()
	// 	log.Println(err.Error())
	// 	return 0, err
	// }
	// log.Println(index)

	// sqlQuery = "insert into t_message(t_message_id,t_message_session_id,t_message_user_id,t_message_type,t_message_content,t_message_index) value(?,?,?,?,?,?)"

	// _, err = session.Tx.Exec(sqlQuery, message.Id, message.SessionId, message.UserId, message.Type, message.Content, index+1)

	// if err != nil {
	// 	session.Rollback()
	// 	log.Println(err.Error())
	// 	return 0, err
	// }
	// session.Commit()

	// return index + 1, nil
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
