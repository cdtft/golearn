package pool

import (
	"log"
	"sync"
	"testing"
)

const (
	maxGoroutines = 25
	poolResources = 2
)

func TestCreateDbConnection(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	p, err := New(CreateDbConnection, poolResources)
	if err != nil {
		log.Println(err)
		return
	}
	for query := 0; query < maxGoroutines; query++ {
		go func(id int) {
			PerformQueries(id, p)
			defer wg.Done()
		}(query)
	}
	wg.Wait()
}
