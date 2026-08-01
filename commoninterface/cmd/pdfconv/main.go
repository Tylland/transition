package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/ledongthuc/pdf"
)

func main() {
	inputDir := flag.String("in", ".", "directory containing PDF files")
	outputDir := flag.String("out", ".", "directory to write Markdown files")
	flag.Parse()

	if err := os.MkdirAll(*outputDir, 0755); err != nil {
		log.Fatalf("Error creating output directory: %v", err)
	}

	entries, err := os.ReadDir(*inputDir)
	if err != nil {
		log.Fatalf("Error reading input directory: %v", err)
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.EqualFold(filepath.Ext(entry.Name()), ".pdf") {
			continue
		}

		pdfPath := filepath.Join(*inputDir, entry.Name())

		mdContent, err := convertPDF(pdfPath)
		if err != nil {
			log.Printf("Error converting pdf %s: %v", pdfPath, err)
			continue
		}

		baseName := strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name()))
		outputPath := filepath.Join(*outputDir, baseName+".md")
		if err := os.WriteFile(outputPath, []byte(mdContent), 0644); err != nil {
			log.Printf("Error writing %s: %v", outputPath, err)
			continue
		}

		fmt.Printf("Converted %s -> %s\n", pdfPath, outputPath)
	}
}

// convertPDF extracts text from a PDF using position-aware spacing.
// It uses ledongthuc/pdf directly and applies an adaptive per-row threshold
// to decide whether the gap between consecutive glyph runs represents a real
// word space (insert " ") or just character-run positioning (no space).
func convertPDF(path string) (string, error) {
	f, r, err := pdf.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var sb strings.Builder

	for pageNum := 1; pageNum <= r.NumPage(); pageNum++ {
		page := r.Page(pageNum)
		if page.V.IsNull() {
			continue
		}

		rows, err := page.GetTextByRow()
		if err != nil {
			continue
		}

		// PDF Y-coordinates increase bottom-to-top; reverse to reading order.
		for i, j := 0, len(rows)-1; i < j; i, j = i+1, j-1 {
			rows[i], rows[j] = rows[j], rows[i]
		}

		var prevPos int64 = -1
		for _, row := range rows {
			if len(row.Content) == 0 {
				continue
			}
			// Large vertical gap → blank line between paragraphs.
			if prevPos >= 0 {
				yGap := math.Abs(float64(row.Position - prevPos))
				if yGap > 20 {
					sb.WriteString("\n")
				}
			}
			prevPos = row.Position

			line := buildRowText(row.Content)
			if line != "" {
				sb.WriteString(line)
				sb.WriteString("\n")
			}
		}

		sb.WriteString("\n")
	}

	return sb.String(), nil
}

// buildRowText joins glyph-run Text elements within a single row, inserting a
// space only when the normalised gap between consecutive runs exceeds the
// adaptive threshold found from the gap distribution of that row.
func buildRowText(texts []pdf.Text) string {
	if len(texts) == 0 {
		return ""
	}
	if len(texts) == 1 {
		return strings.TrimSpace(texts[0].S)
	}

	// Normalised gap[i] = (X[i] - X[i-1]) - estimatedWidth(texts[i-1])
	// A roughly constant character-width factor of 0.5× font-size is used.
	const charWidthFactor = 0.5

	normGaps := make([]float64, len(texts)-1)
	for i := 1; i < len(texts); i++ {
		prev := texts[i-1]
		rawGap := texts[i].X - prev.X
		prevWidth := float64(len([]rune(prev.S))) * prev.FontSize * charWidthFactor
		normGaps[i-1] = rawGap - prevWidth
	}

	threshold := spaceThreshold(normGaps)

	var sb strings.Builder
	sb.WriteString(texts[0].S)
	for i := 1; i < len(texts); i++ {
		if normGaps[i-1] > threshold {
			sb.WriteString(" ")
		}
		sb.WriteString(texts[i].S)
	}
	return strings.TrimSpace(sb.String())
}

// spaceThreshold finds the mid-point of the largest jump in the sorted gap
// distribution.  Small (within-word) gaps cluster below the jump; larger
// (between-word) gaps cluster above it.
func spaceThreshold(gaps []float64) float64 {
	if len(gaps) == 0 {
		return 2.0
	}

	sorted := make([]float64, len(gaps))
	copy(sorted, gaps)
	sort.Float64s(sorted)

	maxJump := 0.0
	// Default: split at the midpoint between median and max.
	threshold := (sorted[len(sorted)/2] + sorted[len(sorted)-1]) / 2

	for i := 1; i < len(sorted); i++ {
		if jump := sorted[i] - sorted[i-1]; jump > maxJump {
			maxJump = jump
			threshold = (sorted[i] + sorted[i-1]) / 2
		}
	}
	return threshold
}
