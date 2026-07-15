package signage

import "testing"

func validSpec() Spec {
	return Spec{
		QRWidthMM:       35,
		ScanDistanceMM:  300,
		QuietZoneModule: 4,
		ContrastRatio:   8,
		ShortURL:        "https://example.com/e",
		Instruction:     "Scan to add photos to the private event album.",
		TestedDevices:   2,
	}
}

func TestValidateCompleteSpec(t *testing.T) {
	if findings := Validate(validSpec()); len(findings) != 0 {
		t.Fatalf("expected no findings, got %#v", findings)
	}
}

func TestValidateRejectsSmallCode(t *testing.T) {
	spec := validSpec()
	spec.QRWidthMM = 20
	spec.ScanDistanceMM = 500
	if findings := Validate(spec); len(findings) != 1 || findings[0].Severity != Error {
		t.Fatalf("expected one blocking finding, got %#v", findings)
	}
}

func TestValidateWarnsForOneDevice(t *testing.T) {
	spec := validSpec()
	spec.TestedDevices = 1
	if findings := Validate(spec); len(findings) != 1 || findings[0].Severity != Warning {
		t.Fatalf("expected one warning, got %#v", findings)
	}
}
