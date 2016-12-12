package dao

import (
	"database/sql"
	//"fmt"
	"im/logicserver/bean"
	"im/logicserver/mysql"
)

//t_message

// defer session.resetStatement()
// 	if session.IsAutoClose {
// 		defer session.Close()
// 	}

// 	var sqlStr string
// 	var args []interface{}
// 	if session.Statement.RawSQL == "" {
// 		sqlStr, args = session.Statement.genCountSQL(bean)
// 	} else {
// 		sqlStr = session.Statement.RawSQL
// 		args = session.Statement.RawParams
// 	}

// 	session.queryPreprocess(&sqlStr, args...)

// 	var err error
// 	var total int64
// 	if session.IsAutoCommit {
// 		err = session.DB().QueryRow(sqlStr, args...).Scan(&total)
// 	} else {
// 		err = session.Tx.QueryRow(sqlStr, args...).Scan(&total)
// 	}

// 	if err == sql.ErrNoRows || err == nil {
// 		return total, nil
// 	}

// 	return 0, err

func MessageInsert(message *bean.Message) (int64, error) {

	engine := mysql.GetXormEngine()

	_, err := engine.Exec("call t_message_get_increment_index(?,?,?,?,@p_out);", message.SessionId, message.Type, message.Content, 1)

	if err != nil {
		return 0, err
	}
	return 1, nil
}

func MessageMaxIndex(sessionId int64) (int64, error) {

	engine := mysql.GetXormEngine()
	sqlQuery := "select count(t_message_index) from t_message where t_message_session_id = ?"

	var index int64
	err := engine.DB().QueryRow(sqlQuery, sessionId).Scan(&index)
	if err == sql.ErrNoRows || err == nil {
		return index, nil
	}

	return 0, err
}
