package parsers

import "abcpan.net.cn/crawl/model"

// EmptyParser 空的解析器, 返回空的结果
func EmptyParser(contents []byte) *model.Result {
	return &model.Result{}
}