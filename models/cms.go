/*
@Time : 2020/7/15 21:41
@Author : xuyiqing
@File : cms.py
*/

package models

import (
	"fmt"
	"gin-template/conf"
	"gin-template/pkg/util"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"time"
)

type FileModel struct {
	BaseModel
	Name string `gorm:"comment:'文件名';column:name;not null;" json:"name"`
	Path string `gorm:"comment:'路径';column:path;not null;" json:"path"`
}

func (file *FileModel) TableName() string {
	return "cms_files"
}

// 获取年月日文件夹
func (file *FileModel) DatePath() string {
	now := time.Now()
	year := now.Year()
	month := now.Month().String()
	day := now.Day()
	return fmt.Sprintf("%s/%s/%s/%s",
		conf.ProjectCfg.MediaFilePath,
		strconv.Itoa(year),
		month,
		strconv.Itoa(day))
}

// 创建 年/月/日 文件夹
func (file *FileModel) MkMediaDir() (string, error) {
	dir := conf.ProjectCfg.MediaFilePath + time.Now().Format("2006/01/02")
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return dir, err
	} else {
		return dir, nil
	}
}

// 构建全路径url
func (file *FileModel) BuildAbsoluteUri(ctx *gin.Context) {
	file.Path = util.BuildAbsoluteUri(ctx, file.Path)
}
