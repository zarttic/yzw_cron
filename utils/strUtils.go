package utils

import "time"

func TimeStr(formatter string) string {
	return time.Now().Format(formatter)
}
