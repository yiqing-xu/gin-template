package util

import (
	"fmt"
	"time"
)
var t = time.Unix(1594909810, 0)

var flake = NewSonyflake(Settings{

})

func GenSonyFlakeId() uint64 {
	uuid, err := flake.NextID()
	if err != nil {
		fmt.Println(err)
	}
	return uuid
}
