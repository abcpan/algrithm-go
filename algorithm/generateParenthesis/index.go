package generateParenthesis


type Parenthesis struct {
	limit int
	result []string
	path string
}


func (this *Parenthesis) solve() []string {
	if this.limit == 0 {
		return this.result
	}
	this.backTrack("", 0, 0)
	return this.result
}

func (this *Parenthesis) backTrack(str string, left int, right int)  {
	//
	if left == this.limit && right == this.limit {
		this.result = append(this.result, str)
		return
	}
	// 剪枝操作
	if left < right {
		return
	}

	if left < this.limit {
		this.backTrack(str + "(", left + 1, right)
	}
	if right < this.limit {
		this.backTrack(str + ")",left, right + 1)
	}
}

func generateParenthesis(n int) []string {
  solution := &Parenthesis{limit: n}
  return solution.solve()
}