package logic

import (
	"fmt"
	"regexp"
)

//后面
func GetKeyAndContentMapForAfter(newText string) (keyMap map[string]string, err error) {
	keyMap = make(map[string]string)
	re := regexp.MustCompile(`<\*(\d+|\*\*)\*>`)
	results := re.FindAllStringSubmatchIndex(newText, -1)
	for i := 0; i < len(results); i++ {
		var value string
		key := newText[results[i][0]:results[i][1]]
		if i < len(results)-1 {
			value = newText[results[i][1]:results[i+1][0]]
		} else {
			value = newText[results[i][1]:]
		}
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
	}
	return keyMap, nil
}
