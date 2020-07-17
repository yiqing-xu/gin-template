/*
@Time : 2020/7/1 20:50
@Author : xuyiqing
@File : cms.py
*/

package handlers

import (
	"fmt"
	"gin-template/models"
	"gin-template/pkg/util"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

// 上传文件
func UploadFileHandler(ctx *gin.Context) {
	response := Response{Ctx: ctx}
	file, _ := ctx.FormFile("file")
	if len([]rune(file.Filename)) > 100 {
		response.BadRequest("文件名过长")
		return
	}
	fileModel := models.FileModel{
		Name: file.Filename,
	}
	fileDir, err := fileModel.MkMediaDir()
	if err != nil {
		panic(err)
	}
	filePath := fmt.Sprintf("%s/%s", fileDir, fileModel.Name)
	if isExist := util.FileExists(filePath); isExist {  // 判断文件名是否重复
		fileNames := strings.Split(fileModel.Name, ".")
		fileNames[0] += strconv.Itoa(int(util.GenSonyFlakeId()))
		fileName := strings.Join(fileNames, ".")
		filePath = fmt.Sprintf("%s/%s", fileDir, fileName)
	}
	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		panic(err)
	}
	fileModel.Path =  filePath
	models.DB.Create(&fileModel)
	fileModel.BuildAbsoluteUri(ctx)
	response.Response(fileModel, nil)
}

// 下载文件
func DownloadFileHandler(ctx *gin.Context) {
	fileId := ctx.Param("id")
	fileIdInt, _ := strconv.Atoi(fileId)
	var file models.FileModel
	models.DB.First(&file, fileIdInt)
	ctx.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", file.Name))
	ctx.Writer.Header().Set("Content-Type", "application/octet-stream")
	ctx.File(file.Path)
}
