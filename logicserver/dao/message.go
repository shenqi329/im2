package dao

import (
	"database/sql"
	"im/logicserver/bean"
	"im/logicserver/mysql"
	"log"
)

func MessageInsert(message *bean.Message) (int64, error) {

	sqlQuery := "select t_message_index from t_message where t_message_user_id=? order by t_message_index desc limit 1 for update"
	engine := mysql.GetXormEngine()
	session := engine.NewSession()
	defer session.Close()

	err := session.Begin()
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}

	var index int64 = 0
	err = session.Tx.QueryRow(sqlQuery, message.UserId).Scan(&index)
	if err != nil && err != sql.ErrNoRows {
		session.Rollback()
		log.Println(err.Error())
		return 0, err
	}
	log.Println(index)

	sqlQuery = "insert into t_message(t_message_id,t_message_session_id,t_message_user_id,t_message_type,t_message_content,t_message_index) value(?,?,?,?,?,?)"

	_, err = session.Tx.Exec(sqlQuery, message.Id, message.SessionId, message.UserId, message.Type, message.Content, index+1)

	if err != nil {
		session.Rollback()
		log.Println(err.Error())
		return 0, err
	}
	session.Commit()

	return index + 1, nil
}

func MessageMaxIndexByUserId(userId string) (int64, error) {

	engine := mysql.GetXormEngine()
	sqlQuery := "select max(t_message_index) from t_message where t_message_user_id = ?"

	var index int64
	err := engine.DB().QueryRow(sqlQuery, userId).Scan(&index)
	if err == sql.ErrNoRows || err == nil {
		log.Println(err.Error())
		return index, nil
	}
	return 0, err
}
