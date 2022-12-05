package series

func All(n int, s string) (result []string) {
	for i := range s {
		if i+n > len(s) {
			break
		}
		result = append(result, s[i:i+n])
	}
	return
}

func UnsafeFirst(n int, s string) string {
	return All(n, s)[0]
}

/* bonus exercise */
func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return
	}
	return UnsafeFirst(n, s), true
}
