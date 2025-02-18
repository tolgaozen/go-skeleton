package internal

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

const (
	// Version is the last release (e.g. v0.1.0)
	Version = "v0.1.0"
)

// Function to create a single line of the ASCII art with centered content and color
func createLine(content string, totalWidth int, borderColor, contentColor *color.Color) string {
	contentLength := len(content)
	paddingWidth := (totalWidth - contentLength - 4) / 2
	if paddingWidth < 0 {
		paddingWidth = 0
	}
	leftPadding := strings.Repeat(" ", paddingWidth)
	rightPadding := strings.Repeat(" ", totalWidth-2-contentLength-paddingWidth)
	border := borderColor.Sprint("│")
	contentWithColor := contentColor.Sprintf("%s%s%s", leftPadding, content, rightPadding)
	return fmt.Sprintf("%s%s%s", border, contentWithColor, border)
}

// PrintBanner displays the project banner with version and repository info.
func PrintBanner() {
	borderColor := color.New(color.FgWhite)
	contentColor := color.New(color.FgWhite)

	versionInfo := fmt.Sprintf("GoSkeleton %s", Version)

	lines := []string{
		borderColor.Sprint("┌────────────────────────────────────────────────────────┐"),
		createLine(versionInfo, 58, borderColor, color.New(color.FgBlue)),
		createLine("A minimal Go project template", 58, borderColor, contentColor),
		createLine("", 58, borderColor, contentColor),
		createLine("GitHub: https://github.com/tolgaOzen/go-skeleton", 58, borderColor, contentColor),
		createLine("", 58, borderColor, contentColor),
		borderColor.Sprint("└────────────────────────────────────────────────────────┘"),
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}
