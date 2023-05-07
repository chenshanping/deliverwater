package model

type UserInfo struct {
	Id       int32  `json:"id,omitempty"`       // ID
	Email    string `json:"email,omitempty"`    // 邮箱
	Password string `json:"password,omitempty"` // 密码
	Username string `json:"username,omitempty"` // 姓名
}

func (UserInfo) TableName() string {
	return "user"
}
