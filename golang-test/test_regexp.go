package main

import (
	"fmt"
	"regexp"
)

func main() {
	tailNumber := "21300001234"

	pattern := fmt.Sprintf(`^*(%s)$`, tailNumber)

	regex := regexp.MustCompile(pattern)

	phoneNumbers := []string{
		"11300001234",
	}

	for _, phoneNumber := range phoneNumbers {
		if regex.MatchString("") {
			fmt.Println(phoneNumber, "匹配成功")
		} else {
			fmt.Println(phoneNumber, "匹配失败")
		}
	}
}
