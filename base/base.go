package base

import "fmt"

var (
	a int
	b int
)

const LEVEL_MAX  = 120

func Yusuan() {
	fmt.Println("基本运算")

	a := 98
	b := 89

	fmt.Println(fmt.Sprintf("%d+%d=%d",a,b,a+b))

	fmt.Println(fmt.Sprintf("%d-%d=%d",a,b,a-b))

	fmt.Println(fmt.Sprintf("%d*%d=%d",a,b,a*b))

	fmt.Println(fmt.Sprintf("%d/%d=%d",a,b,a/b))

	fmt.Println(add(a,b))
	_,_,a = add(a,b) //只获取函数返回值的最后一个
	fmt.Println(a)

	fmt.Printf("最大等级为%d\n",LEVEL_MAX)
}

func add(a int,b int) (int,int,int){
	fmt.Println(a + b)
	return a,b,a+b
}
func HelloWorld(){
	fmt.Println("你好,世界")
}
