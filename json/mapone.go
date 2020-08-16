package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"strings"
)

type PublisherBrowserReport struct {

	Id int `json:"id"`
	Browser string `json:"browser"`
	Impressions int `json:"impressions"`
	Inventory int `json:"inventory"`
	//Clicks int64 `json:"clicks"`
	//ImpressionCost float64 `json:"impression_cost"`
	//ClickCost float64 `json:"click_cost"`
	//Conversion int `json:"conversion"`
	//Bidloss int64 `json:"bidloss"`
	//Dates time.Time `json:"dates"`
	SiteId int `json:"site_id"`
}



func main(){

	data := []*PublisherBrowserReport{
		{   1,
			"chrome",
			22,
			500,
			1,
		},
		{   2,
			"chrome",
			23,
			600,
			2,
		},
		{   3,
			"chrome",
			24,
			700,
			1,
		},
		{   4,
			"chrome",
			25,
			800,
			2,
		},
		{   5,
			"opera",
			25,
			800,
			2,
		},
		{   6,
			"opera",
			25,
			800,
			1,
		},
	}

	pubdata := make(map[int]map[string]PublisherBrowserReport)

	for _, val := range data{
		sitedata, ok := pubdata[val.SiteId]
		spew.Dump(sitedata)
		fmt.Println(strings.Repeat("#",4))
		sitedata = make(map[string]PublisherBrowserReport)
		spew.Dump(sitedata)
		fmt.Println("#######")
		if !ok{
			sitedata = make(map[string]PublisherBrowserReport)
			spew.Dump(sitedata)
			pubdata[val.SiteId] = sitedata
		}
		//browserData, _ := pubdata[val.SiteId][val.Browser]
		//browserData.Impressions++
		//pubdata[val.SiteId][val.Browser] = browserData
		fmt.Println(strings.Repeat("=",14))
	}
	//spew.Dump(pubdata)
}