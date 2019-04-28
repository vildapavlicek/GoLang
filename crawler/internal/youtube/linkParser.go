package youtube

import (
	"errors"
	"net/http"

	"golang.org/x/net/html"
)

// Parse parses youTube html for next video sufix and title
func Parse(res *http.Response) (title, link string, err error) {
	defer res.Body.Close()
	doc, err := html.Parse(res.Body)
	if err != nil {

		return "", "", err
	}
	title, link, err = parseNode(doc)

	if err != nil {

	}

	return title, link, err
}

func parseNode(n *html.Node) (title, link string, err error) {
	if n.Type == html.ElementNode && n.Data == "ul" {
		for _, v := range n.Attr {
			if v.Key == "class" && v.Val == "video-list" {
				title, link, err = parseNextLink(n)
				return title, link, err
			}
		}

	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if title == "" || link == "" {
			title, link, err = parseNode(c)
		}
	}

	return title, link, err
}

func parseNextLink(n *html.Node) (title, link string, err error) {

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, v := range n.Attr {
			if v.Key == "href" {
				link = v.Val
			}

			if v.Key == "title" {
				title = v.Val
			}

			if title != "" && link != "" {
				return title, link, nil
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {

		if title == "" || link == "" {
			title, link, err = parseNextLink(c)
		}
	}

	if title == "" || link == "" {
		return title, link, errors.New("Failed to parse title or link")
	}
	return title, link, nil

}
