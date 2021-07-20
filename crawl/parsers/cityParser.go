package parsers

import (
	"abcpan.net.cn/crawl/model"
	"regexp"
)
// 列表单项匹配
var CityMatcher = regexp.MustCompile( `<div class="list-item">(.*)</div>`)

var PhotoNameMatcher = regexp.MustCompile(`<img src="([^>]+)" alt="([^>]+)">`)

// ParseCity 解析城市下的用户
func ParseCity(contents []byte) *model.Result{
	str := string(contents)

	matches := CityMatcher.FindAllStringSubmatch(str, -1)

	result := &model.Result{}

	for _, match := range matches {
		 PhotoNameMatcher.FindAllStringSubmatch(match[1], -1)

	}

	return result
}