package method

import (
	"golang.org/x/net/html"
	"net/http"
)

const url = "https://xn--80aesfpebagmfblc0a.xn--p1ai/"

type Data struct {
	Region string `json:"region"`
	Sick   string `json:"sick"`
	Healed string `json:"healed"`
	Died   string `json:"died"`
}

type DataList []Data

func TakeData() DataList {
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	defer res.Body.Close()

	node, err := html.Parse(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var dataList DataList

	table := parse(node)

	for row := table.FirstChild; row != nil; row = row.NextSibling {
		data := Data{
			Region: row.FirstChild.FirstChild.Data,
			Sick:   row.FirstChild.NextSibling.LastChild.Data,
			Healed: row.FirstChild.NextSibling.NextSibling.LastChild.Data,
			Died:   row.FirstChild.NextSibling.NextSibling.NextSibling.LastChild.Data,
		}

		dataList = append(dataList, data)
	}

	return dataList
}

func parse(node *html.Node) *html.Node {
	if node.Data == "div" && node.Attr[0].Val == "d-map__list" {
		return node.FirstChild.FirstChild
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		result := parse(child)
		if result != nil {
			return result
		}
	}

	return nil
}
