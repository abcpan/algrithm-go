package model

// Job 解析任务
type Job struct {
	Url string
	Parse func(content []byte) *Result
}

// Result 解析结果
type Result struct {
	Jobs []*Job
	Items [] interface{}
}