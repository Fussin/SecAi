package reporter

import "fmt"

func FormatForPlatform(report Report, platform string) string {
	switch platform {
	case "hackerone":
		return formatForHackerOne(report)
	case "bugcrowd":
		return formatForBugcrowd(report)
	default:
		return ""
	}
}

func formatForHackerOne(report Report) string {
	fmt.Println("Formatting report for HackerOne...")
	// Placeholder for actual HackerOne formatting logic
	return ""
}

func formatForBugcrowd(report Report) string {
	fmt.Println("Formatting report for Bugcrowd...")
	// Placeholder for actual Bugcrowd formatting logic
	return ""
}
