package service

import "testing"

func TestParsePackageExt(t *testing.T) {
	filename := "Plinko_v1.0.58_cv108_2305121850_GOOGLE_6_810_release-cn.apk"
	ex := ParsePackageExt(filename, "Test")
	if ex.env != "release" {
		t.Fatalf("invalid env: %s", ex.env)
	}
	if ex.platformID != 6 {
		t.Fatalf("invalid platformID: %d", ex.platformID)
	}
	if ex.projectID != 810 {
		t.Fatalf("invalid projectID: %d", ex.projectID)
	}
	if ex.region != "cn" {
		t.Fatalf("invalid region: %s", ex.region)
	}
	if ex.description != "Test" {
		t.Fatalf("invalid description: %s", ex.description)
	}

	filename = "Plinko_v1.0.58_cv108_2305121850_GOOGLE_6_810_release.apk"
	ex = ParsePackageExt(filename, "Test")
	if ex.env != "release" {
		t.Fatalf("invalid env: %s", ex.env)
	}
	if ex.platformID != 6 {
		t.Fatalf("invalid platformID: %d", ex.platformID)
	}
	if ex.projectID != 810 {
		t.Fatalf("invalid projectID: %d", ex.projectID)
	}
	if ex.region != "" {
		t.Fatalf("invalid region: %s", ex.region)
	}
	if ex.description != "Test" {
		t.Fatalf("invalid description: %s", ex.description)
	}

}
