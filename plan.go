package eventmedia

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

// ModerationMode defines how guest submissions become visible.
type ModerationMode string

const (
	// ModerationApprovedOnly allows only explicitly approved media to appear.
	ModerationApprovedOnly ModerationMode = "approved_only"
	// ModerationManual keeps uploads in a review workflow.
	ModerationManual ModerationMode = "manual"
	// ModerationNone publishes without moderation and is unsafe for live displays.
	ModerationNone ModerationMode = "none"
)

// Plan captures the minimum operational decisions for an event media workflow.
type Plan struct {
	EventName       string
	UploadURL       string
	ShortURL        string
	PrivacyNotice   string
	RemovalContact  string
	ModerationMode  ModerationMode
	LiveDisplay     bool
	RetentionDays   int
	UploadClosesAt  time.Time
	NetworkTests    []string
	FallbackPlan    string
	ResponsibleRole []string
}

// Severity identifies whether a finding blocks readiness.
type Severity string

const (
	// SeverityError blocks readiness.
	SeverityError Severity = "error"
	// SeverityWarning should be reviewed but does not block readiness.
	SeverityWarning Severity = "warning"
)

// Finding describes one incomplete or unsafe planning decision.
type Finding struct {
	Severity Severity
	Field    string
	Message  string
}

// Result is the deterministic outcome of Validate.
type Result struct {
	Score    int
	Ready    bool
	Findings []Finding
}

// Validate checks an event media plan at the supplied time.
func Validate(plan Plan, now time.Time) Result {
	findings := make([]Finding, 0)
	add := func(severity Severity, field, message string) {
		findings = append(findings, Finding{Severity: severity, Field: field, Message: message})
	}

	if strings.TrimSpace(plan.EventName) == "" {
		add(SeverityError, "event_name", "add a human-readable event name")
	}
	if !isHTTPS(plan.UploadURL) {
		add(SeverityError, "upload_url", "use a valid HTTPS guest upload URL")
	}
	if !isHTTPS(plan.ShortURL) {
		add(SeverityWarning, "short_url", "provide a tested HTTPS text fallback")
	}
	if len(strings.TrimSpace(plan.PrivacyNotice)) < 30 {
		add(SeverityError, "privacy_notice", "explain audience, purpose, and visibility")
	}
	if strings.TrimSpace(plan.RemovalContact) == "" {
		add(SeverityError, "removal_contact", "document a removal or privacy contact")
	}

	switch plan.ModerationMode {
	case ModerationApprovedOnly, ModerationManual, ModerationNone:
	default:
		add(
			SeverityError,
			"moderation_mode",
			fmt.Sprintf("choose %q, %q, or %q", ModerationApprovedOnly, ModerationManual, ModerationNone),
		)
	}
	if plan.LiveDisplay && plan.ModerationMode == ModerationNone {
		add(SeverityError, "live_display", "do not publish unreviewed guest media")
	}

	if plan.RetentionDays <= 0 {
		add(SeverityError, "retention_days", "set a positive retention period")
	} else if plan.RetentionDays > 730 {
		add(SeverityWarning, "retention_days", "record the reason for retention beyond two years")
	}

	if plan.UploadClosesAt.IsZero() {
		add(SeverityWarning, "upload_closes_at", "set and communicate an upload closing time")
	} else if !now.IsZero() && !plan.UploadClosesAt.After(now) {
		add(SeverityWarning, "upload_closes_at", "the configured upload window has already closed")
	}

	if !contains(plan.NetworkTests, "venue_wifi") {
		add(SeverityWarning, "network_tests", "test a real upload over venue guest Wi-Fi")
	}
	if !contains(plan.NetworkTests, "mobile_data") {
		add(SeverityWarning, "network_tests", "test a real upload over mobile data")
	}
	if len(strings.TrimSpace(plan.FallbackPlan)) < 25 {
		add(SeverityWarning, "fallback_plan", "document a privacy-preserving connectivity fallback")
	}
	if !contains(plan.ResponsibleRole, "moderation") {
		add(SeverityWarning, "responsible_role", "assign an explicit moderation owner")
	}
	if !contains(plan.ResponsibleRole, "export") {
		add(SeverityWarning, "responsible_role", "assign an explicit export and archive owner")
	}

	errors, warnings := 0, 0
	for _, finding := range findings {
		if finding.Severity == SeverityError {
			errors++
		} else {
			warnings++
		}
	}
	score := 100 - errors*15 - warnings*5
	if score < 0 {
		score = 0
	}

	return Result{Score: score, Ready: errors == 0, Findings: findings}
}

func isHTTPS(value string) bool {
	parsed, err := url.ParseRequestURI(value)
	return err == nil && parsed.Scheme == "https" && parsed.Host != ""
}

func contains(values []string, expected string) bool {
	for _, value := range values {
		if value == expected {
			return true
		}
	}
	return false
}
