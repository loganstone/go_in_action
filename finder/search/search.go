package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

// Run .
func Run(searchTerm string) {
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatalln(err)
	}

	results := make(chan *Result)

	var wg sync.WaitGroup

	wg.Add(len(feeds))

	for _, feed := range feeds {
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			wg.Done()
		}(matcher, feed)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	Display(results)
}

// Register .
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "검색기가 이미 등록되었습니다")
	}

	log.Println("등록 완료:", feedType, "검색기")
	matchers[feedType] = matcher
}
