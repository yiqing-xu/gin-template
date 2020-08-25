/*
@Time : 2020/7/15 23:27
@Author : xuyiqing
@File : util.go
*/

package util

import (
	"bytes"
	"encoding/json"
	"os"
)


// 判断所给路径文件/文件夹是否存在
func FileOrDirExists(path string) bool {
	_, err := os.Stat(path)    //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 解决json字符串整型精度缺失
func PrecisionLost(data interface{}) (map[string]interface{}, error) {
	bdata, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	var val map[string]interface{}
	d := json.NewDecoder(bytes.NewBuffer(bdata))
	d.UseNumber()
	err = d.Decode(&val)
	if err != nil {
		return nil, err
	}
	return val, nil
}
