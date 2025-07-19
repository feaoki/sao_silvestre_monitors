package scraper

import (
	"strings"

	"github.com/gocolly/colly"
)

type CollyChecker struct {
	URL string
}

func NewCollyChecker(url string) *CollyChecker {
	return &CollyChecker{URL: url}
}

func (c *CollyChecker) Check() (bool, error) {
	open := false
	collector := colly.NewCollector()

	collector.OnHTML("body", func(e *colly.HTMLElement) {
		if strings.Contains(strings.ToLower(e.Text), "inscrições abertas") {
			open = true
		}
	})

	err := collector.Visit(c.URL)
	return open, err
}
