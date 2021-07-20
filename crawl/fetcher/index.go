package fetcher

import (
	"bufio"
	"errors"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"io/ioutil"
	"log"
	"net/http"
)

const UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36"

// Fetch 获取内容
func Fetch(url string)([]byte, error) {

	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Printf("create request error, %v", err)
		return nil, err
	}

	request.Header.Set("User-Agent", UserAgent)
	res, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		msg := fmt.Sprintf("wrong status code: %d", res.StatusCode )
		return nil, errors.New(msg)
	}


	bytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}
	return bytes, nil
}


// FindDecoderFactory 根据内容获取解码器工厂
func createDecoding(reader *bufio.Reader) encoding.Encoding {
	bytes, err := reader.Peek(1024)
	if err != nil {
		log.Printf("get Encoding failed: %v", err)
		return unicode.UTF8
	}

	encoding, _, _ := charset.DetermineEncoding(bytes, "")
	return encoding
}