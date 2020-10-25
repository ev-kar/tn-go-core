package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"../../../crawler/pkg/spider"
)

// Scanner interface
type Scanner interface {
	Scan() (map[string]string, error)
}

type site struct {
	url   string
	depth int
}

func (s *site) Scan() (map[string]string, error) {
	return spider.Scan(s.url, s.depth)
}

func main() {
	mySites := []site{
		{url: "https://go.dev/", depth: 3},
		{url: "https://golang.org/", depth: 3},
	}

	data := make(map[string]string)
	for _, mySite := range mySites {
		nextData, err := mySite.Scan()
		if err != nil {
			fmt.Println("ошибка при сканировании сайта", mySite.url, err)
			return
		}
		for k, v := range nextData {
			data[k] = v
		}
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите слово для поиска(вЫход):")
	fmt.Print("-> ")
	for scanner.Scan() {
		text := scanner.Text()
		if text == "Ы" {
			return
		}
		result := make(map[string]string)
		for k, v := range data {
			if strings.Contains(v, text) {
				result[k] = v
			}
		}
		if len(result) == 0 {
			fmt.Println("---------------------")
			fmt.Println("Не найдено")
		}
		for k, v := range result {
			fmt.Println("---------------------")
			fmt.Println(k)
			fmt.Println(v)
		}
		fmt.Println("---------------------")
		fmt.Println("Введите слово для поиска(вЫход):")
		fmt.Print("-> ")
	}
}
