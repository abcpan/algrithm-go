package parsers

import (
	"abcpan.net.cn/crawl/model"
	"fmt"
	"regexp"
	"strconv"
)

var NamePattern = regexp.MustCompile(`<h1 class="nickName" data-v-499ba28c>(.*)</h1>`)

var AgePattern = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>(\d+)岁</div>`)

var MarriagePattern = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([未婚|离异|丧偶]+)</div>`)

var HeightPattern = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>(\d+)cm</div>`)

var WeightPattern = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>(\d+)kg</div>`)

var IncomePattern = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>月收入:([^<]+)</div>`)

var EducationPattern = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([高中及以下|大专|中专|大学本科|硕士|博士|博导]+)</div>`)

var HousePattern = regexp.MustCompile(`<div class="m-btn pink" data-v-8b1eac0c>([已购房|未购房]+)</div>`)
var CarPattern = regexp.MustCompile(`<div class="m-btn pink" data-v-8b1eac0c>([未买车|已买车]+)</div>`)


func ParseProfile(content [] byte, nickName string) *model.Result {
 profile := &model.Profile{}

 // 解析姓名

 profile.Name = nickName

 // 解析年龄
 if age, err := strconv.Atoi(seekTargetString(AgePattern, content)); err == nil {
 		profile.Age =  age
 }
 // 解析婚姻
 profile.Marriage = seekTargetString(MarriagePattern, content)

 // 解析身高
 if height, err := strconv.Atoi(seekTargetString(HeightPattern, content)); err == nil {
 		profile.Height = height
 }

 // 解析体重
 if weight, err := strconv.Atoi(seekTargetString(WeightPattern, content)); err == nil {
 	profile.Weight = weight
 }
 // 解析收入
 profile.Income = seekTargetString(IncomePattern, content)

 // 解析教育

 profile.Education = seekTargetString(EducationPattern, content)

 // 解析车房
 profile.Car = seekTargetString(CarPattern, content)

 profile.House = seekTargetString(HousePattern, content)
 fmt.Println("user: %s, profile: %v", nickName, profile)
 result := &model.Result{}
 result.Items = append(result.Items, profile)
 return result
}


// seekTargetString 找到目标字符串
func seekTargetString(pattern *regexp.Regexp, content []byte) string {
	str := string(content)
	matches := pattern.FindStringSubmatch(str)
	if matches != nil {
		ret := matches[1]
		return ret
	}
	return ""
}