package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyz"
	numbers  = "0123456789"
)

var (
	loremIpsum = []string{
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		"Praesent euismod, quam at lacinia bibendum, augue lacus bibendum quam.",
		"Donec euismod, nisl eget ultricies ultricies, nunc nisl ultricies nunc.",
		"Sed euismod, quam at lacinia bibendum, augue lacus bibendum quam.",
		"Etiam euismod, nisl eget ultricies ultricies, nunc nisl ultricies nunc.",
	}

	// Create a new random number generator
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// RandomInt generates a random integer between min and max
func RandomInt(min, max int) int {
	return min + rng.Intn(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rng.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomTitle generates a random blog title
func RandomTitle() string {
	adjectives := []string{"Amazing", "Incredible", "Insightful", "Revolutionary", "Game-Changing",
		"Remarkable", "Exceptional", "Phenomenal", "Outstanding", "Spectacular", "Groundbreaking",
		"Cutting-Edge", "Innovative", "Transformative", "Pioneering"}

	nouns := []string{"Discovery", "Innovation", "Technique", "Strategy", "Approach",
		"Breakthrough", "Method", "Solution", "Concept", "Idea"}
	return fmt.Sprintf("%s %s", adjectives[rng.Intn(len(adjectives))], nouns[rng.Intn(len(nouns))])
}

// RandomSlug generates a random URL-friendly slug
func RandomSlug() string {
	return strings.ToLower(strings.ReplaceAll(RandomTitle(), " ", "-"))
}

// RandomParagraph generates a random paragraph
func RandomParagraph() string {
	numSentences := RandomInt(3, 6)
	var paragraph strings.Builder
	for i := 0; i < numSentences; i++ {
		paragraph.WriteString(loremIpsum[rng.Intn(len(loremIpsum))])
		paragraph.WriteString(" ")
	}
	return strings.TrimSpace(paragraph.String())
}

// RandomContent generates random blog content
func RandomContent() string {
	numParagraphs := RandomInt(3, 5)
	var content strings.Builder
	for i := 0; i < numParagraphs; i++ {
		content.WriteString(RandomParagraph())
		content.WriteString("\n\n")
	}
	return strings.TrimSpace(content.String())
}

// RandomDescription generates a random blog description
func RandomDescription() string {
	return loremIpsum[rng.Intn(len(loremIpsum))]
}

// RandomImageURL generates a random image URL
func RandomImageURL() string {
	width := RandomInt(800, 1200)
	height := RandomInt(400, 800)
	return fmt.Sprintf("https://picsum.photos/%d/%d", width, height)
}

// RandomUUID generates a random UUID
func RandomUUID() uuid.UUID {
	return uuid.New()
}

// RandomUsername generates a random username
func RandomUsername() string {
	return RandomString(RandomInt(5, 10))
}

// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@example.com", RandomUsername())
}

// RandomDate generates a random date within the last year
func RandomDate() time.Time {
	now := time.Now()
	days := rng.Intn(365)
	return now.AddDate(0, 0, -days)
}
