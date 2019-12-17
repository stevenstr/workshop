package arraysum

func Sum(arr []int) int {
	var sum int
	for _, d := range arr {
		sum += d
	}

	return sum
}

func SumParallel(arr []int) int {
	//split array to n parts, n equals to logic cores count on current machine

	//initialize structure for sub results or one global what you want

	//run parallel computing for all parts

	//wait unitl all routines end computing

	//summarize and return results

	return 0
}
