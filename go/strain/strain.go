package strain

// BenchmarkKeepInts-2              2592338               414.6 ns/op           104 B/op          7 allocs/op
// BenchmarkDiscardInts-2           2457949               436.9 ns/op           120 B/op          7 allocs/op
type Ints []int
type Lists [][]int
type Strings []string

func keepSlice[T any](i []T, filter func(T) bool) (result []T) {
	for _, v := range i {
		if filter(v) {
			result = append(result, v)
		}
	}
	return
}
func discardSlice[T int](i []T, filter func(T) bool) []T {
	return keepSlice(i, func(x T) bool { return !filter(x) })
}
func (i Ints) Keep(filter func(int) bool) Ints {
	return keepSlice(i, filter)
}
func (i Ints) Discard(filter func(int) bool) Ints {
	return discardSlice(i, filter)
}
func (l Lists) Keep(filter func([]int) bool) Lists {
	return keepSlice(l, filter)
}
func (s Strings) Keep(filter func(string) bool) Strings {
	return keepSlice(s, filter)
}
