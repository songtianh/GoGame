package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//随机范围[0,i)

var R = rand.New(rand.NewSource(time.Now().UnixNano()))
func Random(rand int) int{
	randInt := R.Intn(rand)
	//str := fmt.Sprintf("随机最小值%d，随机最大值%d，随机值%d",0,rand-1,randInt)
	//fmt.Println(str)
	return randInt
}

func LinkStr(str1 string,str2 string) string{
	return strings.Join([]string{str1,str2},"")
}

func Trace(str string){
	fmt.Println(str)
}