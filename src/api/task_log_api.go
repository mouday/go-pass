package api

import (

	"github.com/gin-gonic/gin"
	"github.com/mouday/go-pass/src/config"
	"github.com/mouday/go-pass/src/form"
	"github.com/mouday/go-pass/src/model"
	"github.com/mouday/go-pass/src/service"
	"github.com/mouday/go-pass/src/vo"
)

func GetTaskLogList(ctx *gin.Context) {
	// taskList := service.GetTaskList()
	// database
	params := &form.PageForm{
		Page:   1,
		Size:   10,
		Status: 0,
		TaskId: "",
	}

	ctx.BindJSON(&params)

	db := config.GetDB()

	taskList := []model.TaskLogModel{}

	var count int64
	tx := db.Model(&model.TaskLogModel{})

	if params.Status != 0 {
		tx = tx.Where("status = ?", params.Status)
	}

	if params.TaskId != "" {
		tx = tx.Where("task_id = ?", params.TaskId)
	}

	tx.Count(&count)
	tx.Order("id desc").Limit(params.Size).Offset(params.PageOffset()).Find(&taskList)
	// ctx.JSON(http.StatusOK, taskList)
	vo.Success(ctx, gin.H{
		"list":  taskList,
		"total": count,
	})
}


func GetTaskLogDetail(ctx *gin.Context) {
	params := &model.TaskLogModel{}

	ctx.BindJSON(&params)

	row := model.TaskLogModel{}

	db := config.GetDB()
	db.Model(&model.TaskLogModel{}).Where("task_log_id = ?", params.TaskLogId).Find(&row)

	content := service.ReadLog(row.TaskId, params.TaskLogId)

	vo.Success(ctx, content)
}
