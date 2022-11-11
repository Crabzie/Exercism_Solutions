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
		for j := i; j <= limit/i; j++ {
			ok := marked[i*j]
			if ok {
				continue
			}
			marked[i*j] = false
		}
	}
	return result
}
