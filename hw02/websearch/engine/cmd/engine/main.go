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
	Scan(url string, depth int) (map[string]string, error)
}

type siteScanner func(string, int)

func (s siteScanner) Scan(url string, depth int) (data map[string]string, err error) {
	data, err = spider.Scan(url, depth)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func run(s Scanner, urls []string, depth int) (map[string]string, error) {
	data := make(map[string]string)
	for _, u := range urls {
		nextData, err := s.Scan(u, depth)
		if err != nil {
			return nil, err
		}
		for k, v := range nextData {
			data[k] = v
		}
	}
	return data, nil
}

func main() {
	urls := []string{"https://go.dev/", "https://golang.org/"}
	const depth int = 2
	var s siteScanner
	data, err := run(s, urls, depth)
	if err != nil {
		fmt.Println("ошибка при сканировании сайта", err)
		return
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
