package bugbounty

import (
	"fmt"
	"log"
	"os"
)

type Platform struct {
	Name      string
	APIKey    string
	RateLimit int
}

func InitializePlatforms() []Platform {
	platforms := []Platform{
		{Name: "HackerOne", APIKey: os.Getenv("HACKERONE_API_KEY"), RateLimit: 60},
		{Name: "Bugcrowd", APIKey: os.Getenv("BUGCROWD_API_KEY"), RateLimit: 0},
		{Name: "Intigriti", APIKey: os.Getenv("INTIGRITI_API_KEY"), RateLimit: 0},
		{Name: "YesWeHack", APIKey: os.Getenv("YESWEHACK_API_KEY"), RateLimit: 0},
	}

	for _, p := range platforms {
		if p.APIKey == "" {
			log.Printf("API key for %s not found. Skipping.", p.Name)
			continue
		}
		fmt.Printf("Initializing %s integration...\n", p.Name)
		// testConnectivity(p)
		// fetchRateLimit(p)
	}

	return platforms
}
