package sort

func Quick(arr []int) {
	qs(&arr, 0, len(arr)-1)
}

func partition(arr *[]int, lo int, hi int) int {
	pivot := (*arr)[hi]
	idx := lo - 1
	for i := lo; i < hi; i++ {
		if (*arr)[i] <= pivot {
			idx++
			tmp := (*arr)[i]
			(*arr)[i] = (*arr)[idx]
			(*arr)[idx] = tmp
		}
	}
	idx++
	(*arr)[hi] = (*arr)[idx]
	(*arr)[idx] = pivot
	return idx
}

func qs(arr *[]int, lo int, hi int) {
	if lo == hi {
		return
	}
	pivot_idx := partition(arr, lo, hi)
	qs(arr, lo, pivot_idx-1)
	qs(arr, pivot_idx+1, hi)
}
