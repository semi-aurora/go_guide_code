package main

import "fmt"

func main() {
	var num1 int = 99

	//结构一
	switch num1 {
	case 98, 99://条件可以多个
		fmt.Println("It's equal to 98")
	case 100:
		fmt.Println("It's equal to 100")
	default:
		fmt.Println("It's not equal to 98 or 100")
	}
	//结构二
	var num2 int = -7

	switch {
	case num2 < 0:
		fmt.Println("Number is negative")
	case num2 > 0 && num2 < 10:
		fmt.Println("Number is between 0 and 10")
	default:
		fmt.Println("Number is 10 or greater")
	}
	//结构三 包含初始化语句
	//switch a, b := x[i], y[j]; {
	//	case a < b: t = -1
	//	case a == b: t = 0
	//	case a > b: t = 1
	//}

	//例子
	var k int = 6
	switch k {
		case 4: fmt.Println("was <= 4"); fallthrough
		case 5: fmt.Println("was <= 5"); fallthrough
		case 6: fmt.Println("was <= 6"); fallthrough
		case 7: fmt.Println("was <= 7"); fallthrough
		case 8: fmt.Println("was <= 8"); fallthrough
		default: fmt.Println("default case")
	}
}