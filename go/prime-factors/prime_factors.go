package prime

func numbers(n int64, result *[]int64) int64 {
	for i := 2; i <= int(n) && n >= 1; i++ {
		if n%int64(i) == 0 {
			*result = append(*result, int64(i))
			n = numbers(n/int64(i), result)
		}
	}
	return 0
}
func Factors(n int64) (result []int64) {
	numbers(n, &result)
	return
}
