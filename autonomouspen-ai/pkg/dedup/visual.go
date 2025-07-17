package dedup

import (
	"fmt"
	"image"
	_ "image/png"
	"os"

	"github.com/nfnt/resize"
)

func VisualDeduplication(targets []string) {
	fmt.Println("Performing visual deduplication...")
	hashes := make(map[string]string)

	for _, t := range targets {
		// screenshotPath := captureScreenshot(t)
		hashes[t] = pHash("dummy.png")
	}

	for t1, h1 := range hashes {
		for t2, h2 := range hashes {
			if t1 != t2 && hammingDistance(h1, h2) < 5 {
				fmt.Printf("%s is visually similar to %s\n", t1, t2)
			}
		}
	}
}

func pHash(imagePath string) string {
	file, err := os.Open(imagePath)
	if err != nil {
		return ""
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return ""
	}

	// 1. Resize image to 8x8
	smallImg := resize.Resize(8, 8, img, resize.Lanczos3)

	// 2. Convert to grayscale
	grayImg := image.NewGray(smallImg.Bounds())
	for y := 0; y < smallImg.Bounds().Max.Y; y++ {
		for x := 0; x < smallImg.Bounds().Max.X; x++ {
			grayImg.Set(x, y, smallImg.At(x, y))
		}
	}

	// 3. Calculate the average pixel value
	var total uint32
	for y := 0; y < grayImg.Bounds().Max.Y; y++ {
		for x := 0; x < grayImg.Bounds().Max.X; x++ {
			total += uint32(grayImg.GrayAt(x, y).Y)
		}
	}
	avg := total / 64

	// 4. Create the hash
	var hash string
	for y := 0; y < grayImg.Bounds().Max.Y; y++ {
		for x := 0; x < grayImg.Bounds().Max.X; x++ {
			if uint32(grayImg.GrayAt(x, y).Y) > avg {
				hash += "1"
			} else {
				hash += "0"
			}
		}
	}

	return hash
}

func hammingDistance(s1, s2 string) int {
	var distance int
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			distance++
		}
	}
	return distance
}
