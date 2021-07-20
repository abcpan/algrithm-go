package main

import (
	"abcpan.net.cn/crawl/engine"
	"abcpan.net.cn/crawl/model"
	"abcpan.net.cn/crawl/parsers"
)

const ZhenAiHome = "https://www.zhenai.com"



func main() {
	 engine.Run(&model.Job{
	 	Url: ZhenAiHome + "/zhenghun",
	 	Parse: parsers.ParseCityList,
	 })
}


