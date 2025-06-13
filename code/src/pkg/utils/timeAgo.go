package utils

import (
	"fmt"
	"time"
)

// TimeAgo returns a human-readable string representing the time elapsed since the given time.
func TimeAgo(t time.Time) string {
	now := time.Now()
	diff := now.Sub(t)

	seconds := int(diff.Seconds())
	minutes := int(diff.Minutes())
	hours := int(diff.Hours())
	days := int(diff.Hours() / 24)

	switch {
	case seconds < 60:
		return fmt.Sprintf("%d seconds ago", seconds)
	case minutes < 60:
		return fmt.Sprintf("%d minutes ago", minutes)
	case hours < 24:
		return fmt.Sprintf("%d hours ago", hours)
	case days == 1:
		return "hier"
	case days < 30:
		return fmt.Sprintf("%d days ago", days)
	default:
		return t.Format("02/01/2006") // fallback to date
	}
}
