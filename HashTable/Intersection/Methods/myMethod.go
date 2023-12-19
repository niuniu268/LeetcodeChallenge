package Methods

func Intersection(nums1 []int, nums2 []int) []int {
	arr1 := make([]int, 10)

	for i := range nums1 {
		arr1[nums1[i]-0]++
	}

	arr2 := make([]int, 10)

	for i := range nums2 {
		arr2[nums2[i]-0]++
	}

	arr := make([]int, 0)

	for i := 0; i < 10; i++ {
		if arr1[i] > 0 && arr2[i] > 0 {
			arr = append(arr, i)
		}
	}

	return arr

}
