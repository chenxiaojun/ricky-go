package main

import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
)

func main() {
	JobList()
}

func JobList(){
	doc, err := goquery.NewDocument("https://www.lagou.com/jobs/list_golang")
	if err != nil {
		panic("err")
	}

	//var r_search []string
	//
	//obj := doc.Find(".r_search")
	//fmt.Println(obj.Html())
	//fmt.Println(obj.Eq(0).Html()) // 10
	//obj.Each(func(i int, s *goquery.Selection){
	//	r_search = append(r_search, s.Text())
	//})
	//fmt.Println(r_search)

	obj := doc.Find("body")
	fmt.Println(obj.Html())
}
