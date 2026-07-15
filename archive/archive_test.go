package archive

import "testing"

func validExport() Export {
	return Export{
		EventID:      "event-1",
		ArchiveName:  "event-1-originals.zip",
		Checksum:     "sha256:9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
		Photos:       10,
		Videos:       2,
		Audio:        1,
		ManifestRows: 13,
		ReadOnlyCopy: true,
	}
}

func TestValidateCompleteExport(t *testing.T) {
	if findings := Validate(validExport()); len(findings) != 0 {
		t.Fatalf("expected no findings, got %#v", findings)
	}
}

func TestValidateDetectsManifestMismatch(t *testing.T) {
	export := validExport()
	export.ManifestRows = 12
	if findings := Validate(export); len(findings) != 1 {
		t.Fatalf("expected one finding, got %#v", findings)
	}
}

func TestValidateRequiresImmutableCopy(t *testing.T) {
	export := validExport()
	export.ReadOnlyCopy = false
	if findings := Validate(export); len(findings) != 1 {
		t.Fatalf("expected one finding, got %#v", findings)
	}
}
