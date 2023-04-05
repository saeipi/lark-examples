package main

import (
	"fmt"
	"lark/pkg/common/xjwt"
)

func main() {
	token, err := xjwt.CreateToken(1, 1, true, 3600*12)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(token.Token)
}
