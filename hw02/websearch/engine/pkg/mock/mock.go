package mock

type site struct {
	url   string
	depth int
}

func (s *site) Scan() (map[string]string, error) {
	data := map[string]string{
		"https://go.dev/":                       "go.dev",
		"https://go.dev/solutions":              "Why Go - go.dev",
		"https://go.dev/about":                  "About - go.dev",
		"https://go.dev/solutions#use-cases":    "Why Go - go.dev",
		"https://go.dev/solutions#case-studies": "Why Go - go.dev",
		"https://go.dev/learn":                  "Getting Started - go.dev",
		"https://golang.org/":                   "The Go Programming Language",
	}
	return data, nil
}
