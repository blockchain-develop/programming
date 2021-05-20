package problem8

import (
	"fmt"
	"github.com/programming/utils"
	"testing"
)

func TestAddTwoNumbers1(t *testing.T) {
	num1 := utils.Slice2List([]int{2,4,3})
	num2 := utils.Slice2List([]int{5,6,4})
	{
		result := addTwoNumbers(num1, num2)
		fmt.Printf("result: %v\n", utils.PrintList(result))
	}
}

func TestAddTwoNumbers2(t *testing.T) {
	num1 := utils.Slice2List([]int{0})
	num2 := utils.Slice2List([]int{0})
	{
		result := addTwoNumbers(num1, num2)
		fmt.Printf("result: %v\n", utils.PrintList(result))
	}
}

func TestAddTwoNumbers3(t *testing.T) {
	num1 := utils.Slice2List([]int{9,9,9,9,9,9,9})
	num2 := utils.Slice2List([]int{9,9,9,9})
	{
		result := addTwoNumbers(num1, num2)
		fmt.Printf("result: %v\n", utils.PrintList(result))
	}
}


