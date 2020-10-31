package mock

// SiteScanner - сканер-заглушка
type SiteScanner func(string, int)

// Scan - метод-заглушка, возвращающий жестко закодированные данные
func (s SiteScanner) Scan(url string, depth int) (data map[string]string, err error) {
	data = map[string]string{
		"https://go.dev/":     "go.dev",
		"https://golang.org/": "The Go Programming Language",
	}
	return data, nil
}
