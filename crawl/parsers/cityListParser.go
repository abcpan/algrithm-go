package parsers

import (
	"abcpan.net.cn/crawl/model"
	"fmt"
	"regexp"
)

var CityListMatcher = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)" [^>]*>([^<]+)</a>`)

// ParseCityList 处理内容
func ParseCityList(contents []byte) *model.Result {

	matches := CityListMatcher.FindAllStringSubmatch(string(contents), -1)

	result := &model.Result{}
	// 限制请求数量
	var limit = 10
	for _, match := range matches {
		// 新建内容
		item := fmt.Sprintf("City %s", match[2])
		result.Items = append(result.Items, item)


		// 新建任务
		url :=  match[1]
		job := &model.Job{ Url: url, Parse: ParseCity }

		result.Jobs = append(result.Jobs, job)
		limit--
		if limit == 0 {
			break
		}
	}

	return result
}