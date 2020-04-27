package webcrawler

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gmdutra/jowfuzz/pkg/collectlinks"
	"github.com/gmdutra/jowfuzz/pkg/parameters"
	"github.com/gmdutra/jowfuzz/pkg/requests"
)

var newLinks []string

func GetUrls(host string) (links []string, err error) {
	req, err := requests.Get(host)
	links = collectlinks.All(req.Body)

	return links, err
}

func checkBadUrl(host string) bool {
	badUrls := []string{"https://www.facebook.com", "instagram.com", "twitter.com"}

	for _, badUrl := range badUrls {
		check := strings.Contains(host, badUrl)

		if check {
			return true
		}
	}

	return false
}

func checkLink(host string, str string) (string, bool) {

	checkUrl, _ := regexp.MatchString("http.*://", str)

	if checkUrl {
		badUrl := checkBadUrl(host)
		if badUrl == true {
			return "", false
		}
	}

	//eturn host + "/" + str, true
	return str, true
}

func InsertLink(host string, link string) {
	newLink, checkLink := checkLink(host, link)

	if collectlinks.Check(newLinks, newLink) == false {
		if checkLink {
			fmt.Println(newLink)
			newLinks = append(newLinks, newLink)
		}
	}
}

func RunCrawler(host string) []string {
	var links []string
	links, _ = GetUrls(host)

	for _, link := range links {
		InsertLink(host, link)

		for _, newLink := range newLinks {
			links, _ = GetUrls(newLink)

			for _, link := range links {
				InsertLink(host, link)
			}
		}
	}

	return newLinks
}

func Crawler(parameters *parameters.HostParameters) {
	host := parameters.Host
	links := RunCrawler(host)

	for _, link := range links {
		fmt.Println(link)
	}

}
