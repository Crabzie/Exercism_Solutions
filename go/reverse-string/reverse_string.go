// BenchmarkReverse-2        552910              1945 ns/op             176 B/op         34 allocs/op
package reverse

func Reverse(input string) (result string) {
	for _, v := range input {
		result = string(v) + result
	}
	return
}
