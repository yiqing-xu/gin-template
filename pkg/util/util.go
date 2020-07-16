/*
@Time : 2020/7/15 23:27
@Author : xuyiqing
@File : util.py
*/

package util

import (
	"bytes"
	"encoding/json"
	"math"
	"os"
)


// 判断所给路径文件/文件夹是否存在
func FileExists(path string) bool {
	_, err := os.Stat(path)    //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

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

//将float64转成精确的int64
func Wrap(num float64, retain int) int64 {
	return int64(num * math.Pow10(retain))
}

//将int64恢复成正常的float64
func Unwrap(num int64, retain int) float64 {
	return float64(num) / math.Pow10(retain)
}

//精准float64
func WrapToFloat64(num float64, retain int) float64 {
	return num * math.Pow10(retain)
}

//精准int64
func UnwrapToInt64(num int64, retain int) int64 {
	return int64(Unwrap(num, retain))
}
