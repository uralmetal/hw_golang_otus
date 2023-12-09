package test

import (
	"fmt"

	hw02 "github.com/uralmetal/hw_golang_otus/hw02_unpack_string"
)

func main() {
	fmt.Println("---Test valid string---")
	var sucTests = []string{
		"a4bc2d5e", "abcd", "aaa0b",
		"", "d\n5abc", `qwe\4\5`, `qwe\45`, `qwe\\5`,
		"Hello, 世4界",
	}
	for _, test := range sucTests {
		var newString, strError = hw02.Unpack(test)
		fmt.Println(test, newString, strError)
	}
	fmt.Println("---Test invalid string---")
	var failTests = []string{
		"3abc", "45", "aaa10b", `qw\ne`, `test\`,
	}
	for _, test := range failTests {
		var newString, strError = hw02.Unpack(test)
		fmt.Println(test, newString, strError)
	}
}
