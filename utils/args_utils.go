package utils

import (
	"os"
)

func FindArg(key string) bool{
	for _, arg := range os.Args {
		if arg==key{
			return true
		}
	}
	return false
}
