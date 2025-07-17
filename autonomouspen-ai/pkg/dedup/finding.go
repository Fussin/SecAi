package dedup

import (
	"crypto/sha256"
	"fmt"
)

type Finding struct {
	Endpoint    string
	Parameter   string
	Type        string
	Impact      string
	Hash        string
}

func DeduplicateFindings(findings []Finding) []Finding {
	fmt.Println("Deduplicating findings...")
	var deduplicatedFindings []Finding
	hashes := make(map[string]bool)

	for _, finding := range findings {
		finding.Hash = calculateFindingHash(finding)
		if !hashes[finding.Hash] {
			hashes[finding.Hash] = true
			deduplicatedFindings = append(deduplicatedFindings, finding)
		}
	}

	return deduplicatedFindings
}

func calculateFindingHash(finding Finding) string {
	data := fmt.Sprintf("%s-%s-%s-%s", finding.Endpoint, finding.Parameter, finding.Type, finding.Impact)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(data)))
}

func recognizeSimilarPatterns(findings []Finding) {
	fmt.Println("Recognizing similar patterns...")
	// Placeholder for actual similar pattern recognition logic
}
