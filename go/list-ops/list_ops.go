package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	if s.Length() != 0 {
		result := initial
		for _, v := range s {
			result = fn(result, v)
		}
		return result
	}
	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	if s.Length() != 0 {
		result := initial
		for i := s.Length() - 1; i >= 0; i-- {
			result = fn(s[i], result)
		}
		return result
	}
	return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
	if s.Length() != 0 {
		result := make(IntList, s.Length())
		var i int
		for _, v := range s {
			if fn(v) {
				result[i] = v
				i++
			}
		}
		return result[:i]
	}
	return []int{}
}
func (s IntList) Length() int {
	var result int
	for range s {
		result++
	}
	return result
}

func (s IntList) Map(fn func(int) int) IntList {
	result := make(IntList, s.Length())
	if s.Length() != 0 {
		for i, v := range s {
			result[i] = fn(v)
		}
	}
	return result
}

func (s IntList) Reverse() IntList {
	result := make(IntList, s.Length())
	for i, j := 0, s.Length()-1; i < s.Length(); i, j = i+1, j-1 {
		result[i] = s[j]
	}
	return result
}

func (s IntList) Append(lst IntList) IntList {
	result := make(IntList, s.Length()+lst.Length())
	if s.Length() == 0 {
		copy(result, lst)
		return result
	}
	copy(result, s)
	for i, j := s.Length(), 0; i < len(result); i, j = i+1, j+1 {
		result[i] = lst[j]
	}
	return result
}
func (s IntList) Concat(lists []IntList) IntList {
	var result IntList
	result = result.Append(s)
	for _, v := range lists {
		if v.Length() != 0 {
			result = result.Append(v)
		}
	}
	return result
}
