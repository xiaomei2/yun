package logic

import (
	"fmt"
	"regexp"
)

func GetKeyInformation(text string) (info string, err error) {
	pattern := `\w+-[\w-]+`
	re, err := regexp.Compile(pattern)
	if err != nil {
		return "", fmt.Errorf("正则表达式编译错误: %s", err.Error())
	}
	match := re.FindString(text)
	if match == "" {
		return "", nil
	}
	return match, nil
}
