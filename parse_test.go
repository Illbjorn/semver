package semver

import (
	"testing"

	"gotest.tools/assert"
)

func TestParse(t *testing.T) {
	t.Run("good cases", func(t *testing.T) {
		// Prefix, Major, Minor, Patch, More
		assertVersion(t, "v1.22.333-rc0", Version{
			Prefix: "v",
			Major:  1,
			Minor:  22,
			Patch:  333,
			More:   "-rc0",
		}, "")
		// Prefix, Major, Minor, More
		assertVersion(t, "v1.22-rc0", Version{
			Prefix: "v",
			Major:  1,
			Minor:  22,
			More:   "-rc0",
		}, "")
		// Prefix, Major, More
		assertVersion(t, "v1-rc0", Version{
			Prefix: "v",
			Major:  1,
			More:   "-rc0",
		}, "")
		// Major, Minor, Patch, More
		assertVersion(t, "1.22.333-rc0", Version{
			Major: 1,
			Minor: 22,
			Patch: 333,
			More:  "-rc0",
		}, "")
		// Major, Minor, More
		assertVersion(t, "1.22-rc0", Version{
			Major: 1,
			Minor: 22,
			More:  "-rc0",
		}, "")
		// Major, Minor
		assertVersion(t, "1.22", Version{
			Major: 1,
			Minor: 22,
		}, "")
		// Major (with leading zeroes), Minor
		assertVersion(t, "001.22", Version{
			Major: 1,
			Minor: 22,
		}, "")
	})

	t.Run("bad cases", func(t *testing.T) {
		// No hyphen following patch
		assertVersion(t, "v1.22.333rc0", Version{
			Prefix: "v",
			Major:  1,
			Minor:  22,
			Patch:  333,
		}, "[1:9]: received ill-formed version string (v1.22.333rc0)")
		// Major overflows uint16
		assertVersion(t, "v65536.22.333-rc0", Version{
			Prefix: "v",
			Major:  1,
		}, "input overflows 16-bit unsigned integer")
		// Minor overflows uint16
		assertVersion(t, "v1.65536.333-rc0", Version{
			Prefix: "v",
			Major:  1,
			Minor:  22,
		}, "input overflows 16-bit unsigned integer")
		// Patch overflows uint16
		assertVersion(t, "v1.22.65536-rc0", Version{
			Prefix: "v",
			Major:  1,
			Minor:  22,
			Patch:  333,
		}, "input overflows 16-bit unsigned integer")
		// Non-"v" prefix
		assertVersion(
			t,
			"d1.22.333rc0",
			Version{},
			"[1:9]: received ill-formed version string (d1.22.333rc0)",
		)
	})
}

func assertVersion(t *testing.T, input string, expect Version, expectErr string) {
	t.Helper()

	got, err := Parse(input)
	if expectErr != "" {
		assert.Check(t, err != nil)
		assert.Equal(t, e{expectErr}, err)
	} else {
		assert.NilError(t, err)
		assert.Equal(t, expect, got)
	}
}

func BenchmarkParse(b *testing.B) {
	input := "v1.22.333-rc_alpha1a1a1a"
	for b.Loop() {
		version, err := Parse(input)
		_ = version
		_ = err
	}
}
