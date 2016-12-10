package bean

import (
	"time"
)

type (
	Message struct {
		Id         int64      `xorm:"'t_message_id'" json:"id,omitempty"`
		SessionId  int64      `xorm:"'t_message_session_id'" json:"sessionId,omitempty"`
		Type       int        `xorm:"'t_message_type'" json:"type,omitempty"`
		Content    string     `xorm:"'t_message_content'" json:"content,omitempty"`
		Index      int64      `xorm:"'t_message_index'" json:"index,omitempty"`
		CreateTime *time.Time `xorm:"'t_message_create_time'" json:"createTime,omitempty"`
	}
)

func (t Message) TableName() string {
	return "t_message"
}
