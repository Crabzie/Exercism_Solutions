package flatten

func Flatten(nested interface{}) []interface{} {
	var result = make([]interface{}, 0)
	switch t := nested.(type) {
	case []interface{}:
		for _, v := range t {
			result = append(result, Flatten(v)...)
		}
	case interface{}:
		result = append(result, t)
	}
	return result
}
