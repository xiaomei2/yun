package logic

import (
	"fmt"
	"regexp"
)

//前面
func GetKeyAndContentMapForFront(newText string) (keyMap map[string]string, err error) {
	keyMap = make(map[string]string)
	re := regexp.MustCompile(`<\*(\d+|\*\*)\*>`)
	results := re.FindAllStringSubmatchIndex(newText, -1)
	startIndex := 0
	for i := 0; i < len(results); i++ {
		//关键字
		key := newText[results[i][0]:results[i][1]]
		//关键字及其前面文本
		value := newText[startIndex:results[i][0]]
		//
		result, err := GetKeyInformation(value)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if result == "" {
			keyMap[key] = "未找到匹配项"
		} else {
			keyMap[key] = result
		}
		if i < len(results)-1 {
			startIndex = results[i][1]
		}
	}
	return keyMap, nil
}
