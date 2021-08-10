package isIsomorphic

// 是否是同构字符串
func isIsomorphic(s1 string, s2 string) bool {
	chars1 := []byte(s1)
	chars2 := []byte(s2)
	memo1 := make(map[byte]int)
	memo2 := make(map[byte]int)
	size := len(chars2)
	for i := 0; i < size;i++ {
		c1 := chars1[i]
		c2 := chars2[i]
		if memo1[c1] != memo2[c2] {
			return false
		}
		memo1[c1] = i+1
		memo2[c2] = i+1
	}
	return true
}