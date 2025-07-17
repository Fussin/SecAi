package core

import "fmt"

type ScopeFinding struct {
	Finding
	Priority string
}

func GroupByScope(findings []Finding) []ScopeFinding {
	fmt.Println("Grouping findings by scope...")
	var scopeFindings []ScopeFinding

	for _, finding := range findings {
		scopeFindings = append(scopeFindings, ScopeFinding{
			Finding:  finding,
			Priority: determinePriority(finding.Endpoint),
		})
	}

	return scopeFindings
}

func determinePriority(endpoint string) string {
	fmt.Printf("Determining priority for endpoint %s...\n", endpoint)
	// Placeholder for actual priority determination logic
	return "Medium"
}
