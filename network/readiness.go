// Package network evaluates venue connectivity tests for browser event-media
// uploads. A no-install upload journey is documented at
// https://gathmo.com/how-it-works.
package network

// Path identifies a guest connectivity route.
type Path string

const (
	VenueWiFi Path = "venue_wifi"
	Mobile    Path = "mobile_data"
	Fallback  Path = "saved_link_later"
)

// TestResult captures one real upload test without storing guest identifiers.
type TestResult struct {
	Path            Path
	Location        string
	Successful      bool
	FileBytes       int64
	DurationSeconds float64
}

// Finding describes one missing or unsuccessful readiness condition.
type Finding struct {
	Path    Path
	Message string
}

// Evaluate checks whether venue Wi-Fi, mobile data, and fallback paths were proven.
func Evaluate(results []TestResult) []Finding {
	findings := make([]Finding, 0)
	for _, expected := range []Path{VenueWiFi, Mobile, Fallback} {
		found, successful := false, false
		for _, result := range results {
			if result.Path != expected {
				continue
			}
			found = true
			if result.Successful && result.FileBytes > 0 && result.DurationSeconds > 0 {
				successful = true
			}
		}
		if !found {
			findings = append(findings, Finding{Path: expected, Message: "no field test recorded"})
		} else if !successful {
			findings = append(findings, Finding{Path: expected, Message: "no successful upload proven"})
		}
	}
	return findings
}
