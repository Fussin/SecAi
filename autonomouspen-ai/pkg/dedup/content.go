package dedup

import (
	"fmt"

	"github.com/mfonda/simhash"
)

func ContentDeduplication(targets []string) {
	fmt.Println("Performing content-based deduplication...")
	hashes := make(map[string]uint64)

	for _, t := range targets {
		// htmlContent := fetchHTML(t)
		hashes[t] = simhash.Simhash(simhash.NewWordFeatureSet([]byte("")))
	}

	for t1, h1 := range hashes {
		for t2, h2 := range hashes {
			if t1 != t2 && simhash.Compare(h1, h2) > 85 {
				fmt.Printf("%s is similar to %s\n", t1, t2)
			}
		}
	}
}
