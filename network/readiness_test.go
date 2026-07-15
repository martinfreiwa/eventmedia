package network

import "testing"

func completeResults() []TestResult {
	return []TestResult{
		{Path: VenueWiFi, Location: "reception", Successful: true, FileBytes: 1_000_000, DurationSeconds: 3},
		{Path: Mobile, Location: "dance floor", Successful: true, FileBytes: 1_000_000, DurationSeconds: 4},
		{Path: Fallback, Location: "off-site", Successful: true, FileBytes: 1_000_000, DurationSeconds: 2},
	}
}

func TestEvaluateCompleteResults(t *testing.T) {
	if findings := Evaluate(completeResults()); len(findings) != 0 {
		t.Fatalf("expected no findings, got %#v", findings)
	}
}

func TestEvaluateReportsMissingMobilePath(t *testing.T) {
	results := completeResults()
	results = append(results[:1], results[2:]...)
	if findings := Evaluate(results); len(findings) != 1 || findings[0].Path != Mobile {
		t.Fatalf("expected missing mobile finding, got %#v", findings)
	}
}

func TestEvaluateRequiresTransferredBytes(t *testing.T) {
	results := completeResults()
	results[0].FileBytes = 0
	if findings := Evaluate(results); len(findings) != 1 || findings[0].Path != VenueWiFi {
		t.Fatalf("expected failed Wi-Fi proof, got %#v", findings)
	}
}
