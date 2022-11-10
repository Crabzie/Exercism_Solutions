// BenchmarkSieve-2            5860            190007 ns/op           57962 B/op         92 allocs/op
package sieve

func Sieve(limit int) []int {
	marked := make(map[int]bool)
	var result []int
	for i := 2; i <= limit; i++ {
		_, ok := marked[i]
		switch ok {
		case false:
			marked[i] = true
			result = append(result, i)
		case true:
			continue
		}
		j := i
		for ok := marked[i*j]; j <= limit/2 && i*j <= limit && !ok; j++ {
			marked[i*j] = false
		}
	}
	return result
}
