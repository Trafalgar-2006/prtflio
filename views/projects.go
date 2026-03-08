package views

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type Project struct {
	Title       string
	Description string
	Tags        []string
	Highlight   string // e.g. "85% accuracy", "22 FPS"
}

var AllProjects = []Project{
	{
		Title:       "Webcraft Studios Platform",
		Description: "Full-stack web development agency platform for managing client projects, team collaboration, and seamless deployment workflows.",
		Tags:        []string{"React", "Node.js", "MongoDB", "AWS"},
	},
	{
		Title:       "Myntra Minis Integration",
		Description: "Social commerce integration enabling shoppable mini-apps within the Myntra ecosystem, driving user engagement and conversion.",
		Tags:        []string{"JavaScript", "React Native", "Firebase"},
	},
	{
		Title:       "Str8Fire.io Platform",
		Description: "SaaS platform with integrated payment processing, user dashboards, and real-time analytics for content creators.",
		Tags:        []string{"Next.js", "TypeScript", "PostgreSQL", "Stripe"},
	},
	{
		Title:       "Frog Call Classifier",
		Description: "ML pipeline for classifying frog species from audio recordings using MFCC feature extraction and ensemble methods.",
		Tags:        []string{"Python", "Scikit-learn", "MFCC"},
		Highlight:   "85% accuracy",
	},
	{
		Title:       "Bird Sound Species Detector",
		Description: "Deep learning model using CNN architectures to identify bird species from spectrograms of field recordings.",
		Tags:        []string{"TensorFlow", "CNNs", "Python"},
		Highlight:   "90% accuracy",
	},
	{
		Title:       "ISRO LEOS — Sim-to-Real Vision",
		Description: "Built a synthetic satellite imagery data factory and deployed YOLOv7 object detection on NVIDIA Jetson for real-time inference.",
		Tags:        []string{"YOLOv7", "Jetson", "Synthetic Data", "Python"},
		Highlight:   "22 FPS real-time",
	},
}

func RenderProjects(width, height, cursor, scroll int) string {
	goldStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#FFD700")).Bold(true)
	descStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#E0E0E0"))
	tagStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#BD93F9"))
	greenStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#50FA7B")).Bold(true)
	cyanStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#00DFDF"))
	dimStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#888888"))
	selectedBorder := lipgloss.NewStyle().Foreground(lipgloss.Color("#00DFDF"))
	unselectedBorder := lipgloss.NewStyle().Foreground(lipgloss.Color("#444444"))

	var b strings.Builder

	b.WriteString("\n")
	b.WriteString(" " + cyanStyle.Bold(true).Render("✦ Projects") + "\n")
	b.WriteString(" " + dimStyle.Render("Things I've built and shipped") + "\n\n")

	contentWidth := width - 8
	if contentWidth < 50 {
		contentWidth = 50
	}
	if contentWidth > 80 {
		contentWidth = 80
	}

	for i, p := range AllProjects {
		isSelected := i == cursor

		// Project number
		num := fmt.Sprintf("%02d", i+1)

		// Border character
		var prefix string
		if isSelected {
			prefix = selectedBorder.Render("▎ ")
		} else {
			prefix = unselectedBorder.Render("  ")
		}

		// Title line
		titleLine := prefix + dimStyle.Render(num+". ") + goldStyle.Render(p.Title)
		if p.Highlight != "" {
			titleLine += "  " + greenStyle.Render("⚡ "+p.Highlight)
		}
		b.WriteString(" " + titleLine + "\n")

		// Description (only for selected)
		if isSelected {
			// Wrap description to fit
			desc := wrapText(p.Description, contentWidth-6)
			for _, line := range strings.Split(desc, "\n") {
				b.WriteString("     " + descStyle.Render(line) + "\n")
			}
		}

		// Tags
		var tags []string
		for _, t := range p.Tags {
			tags = append(tags, tagStyle.Render("["+t+"]"))
		}
		b.WriteString("     " + strings.Join(tags, " ") + "\n")
		b.WriteString("\n")
	}

	b.WriteString("\n")
	b.WriteString(" " + dimStyle.Render("[↑↓ to browse · esc to go back]") + "\n")

	return b.String()
}

func wrapText(text string, maxWidth int) string {
	if maxWidth <= 0 {
		return text
	}
	words := strings.Fields(text)
	var lines []string
	var currentLine strings.Builder

	for _, word := range words {
		if currentLine.Len()+len(word)+1 > maxWidth {
			lines = append(lines, currentLine.String())
			currentLine.Reset()
		}
		if currentLine.Len() > 0 {
			currentLine.WriteString(" ")
		}
		currentLine.WriteString(word)
	}
	if currentLine.Len() > 0 {
		lines = append(lines, currentLine.String())
	}
	return strings.Join(lines, "\n")
}
