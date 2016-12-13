package dao

import (
	"database/sql"
	"im/logicserver/bean"
	"im/logicserver/mysql"
)

// DELIMITER //
// create procedure t_message_get_increment_index(in session_id bigint(20), in type int(4) ,in context varchar(20000),in small_index bigint, out index_out bigint)
// begin
// declare oldindex bigint;
// start transaction;
// select max(t_message_index) into oldindex from t_message where t_message_session_id=session_id for update;
// if oldindex is NULL then
// insert into t_message(t_message_session_id,t_message_type,t_message_content,t_message_index) value(session_id, type,context,small_index);
// set index_out=small_index;
// else
// insert into t_message(t_message_session_id,t_message_type,t_message_content,t_message_index) value(session_id, type,context,oldindex+1);
// set index_out=oldindex+1;
// end if;
// commit;
// end;
// //
// DELIMITER ;

func MessageInsert(message *bean.Message) (int64, error) {

	engine := mysql.GetXormEngine()

	_, err := engine.Exec("call t_message_insert(?,?,?,?,?,@p_out);", message.Id, message.SessionId, message.UserId, message.Type, message.Content)

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
