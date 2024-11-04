package main

import (
	"fmt"
	"task-basic/utils"
)

func main() {
	n := 3
	fmt.Println(utils.MagicSum(n))
	fmt.Println(utils.MagicPow(n))
	fmt.Println(utils.MagicOdd(n))
	fmt.Println(utils.MagicGrade(n))
	fmt.Println(utils.MagicName(n))
	fmt.Println(utils.MagicTria(n))
	utils.MagicChange(&n)
	fmt.Println(n)
}
