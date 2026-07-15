// Package signage validates physical QR sign specifications for event media
// collection. A field implementation guide is available at
// https://gathmo.com/weddings/wedding-qr-code-for-photos-setup-guide.
package signage

import "net/url"

// Severity identifies whether a finding blocks field readiness.
type Severity string

const (
	// Error blocks field readiness.
	Error Severity = "error"
	// Warning should be reviewed during the physical venue test.
	Warning Severity = "warning"
)

// Finding describes one incomplete or unsafe sign decision.
type Finding struct {
	Severity Severity
	Field    string
	Message  string
}

// Spec contains measurable properties of one installed QR sign.
type Spec struct {
	QRWidthMM       float64
	ScanDistanceMM  float64
	QuietZoneModule int
	ContrastRatio   float64
	ShortURL        string
	Instruction     string
	TestedDevices   int
}

// Validate returns all deterministic findings for a sign specification.
func Validate(spec Spec) []Finding {
	findings := make([]Finding, 0)
	add := func(severity Severity, field, message string) {
		findings = append(findings, Finding{Severity: severity, Field: field, Message: message})
	}

	if spec.QRWidthMM <= 0 {
		add(Error, "qr_width_mm", "set the printed QR width")
	}
	if spec.ScanDistanceMM <= 0 {
		add(Error, "scan_distance_mm", "set the expected scanning distance")
	}
	if spec.QRWidthMM > 0 && spec.ScanDistanceMM > 0 && spec.QRWidthMM*10 < spec.ScanDistanceMM {
		add(Error, "qr_width_mm", "QR width is below one tenth of the scan distance")
	}
	if spec.QuietZoneModule < 4 {
		add(Error, "quiet_zone_modules", "keep at least four quiet-zone modules")
	}
	if spec.ContrastRatio < 4.5 {
		add(Error, "contrast_ratio", "use a contrast ratio of at least 4.5 to 1")
	}
	if !isHTTPS(spec.ShortURL) {
		add(Error, "short_url", "print a valid HTTPS text fallback")
	}
	if len(spec.Instruction) < 20 {
		add(Error, "instruction", "explain the scan action in plain language")
	}
	if spec.TestedDevices < 2 {
		add(Warning, "tested_devices", "test at least two camera and browser paths")
	}
	return findings
}

func isHTTPS(value string) bool {
	parsed, err := url.ParseRequestURI(value)
	return err == nil && parsed.Scheme == "https" && parsed.Host != ""
}
