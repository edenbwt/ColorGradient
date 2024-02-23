package ColorGradient

import (
	"fmt"
	"math"
)

func ColorGradient(text string, gradient []int) string {

	colorCodes := gradient
	fragmentSize := int(math.Ceil(float64(len(text)) / 6))
	coloredResult := ""

	for i := 0; i < 6; i++ {
		startIdx := i * fragmentSize
		endIdx := (i + 1) * fragmentSize
		if endIdx > len(text) {
			endIdx = len(text)
		}

		fragment := text[startIdx:endIdx]
		coloredFragment := fmt.Sprintf("\x1b[38;5;%dm%s\x1b[0m", colorCodes[i], fragment)
		coloredResult += coloredFragment
	}

	return coloredResult
}
func ColorBACKGradient(text string, gradient []int) string {

	colorCodes := gradient
	fragmentSize := int(math.Ceil(float64(len(text)) / 6))
	coloredResult := ""

	for i := 0; i < 6; i++ {
		startIdx := i * fragmentSize
		endIdx := (i + 1) * fragmentSize
		if endIdx > len(text) {
			endIdx = len(text)
		}

		fragment := text[startIdx:endIdx]
		coloredFragment := fmt.Sprintf("\x1b[48;5;%dm%s\x1b[0m", colorCodes[i], fragment)
		coloredResult += coloredFragment
	}
	return coloredResult
}
