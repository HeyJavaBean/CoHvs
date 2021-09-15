package utils

import (
	"CoHvs/constant"
)

func GetMapPort() int{
	return constant.ServerPortBase
}

func GetPlayerPort(n int) int{
	return constant.ServerPortBase+n+1
}
