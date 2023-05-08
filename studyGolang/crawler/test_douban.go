package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const baseURL = "https://movie.douban.com/top250" // 定义要爬取的页面URL
const userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"

// 获取页面内容
func getPageData(url string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", userAgent)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			return "", err
		}
		html, err := doc.Html()
		if err != nil {
			return "", err
		}
		return html, nil
	} else {
		return "", fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}
}

// 解析页面内容，获取所需信息
func parsePageData(html string) ([]map[string]string, error) {
	movieData := []map[string]string{}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return movieData, err
	}
	doc.Find(".article .grid_view li").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".hd a span:first-of-type").Text()
		rating := s.Find(".rating_num").Text()
		ratingNum := s.Find(".star span:last-of-type").Text()
		ratingNum = strings.Trim(ratingNum, "人评价")
		movieInfo := map[string]string{"title": title, "rating": rating, "rating_num": ratingNum}
		movieData = append(movieData, movieInfo)
	})
	return movieData, nil
}

// 处理一页数据
func handleOnePage(pageNum int) ([]map[string]string, error) {
	pageURL := fmt.Sprintf("%s?start=%d&filter=", baseURL, pageNum*25)
	html, err := getPageData(pageURL)
	if err != nil {
		return nil, err
	}
	movieData, err := parsePageData(html)
	if err != nil {
		return nil, err
	}
	return movieData, nil
}

// 程序主函数
func main() {
	movieData := []map[string]string{}
	for i := 0; i < 10; i++ {
		movieDataPerPage, err := handleOnePage(i)
		if err != nil {
			fmt.Println(err)
			continue
		}
		movieData = append(movieData, movieDataPerPage...)
		time.Sleep(time.Duration(1+rand.Float64()*1) * time.Second)
	}
	// 按评分重新排序
	sort.Slice(movieData, func(i, j int) bool {
		rating1, _ := strconv.ParseFloat(movieData[i]["rating"], 64)
		rating2, _ := strconv.ParseFloat(movieData[j]["rating"], 64)
		return rating1 > rating2
	})
    fmt.Println("---------下面是为您爬取的豆瓣电影Top250---------")
	// 控制输出电影名称、评分和评分人数，只显示前250部
	for i, movie := range movieData[:250] {
		fmt.Printf("%d: %s，评分：%s，评分人数：%s\n", i+1, movie["title"], movie["rating"], movie["rating_num"]+"人")
	}
}