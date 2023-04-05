package main

import (
	"errors"
	"fmt"
)

func check(s string) (string, error) {
	if s == "" {
		err := errors.New("字符串不能为空")
		return "", err
	} else {
		return s, nil
	}
}

func main() {
	s, err := check("")
	if err != nil {
		fmt.Printf("err: %v \n", err.Error())
	} else {
		fmt.Printf("s:%v\n", s)
	}
}
