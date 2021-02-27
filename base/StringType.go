package base

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func StringsType() {
	// 字符串需要熟悉的标注库：strings\strconv\bytes\unicode
	var s1 string
	s1 = "lizhi"
	// strings
	bl := strings.HasPrefix(s1, "l")
	// strconv
	// 1. 整型转字符串
	num := 100
	numToString := strconv.Itoa(num)
	// 2. 字符串转整型
	str := "100"
	strToNum, _ := strconv.Atoi(str)
	// bytes
	// 偏处理流操作的类型定义、bytes是uint8的别名、rune是int32的别名、bytes和rune是两种字符类型：字节类、字符类
	// rune 可以处理一切字符
	strChinese := "我爱中国"
	// unicode
	// 判断是否是大写
	judgeString := "A b C"

	for _, word := range judgeString {
		if unicode.IsUpper(word) {
			fmt.Println(word, "is Upper")
		}
	}

	fmt.Println(s1, bl, numToString+"liziwan", strToNum*10, len(strChinese))

}
