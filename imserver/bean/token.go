package bean

import ()

type (
	Token struct {
		Id         string `xorm:"'t_token_id'" json:"id,omitempty"`
		Token      string `xorm:"'t_token_token'" json:"token,omitempty"`
		DeviceId   string `xorm:"'t_token_device_id'" json:"deviceId,omitempty"`
		AppId      string `xorm:"'t_token_app_id'" json:"appId,omitempty"`
		Platform   string `xorm:"'t_token_platform'" json:"platform,omitempty"`
		CreateTime string `xorm:"'t_token_create_time'" json:"createTime,omitempty"`
		UpdateTime string `xorm:"'t_token_update_time'" json:"updateTime,omitempty"`
	}
)
