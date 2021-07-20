package subsets


type Solution struct {
	result [][]int
	temp [] int
}


func (this *Solution)subsets(nums []int) [][]int{
	this.backTrack(nums, 0)
	return this.result
}


// backTrack 进行回溯
func (this *Solution) backTrack(source []int, start int)  {
	// 将子集添加进结果集
	ret := make([]int, len(this.temp))
	copy(ret, this.temp)
	this.result = append(this.result, ret)

	for i := start; i < len(source);i++ {
		this.temp = append(this.temp, source[i])
		this.backTrack(source, i + 1)
		this.temp = this.temp[:len(this.temp)-1]
	}
}

func subsets(nums []int) [][]int{
	solution := &Solution{}
	return solution.subsets(nums)
}