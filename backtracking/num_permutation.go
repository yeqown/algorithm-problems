package backtracking

import "fmt"

/*
Given a collection of numbers that might contain duplicates, return all possible unique permutations.

Example:
	Input: [1,1,2]
	Output:
	[
		[1,1,2],
		[1,2,1],
		[2,1,1]
	]

Solution: https://www.youtube.com/watch?v=nYFd7VHKyWQ
*/

func permuteUnique(nums []int) [][]int {
	// sort first
	// sort.Ints(nums)
	dict := make(map[int]int)
	for _, num := range nums {
		dict[num]++
	}

	results := new([][]int)

	result := make([]int, len(nums))
	number := make([]int, len(dict))
	count := make([]int, len(dict))

	idx := 0
	for k, v := range dict {
		number[idx] = k
		count[idx] = v
		idx++
	}
	permteHelper(number, count, result, 0, results)
	return *results
}

func permteHelper(numbers, count []int, result []int, level int, results *[][]int) {
	// finnal length of a num slice
	if level == len(result) {
		*results = append(*results, append([]int{}, result...))
		fmt.Printf("%v\n", result)
		return
	}

	for i := 0; i < len(numbers); i++ {
		if count[i] == 0 {
			continue
		}
		result[level] = numbers[i]
		count[i]--
		permteHelper(numbers, count, result, level+1, results)
		count[i]++
	}
}
