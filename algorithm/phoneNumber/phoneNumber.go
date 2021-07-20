package phoneNumber

import "strings"

var phoneMap = map[string] []string {
	"2": strings.Split("abc", ""),
	"3": strings.Split("def", ""),
	"4": strings.Split("ghi", ""),
	"5": strings.Split("jkl", ""),
	"6": strings.Split("mno", ""),
	"7": strings.Split("pqrs", ""),
	"8": strings.Split("tuv", ""),
	"9": strings.Split("wxyz", ""),
}

type PhoneSolution struct {
	temp string
	result [] string
}
//
func(this *PhoneSolution) letterCombinations(digests string) [] string {
	if digests == "" {
		return this.result
	}
	strs := strings.Split(digests, "")
	this.backTrack(strs, 0)
	return this.result
}

// 回溯
func (this *PhoneSolution) backTrack(strs []string, start int) {
	if start == len(strs) {
		this.result = append(this.result, this.temp)
		return
	}

	letters :=phoneMap[strs[start]]

	for i := 0; i < len(letters); i++ {
		this.temp += letters[i]
		this.backTrack(strs, start + 1)
		this.temp = strings.TrimSuffix(this.temp, letters[i])
	}
}

func letterCombinations(digests string) [] string {
	solution := &PhoneSolution{}
	return solution.letterCombinations(digests)
}