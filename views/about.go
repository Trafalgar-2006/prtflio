package views

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func RenderAbout(width, height int) string {
	cyanStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#00DFDF"))
	dimStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#888888"))
	whiteStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#E0E0E0"))
	magentaStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#FF6AC1"))
	goldStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#FFD700"))
	greenStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#50FA7B"))
	purpleStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#BD93F9"))
	divider := dimStyle.Render("  ─────────────────────────────────────────")

	var b strings.Builder
	b.WriteString("\n")
	b.WriteString(" " + cyanStyle.Bold(true).Render("✦ About") + "\n")
	b.WriteString(divider + "\n\n")

	// Name & tagline
	b.WriteString("  " + goldStyle.Bold(true).Render("Mohith Akshay Duggirala") + "\n")
	b.WriteString("  " + magentaStyle.Italic(true).Render("Engineer · Builder · Creator") + "\n\n")

	// Education
	b.WriteString("  " + cyanStyle.Bold(true).Render("🎓 Education") + "\n")
	b.WriteString("  " + whiteStyle.Render("Electronics & Computer Engineering") + "\n")
	b.WriteString("  " + whiteStyle.Render("Manipal Institute of Technology, Bangalore") + "\n")
	b.WriteString("  " + dimStyle.Render("2023 – 2027") + "\n\n")

	// Roles
	b.WriteString("  " + cyanStyle.Bold(true).Render("💼 Roles") + "\n")
	b.WriteString("  " + greenStyle.Render("▸ ") + whiteStyle.Render("Founder, ") + magentaStyle.Render("Webcraft Studios") + "\n")
	b.WriteString("  " + greenStyle.Render("▸ ") + whiteStyle.Render("President, ") + purpleStyle.Render("MBOSC") + dimStyle.Render(" (Manipal Bengaluru Open Source Community)") + "\n\n")

	// Specializations
	b.WriteString("  " + cyanStyle.Bold(true).Render("🔬 Specializations") + "\n")
	b.WriteString("  " + greenStyle.Render("▸ ") + whiteStyle.Render("Computer Vision (Sim-to-Real Transfer)") + "\n")
	b.WriteString("  " + greenStyle.Render("▸ ") + whiteStyle.Render("Full-Stack Development (React, Next.js, Node.js)") + "\n")
	b.WriteString("  " + greenStyle.Render("▸ ") + whiteStyle.Render("AI Perception Pipelines") + "\n")
	b.WriteString("  " + greenStyle.Render("▸ ") + whiteStyle.Render("Edge Deployment (NVIDIA Jetson)") + "\n\n")

	// Interests
	b.WriteString("  " + cyanStyle.Bold(true).Render("🌟 Interests") + "\n")
	b.WriteString("  " + dimStyle.Render("Building products that bridge the gap between") + "\n")
	b.WriteString("  " + dimStyle.Render("research and real-world deployment. Passionate") + "\n")
	b.WriteString("  " + dimStyle.Render("about open source, developer tools, and making") + "\n")
	b.WriteString("  " + dimStyle.Render("AI systems that actually work in production.") + "\n\n")

	b.WriteString(divider + "\n")
	b.WriteString("\n")
	b.WriteString("  " + dimStyle.Render("[esc to go back]") + "\n")

	return b.String()
}
