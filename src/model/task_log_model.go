package model

import "github.com/mouday/go-pass/src/utils"

type TaskLogModel struct {
	Id        uint   `json:"-"`
	TaskLogId string `gorm:"index" json:"taskLogId"`
	TaskId    string `json:"taskId"`
	Title     string `json:"title"`
	Url       string `json:"url"`
	// RunnerId   string          `json:"runnerId"`
	// TaskName   string          `json:"taskName"`
	Status int `json:"status"`
	// Result     string          `json:"result"`
	EndTime    utils.LocalTime `gorm:"type:datetime" json:"endTime"`
	CreateTime utils.LocalTime `gorm:"type:datetime;autoCreateTime" json:"createTime"`
	UpdateTime utils.LocalTime `gorm:"type:datetime;autoUpdateTime" json:"updateTime"`
}

// 自定义表名
func (TaskLogModel) TableName() string {
	return "tb_log_task"
}
