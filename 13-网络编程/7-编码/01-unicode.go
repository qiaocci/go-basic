package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "a"
	fmt.Println([]byte(str)) // 97 unicode编码 https://unicode-table.com/cn/0061/

	str2 := "乔"
	fmt.Println([]byte(str2)) // [228 185 148] https://unicode-table.com/cn/4E54/

	// string to binary 方法一
	_ = fmt.Sprintf("%08b", []byte("A")) // 1000001

	// string to binary 方法一
	str3 := "A"
	fmt.Println([]byte(str3)) // [65]
	var v int64 = 65
	s2 := strconv.FormatInt(v, 2)
	fmt.Println(s2) // 1000001

	// 十六进制转二进制 方法一
	// bitSize
	// bitSize参数表示转换为什么位的int/uint，有效值为0、8、16、32、64。
	// 当bitSize=0的时候，表示转换为int或uint类型。例如bitSize=8表示转换后的值的类型为int8或uint8
	res, _ := strconv.ParseUint("6B66", 16, 64)
	fmt.Println(res)                 // 27494
	fmt.Println(convertToBin(27494)) // 110101101100110

	fmt.Println("====")
	// 十六进制转二进制 方法二
	year := 0x5
	fmt.Printf("十六进制0x7e4对应的二进制表示为:[%b]\n", year) // 110101101100110
}
func convertToBin(num int) string {
	s := ""

	if num == 0 {
		return "0"
	}

	// num /= 2 每次循环的时候 都将num除以2  再把结果赋值给 num
	for ; num > 0; num /= 2 {
		lsb := num % 2
		// strconv.Itoa() 将数字强制性转化为字符串
		s = strconv.Itoa(lsb) + s
	}
	return s
}
