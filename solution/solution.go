package solution

import (
	"sort"
)

func Solution(A []int) int {
	var positiveNumbers []int
	for _, number := range A {
		if number > 0 {
			positiveNumbers = append(positiveNumbers, number)
		}
	}
	if len(positiveNumbers) == 0 {
		return 1
	}
	sort.Ints(positiveNumbers)
	uniqueNumbers := removeDuplicateValues(positiveNumbers)
	majorNumber := 1
	for i := range uniqueNumbers {
		if majorNumber == uniqueNumbers[i] {
			majorNumber = uniqueNumbers[i] + 1
		} else {
			break
		}

	}
	return majorNumber
}

// removeDuplicateValues taken from https://www.geeksforgeeks.org/how-to-remove-duplicate-values-from-slice-in-golang/
func removeDuplicateValues(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}

	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
