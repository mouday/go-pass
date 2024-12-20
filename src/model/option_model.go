package model

type OptionModel struct {
	Id         uint   `json:"id"`
	QuestionId uint   `json:"questionId"`
	Label      string `json:"label"`
	Status     bool   `json:"status"`
}

// 自定义表名
func (OptionModel) TableName() string {
	return "tb_option"
}
