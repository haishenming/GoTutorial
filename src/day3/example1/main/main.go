package main

import (
	"fmt"
	"strings"
)

func main() {
	var st string = "haishenming"

	fmt.Println(strings.HasPrefix(st, "hai"))   // 字符串是否已“hai”开头
	fmt.Println(strings.HasSuffix(st, "hai"))   // 字符串是否已“hai”结尾
	fmt.Println(strings.Index(st, "hai"))   	  // “hai”首次出现在字符串中的位置
	fmt.Println(strings.LastIndex(st, "h"))      // “h”最后吃现在字符串中的位置


}
