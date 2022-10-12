package util

func CompareEqualSlice(a, b []int) (f bool) {
	flag := true
	for i, v := range a {
		if v != b[i] {
			flag = false
			break
		}
	}
	return flag
}
