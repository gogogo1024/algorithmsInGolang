package introduction

func binarySearch(array []int, search int) (find bool, index int) {
	low := 0
	high := len(array) - 1
	for low <= high {
		mid := (low + high) / 2
		guess := array[mid]
		if guess == search {
			return true, mid
		}
		if search < guess {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false, -1
}
