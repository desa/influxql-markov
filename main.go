package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func histogram(doc string) map[string]int {
	s := regexp.MustCompile(" ").Split(doc, -1)
	//s := regexp.MustCompile(" ").Split("123 213 324 234", -1)
	fmt.Println(s)
	return make(map[string]int)
}

func run(starter string, mark map[string][]string, n int) {
	s := starter
	word := starter
	for i := 0; i < n; i++ {
		wordList := mark[word]
		if len(wordList) == 0 {
			break
		}
		randWord := wordList[rand.Intn(len(wordList))]
		s += " " + strings.Split(randWord, " ")[1]
		word = randWord
	}

	fmt.Println(s)
}

func genRandQuery(starter string, mark map[string][]string) string {
	s := starter
	word := starter
	for {
		wordList := mark[word]
		if len(wordList) == 0 || string(word[len(word)-1]) == ";" {
			break
		}
		randWord := wordList[rand.Intn(len(wordList))]
		s += " " + strings.Split(randWord, " ")[1]
		word = randWord
	}

	return s
}

func generateRandQueryN(starters []string, n int, mark map[string][]string) []string {
	queries := make([]string, n)

	for i, _ := range queries {
		queries[i] = genRandQuery(starters[rand.Intn(len(starters))], mark)
	}

	return queries
}

func addTerminalSymbol(doc, sym string) string {
	s := strings.Split(doc, "\n")
	newDoc := ""
	for _, w := range s {
		newDoc += fmt.Sprintf("%v%v\n", w, sym)
	}

	return newDoc
}

func main() {
	url := "https://gist.githubusercontent.com/jackzampolin/201711ff589a7e95f652/raw/904758d9e2e84b0aab8d7f523141ec89e5d90694/refactored_sample_queries"

	resp, err := http.Get(url)
	if err != nil {
		return
	}

	b, _ := ioutil.ReadAll(resp.Body)
	p := addTerminalSymbol(string(b), ";")
	//fmt.Println(p)
	a := regexp.MustCompile(`\s+`).Split(p, 1000)

	for i := 1; i < len(a); i++ {
		a[i-1] = a[i-1] + " " + a[i]
	}

	bi := make(map[string][]string)

	for i := 1; i < len(a); i++ {
		bi[a[i-1]] = append(bi[a[i-1]], a[i])
	}

	hist := make(map[string]int)

	for _, w := range a {
		hist[w]++
	}

	rand.Seed(int64(time.Now().Second()))

	x := generateRandQueryN([]string{"SELECT get", "SELECT put", "SELECT staleness"}, 2, bi)

	for _, w := range x {
		fmt.Println(w)
	}

}
