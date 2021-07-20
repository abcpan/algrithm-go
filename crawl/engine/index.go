package engine

import (
	fetcher2 "abcpan.net.cn/crawl/fetcher"
	"abcpan.net.cn/crawl/model"
	"log"
)

func Run(seeds ...*model.Job) {
	var jobQueue []*model.Job

	for _, job := range seeds {
		jobQueue = append(jobQueue, job)
	}

	for len(jobQueue) > 0 {
		job := jobQueue[0]
		jobQueue = jobQueue[1:]
		log.Printf("fetching: %v", (*job).Url)
		bytes, err := fetcher2.Fetch((*job).Url)
		if err != nil {
			log.Printf("job error: fetch url %s", (*job).Url)
			continue
		}

		result := (*job).Parse(bytes)
		jobQueue = append(jobQueue, result.Jobs...)

		for _, item := range result.Items {
			log.Printf("Got item: %v", item)
		}
	}
}
