package core

import "fmt"

type Severity string

const (
	Critical Severity = "Critical"
	High     Severity = "High"
	Medium   Severity = "Medium"
	Low      Severity = "Low"
)

func ClassifySeverity(vulnerabilityType string) Severity {
	fmt.Printf("Classifying severity for %s...\n", vulnerabilityType)

	switch vulnerabilityType {
	case "rce-system":
		return Critical
	case "sqli-admin":
		return Critical
	case "auth-bypass-admin":
		return Critical
	case "ssrf-cloud-metadata":
		return Critical
	case "rce-limited":
		return High
	case "sqli-user":
		return High
	case "xss-stored-critical":
		return High
	case "xxe-file-disclosure":
		return High
	case "xss-reflected-interaction":
		return Medium
	case "idor-limited":
		return Medium
	case "open-redirect":
		return Medium
	default:
		return Low
	}
}
