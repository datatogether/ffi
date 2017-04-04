package ffi

import (
	"io"
	"testing"
)

func errorMismatch(err error, expect string) bool {
	return (err == nil && expect != "") || (err != nil && expect == "") || (err != nil && err.Error() != expect)
}

func TestFilenameFromUrlString(t *testing.T) {
	cases := []struct {
		in, out, err string
	}{
		{"https://qri.io/files/not/a/path/image.jpg", "image.jpg", ""},
		// {"https://qri.io/files/not/a/path/image", "", nil},
	}

	for i, c := range cases {
		got, err := FilenameFromUrlString(c.in)
		if errorMismatch(err, c.err) {
			t.Errorf("case %d error mismatch. expected: %s, got: %s", i, c.err, err)
		}
		if got != c.out {
			t.Errorf("case %d output mismatch. expected: %s, got: %s", i, c.out, got)
		}
	}
}

func TestMimeTypeFromFilename(t *testing.T) {
	cases := []struct {
		in, out, err string
	}{
		{"", "", "unrecognized MIME-Type: ''"},
	}

	for i, c := range cases {
		got, err := MimeTypeFromFilename(c.in)
		if errorMismatch(err, c.err) {
			t.Errorf("case %d error mismatch. expected: %s, got: %s", i, c.err, err)
		}
		if got != c.out {
			t.Errorf("case %d output mismatch. expected: %s, got: %s", i, c.out, got)
		}
	}
}

func TestMimeTypeExtension(t *testing.T) {
	cases := []struct {
		in, out, err string
	}{
		{"", "", "unrecognized MIME-Type: ''"},
	}

	for i, c := range cases {
		got, err := MimeTypeExtension(c.in)
		if errorMismatch(err, c.err) {
			t.Errorf("case %d error mismatch. expected: %s, got: %s", i, c.err, err)
		}
		if got != c.out {
			t.Errorf("case %d output mismatch. expected: %s, got: %s", i, c.out, got)
		}
	}
}

func TestExtensionMimeType(t *testing.T) {
	cases := []struct {
		in, out, err string
	}{
		{"", "", "unrecognized extension: ''"},
	}

	for i, c := range cases {
		got, err := ExtensionMimeType(c.in)
		if errorMismatch(err, c.err) {
			t.Errorf("case %d error mismatch. expected: %s, got: %s", i, c.err, err)
		}
		if got != c.out {
			t.Errorf("case %d output mismatch. expected: %s, got: %s", i, c.out, got)
		}
	}
}

func TestMimeType(t *testing.T) {
	cases := []struct {
		in  io.Reader
		out string
		err error
	}{
	// {"", "", nil},
	}

	for i, c := range cases {
		got, err := MimeType(c.in)
		if err != c.err {
			t.Errorf("case %d error mismatch. expected: %s, got: %s", i, c.err, err)
		}
		if got != c.out {
			t.Errorf("case %d output mismatch. expected: %s, got: %s", i, c.out, got)
		}
	}
}
