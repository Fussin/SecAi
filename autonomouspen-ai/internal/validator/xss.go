package validator

import (
	"fmt"
	"log"

	"github.com/playwright-community/playwright-go"
)

func ValidateXSS(target string, payloads []string) {
	fmt.Printf("Validating XSS for %s...\n", target)

	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}

	for _, payload := range payloads {
		if isVulnerable, err := checkPayload(page, target, payload); err == nil && isVulnerable {
			fmt.Printf("XSS vulnerability found on %s with payload: %s\n", target, payload)
		}
	}

	if err := browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err := pw.Stop(); err != nil {
		log.Fatalf("could not stop playwright: %v", err)
	}
}

func checkPayload(page playwright.Page, target, payload string) (bool, error) {
	isVulnerable := false
	page.On("dialog", func(dialog playwright.Dialog) {
		isVulnerable = true
		dialog.Dismiss()
	})
	_, err := page.Goto(target + payload)
	return isVulnerable, err
}
