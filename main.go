package main

import (
	"fmt"

	"github.com/julian-alarcon/codility-sample-test/solution"
)

func main() {
	fmt.Println("Just some testResults")
	testResult := solution.Solution([]int{1, 3, 6, 4, 1}) // 2
	fmt.Println(testResult)
	testResult = solution.Solution([]int{1, 3, 6, 4, 1, 2}) // 5
	fmt.Println(testResult)
	testResult = solution.Solution([]int{1, 1, 1, 4, 1, 2}) // 3
	fmt.Println(testResult)
	testResult = solution.Solution([]int{1, 1, 2, 2, 1, 2}) // 3
	fmt.Println(testResult)
	testResult = solution.Solution([]int{1, 2, 3}) // 4
	fmt.Println(testResult)
	testResult = solution.Solution([]int{-1, -32, 3, 5}) // 1
	fmt.Println(testResult)
	testResult = solution.Solution([]int{-1, -5, -5, 1}) // 2
	fmt.Println(testResult)
	testResult = solution.Solution([]int{3, 3, 5}) // 1
	fmt.Println(testResult)
}
