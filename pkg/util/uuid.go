package util

import (
	"github.com/sony/sonyflake"
)

var flake = sonyflake.NewSonyflake(sonyflake.Settings{})

func GenSonyFlakeId() uint64  {
	uuid, err := flake.NextID()
	if err != nil {
		panic(err)
	}
	return uuid
}
