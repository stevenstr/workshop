package pipeline

import (
	"bufio"
	"io"
	"net/http"
)

func SearchWord(urls <-chan string, word string) int {
	// initialize structure to store result

	// run routines to fetch document and count words

	//summarize and store result

	return 271
}

func fetchDocument(url string) io.ReadCloser {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	return resp.Body
}

func countWords(r io.ReadCloser, word string) int {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	var count int
	for sc.Scan() {
		if sc.Text() == word {
			count++
		}
	}

	return count
}