package utils

import (
	"fmt"
	"regexp"
)

// RegMatch true 代表匹配
func RegMatch(pattern string, str string) (bool, error) {
	reg, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Printf("正则表达式错误: %v\n", err)
		return false, err
	}

	match := reg.FindString(str)
	return match != "", nil
}
