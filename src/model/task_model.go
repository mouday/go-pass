package model

import "github.com/mouday/go-pass/src/utils"

type TaskModel struct {
	Id     uint   `json:"-"`
	TaskId string `gorm:"index" json:"taskId"`
	Title  string `json:"title"`
	// RunnerId   string          `json:"runnerId"`
	// TaskName   string          `json:"taskName"`
	// Method     string          `json:"cron"`
	Cron       string          `json:"cron"`
	Url        string          `json:"url"`
	Status     bool            `json:"status"`
	CreateTime utils.LocalTime `gorm:"type:datetime;autoCreateTime" json:"createTime"`
	UpdateTime utils.LocalTime `gorm:"type:datetime;autoUpdateTime" json:"updateTime"`
}

// 自定义表名
func (TaskModel) TableName() string {
	return "tb_task"
}
