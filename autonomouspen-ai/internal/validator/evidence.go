package validator

import (
	"fmt"
	"log"

	"github.com/playwright-community/playwright-go"
)

func CollectVisualEvidence(target string, exploit func(playwright.Page)) {
	fmt.Printf("Collecting visual evidence for %s...\n", target)

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

	// Record video
	page.Video().Path()

	// Before screenshot
	page.Screenshot(playwright.PageScreenshotOptions{
		Path: playwright.String("before.png"),
	})

	// Exploit
	exploit(page)

	// After screenshot
	page.Screenshot(playwright.PageScreenshotOptions{
		Path: playwright.String("after.png"),
	})

	if err := browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err := pw.Stop(); err != nil {
		log.Fatalf("could not stop playwright: %v", err)
	}
}
