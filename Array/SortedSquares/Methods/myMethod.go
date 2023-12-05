package Methods

func Mymethod(nums []int) []int {

	var tmp int
	var tmp2 int
	var count int

	var Nnums = make([]int, len(nums), len(nums))

	for i := 0; i < len(nums); i++ {

		tmp2 = nums[i] * nums[i]
		Nnums[i] = tmp2

	}

	for count < len(Nnums) {

		for i := 0; i < len(Nnums)-1; i++ {
			count++

			if Nnums[i] > Nnums[i+1] {
				tmp = Nnums[i+1]
				Nnums[i+1] = Nnums[i]
				Nnums[i] = tmp
			}
		}

	}

	return Nnums
}
