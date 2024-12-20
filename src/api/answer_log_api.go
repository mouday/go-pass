package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mouday/go-pass/src/config"
	"github.com/mouday/go-pass/src/form"
	"github.com/mouday/go-pass/src/model"

	"github.com/mouday/go-pass/src/vo"
)

func GetAnswerLogList(ctx *gin.Context) {
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

	taskList := []model.AnswerLogModel{}

	var count int64
	tx := db.Model(&model.AnswerLogModel{})

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

func GetAnswerLogDetail(ctx *gin.Context) {
	params := &model.AnswerLogModel{}

	ctx.BindJSON(&params)

	row := model.AnswerLogModel{}

	db := config.GetDB()
	db.Model(&model.AnswerLogModel{}).Where("Id = ?", params.Id).Find(&row)

	// content := service.ReadLog(row.TaskId, params.Id)

	vo.Success(ctx, nil)
}
