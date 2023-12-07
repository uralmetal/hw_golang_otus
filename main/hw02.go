package main

import (
	"fmt"

	hw02 "github.com/uralmetal/hw_golang_otus/hw02_unpack_string"
)

func main() {
	var sucTests = []string{
		"a4bc2d5e", "abcd", "aaa0b",
		"", "d\n5abc", `qwe\4\5`, `qwe\45`, `qwe\\5`,
	}
	for _, test := range sucTests {
		var newString, strError = hw02.Unpack(test)
		fmt.Println(test, newString, strError)
	}
	//var failTests = []string{
	//	"3abc", "45", "aaa10b", `qw\ne`,
	//}
	//for i, test := range failTests {
	//	var newString, strError = Unpack(test)
	//	fmt.Println(i, newString, strError)
	//}
}
