package arrays

func SumAllTails(arraysToSum ...[]int) []int {
	var sums []int

	for _, numbers := range arraysToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
