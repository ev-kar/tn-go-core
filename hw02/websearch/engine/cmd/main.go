package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"../../crawler/pkg/spider"
	"../pkg/search"
)

func main() {
	urls := []string{"https://www.marya.ru/", "https://www.stilkuhni.ru/"}
	data := make(map[string]string)

	for _, url := range urls {
		nextData, err := spider.Scan(url, 3)
		if err != nil {
			log.Printf("ошибка при сканировании сайта %s: %v\n", url, err)
		}

		for k, v := range nextData {
			data[k] = v
		}
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите слово для поиска")
	fmt.Println("---------------------")

	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')
	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)

	result := search.Run(text, data)

	if len(result) == 0 {
		fmt.Println("Не найдено")
		return
	}

	for k, v := range result {
		fmt.Println("---------------------")
		fmt.Println(k)
		fmt.Println(v)
	}
}
