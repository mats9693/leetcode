package mario

func maximumGap(nums []int) int {
	QuickSort(nums)

	max := 0
	for i := 1; i < len(nums); i++ {
		if max < nums[i]-nums[i-1] {
			max = nums[i] - nums[i-1]
		}
	}

	return max
}

func QuickSort(is []int) {
	if len(is) <= 1 {
		return
	}

	var small int
	{
		var big int
		for i := 1; i < len(is); i++ {
			if is[i] > is[0] {
				big++
			} else {
				small++
				if big != 0 {
					is[i], is[small] = is[small], is[i]
				}
			}
		}
		if small != 0 {
			is[0], is[small] = is[small], is[0]
		}
	}

	QuickSort(is[:small])
	QuickSort(is[small+1:])

	return
}
