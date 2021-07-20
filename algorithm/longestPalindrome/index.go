package longestPalindrome

import (
	"math"
	"strings"
)

func longestPalindrome(s string) string {
	if  s == "" {
		return ""
	}
	var start, end = 0, 0
	var maxLen = 0
	strs := strings.Split(s, "")

	for i := 0; i < len(strs); i++ {
	  length1, left1, right1 := expandAroundCenter(strs, i, i)
	  length2, left2, right2 := expandAroundCenter(strs, i, i + 1)
	  length := int(math.Max(float64(length1), float64(length2)))

		if length >= maxLen {
			 maxLen = length
			 if length == length1 {
				 start = left1
				 end = right1
			 }else {
			 	start = left2
			 	end = right2
			 }

		}
	}

	ss := strs[start: end + 1]
	return strings.Join(ss, "")
}

// 中心拓展, 返回符合要求的长度
func expandAroundCenter(s []string, left int, right int) (int, int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	left++
	right--
	return right - left + 1, left, right
}