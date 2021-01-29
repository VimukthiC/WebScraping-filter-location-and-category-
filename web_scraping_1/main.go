package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {

	var location string
	fmt.Print("Enter Location : ")
	fmt.Scan(&location)

	var catagory string
	fmt.Print("Enter Catagory : ")
	fmt.Scan(&catagory)

	c := colly.NewCollector(
		colly.AllowedDomains("ikman.lk"),
	)

	c.OnHTML(".list--3NxGO", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	c.Visit("https://ikman.lk/en/ads/" + location + "/" + catagory)

}
