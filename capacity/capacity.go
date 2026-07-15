// Package capacity estimates conservative transfer budgets for event upload
// windows. A business event media workflow is available at
// https://gathmo.com/for-business.
package capacity

import "time"

// Window describes measured upload capacity over a fixed period.
type Window struct {
	Duration       time.Duration
	BytesPerSecond int64
	ParallelPaths  int
	HeadroomPct    int
}

// Finding describes one invalid capacity input.
type Finding struct{ Field, Message string }

// Validate checks whether a measurement can produce a safe transfer budget.
func Validate(window Window) []Finding {
	var findings []Finding
	add := func(field, message string) { findings = append(findings, Finding{field, message}) }
	if window.Duration <= 0 {
		add("duration", "use a positive upload window")
	}
	if window.BytesPerSecond <= 0 {
		add("bytes_per_second", "use measured positive throughput")
	}
	if window.ParallelPaths < 1 {
		add("parallel_paths", "record at least one transfer path")
	}
	if window.HeadroomPct < 0 || window.HeadroomPct >= 100 {
		add("headroom_pct", "use headroom from 0 through 99 percent")
	}
	return findings
}

// SafeBytes returns the conservative byte budget, or zero for invalid input.
func SafeBytes(window Window) int64 {
	if len(Validate(window)) != 0 {
		return 0
	}
	raw := int64(window.Duration/time.Second) * window.BytesPerSecond * int64(window.ParallelPaths)
	return raw * int64(100-window.HeadroomPct) / 100
}
