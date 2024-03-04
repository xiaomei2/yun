package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	logic2 "wechat-ocr/api/common/logic"
)

func main() {
	text := getImagesText()
	var result logic2.PickupDemo
	result, err := logic2.SPickUpInfo(text)
	if err != nil {
		panic(err)
	}
	fmt.Printf("result:", result.Text)
	for _, info := range result.Info {
		fmt.Println(info.PickupName, ":", info.PickupCode)
	}

}
func getImagesText() string {
	// 读取本地图片文件
	fileData, err := ioutil.ReadFile("./images/6.jpg")
	if err != nil {
		panic(err)
	}
	// 将图片内容进行 Base64 编码
	imageBase64 := base64.StdEncoding.EncodeToString(fileData)

	return imageBase64
}
