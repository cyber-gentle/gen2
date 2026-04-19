package gen2

func Sum(n ...int) int {
	var total int

	for i := range n {
		total += n[i]
	}
	return total
}

func Difference(n ...int) int {
	var total int = n[0]

	for i := 1; i <= len(n)-1; i++ {
		total -= n[i]
	}
	return total
}
