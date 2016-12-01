package bean

import (
	"time"
)

type (
	Token struct {
		Id         int64      `xorm:"'t_token_id'" json:"id,omitempty"`
		UserId     string     `xorm:"'t_token_user_id'" json:"userId,omitempty"`
		DeviceId   string     `xorm:"'t_token_device_id'" json:"deviceId,omitempty"`
		AppId      string     `xorm:"'t_token_app_id'" json:"appId,omitempty"`
		Platform   string     `xorm:"'t_token_platform'" json:"platform,omitempty"`
		CreateTime *time.Time `xorm:"'t_token_create_time'" json:"createTime,omitempty"`
		UpdateTime *time.Time `xorm:"'t_token_update_time'" json:"updateTime,omitempty"`
	}
)

func (u Token) TableName() string {
	return "t_token"
}
