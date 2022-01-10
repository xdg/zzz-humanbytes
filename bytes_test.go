// Copyright 2022 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License") you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package humanbytes

import (
	"strings"
	"testing"
)

type testcase struct {
	input  int
	output string
}

// All test case rounding based on `ls -lh` on Ubuntu 20.04.

var kib = float64(1024)
var base2Cases = []testcase{
	{0, "0"},
	{int(kib - 1), "1023"},
	{int(kib), "1.0 K"},
	{int(kib + 1), "1.1 K"},
	{int(1.1 * kib), "1.1 K"},
	{int(1.1*kib) + 1, "1.2 K"},
	{int(1.9 * kib), "1.9 K"},
	{int(1.9*kib) + 1, "2.0 K"},
	{int(9 * kib), "9.0 K"},
	{int(9*kib + 1), "9.1 K"},
	{int(9.9 * kib), "9.9 K"},
	{int(9.9*kib) + 1, "10 K"},
	{int(10 * kib), "10 K"},
	{int(10*kib + 1), "11 K"},
	{int((kib - 1) * kib), "1023 K"},
	{int((kib-1)*kib + 1), "1.0 M"},
	{int(kib*kib - 1), "1.0 M"},
	{int(kib * kib), "1.0 M"},
	{int(kib*kib + 1), "1.1 M"},
	{int(1.1 * kib * kib), "1.1 M"},
	{int(1.1*kib*kib) + 1, "1.2 M"},
	{int(1.9 * kib * kib), "1.9 M"},
	{int(1.9*kib*kib) + 1, "2.0 M"},
	{int(9 * kib * kib), "9.0 M"},
	{int(9*kib*kib + 1), "9.1 M"},
	{int(9.9 * kib * kib), "9.9 M"},
	{int(9.9*kib*kib) + 1, "10 M"},
	{int(10 * kib * kib), "10 M"},
	{int(10*kib*kib + 1), "11 M"},
	{int((kib - 1) * kib * kib), "1023 M"},
	{int((kib-1)*kib*kib + 1), "1.0 G"},
}

var kb = float64(1000)
var base10Cases = []testcase{
	{0, "0"},
	{int(kb - 1), "999"},
	{int(kb), "1.0 kB"},
	{int(kb + 1), "1.1 kB"},
	{int(1.1 * kb), "1.1 kB"},
	{int(1.1*kb) + 1, "1.2 kB"},
	{int(1.9 * kb), "1.9 kB"},
	{int(1.9*kb) + 1, "2.0 kB"},
	{int(9 * kb), "9.0 kB"},
	{int(9*kb + 1), "9.1 kB"},
	{int(9.9 * kb), "9.9 kB"},
	{int(9.9*kb) + 1, "10 kB"},
	{int(10 * kb), "10 kB"},
	{int(10*kb + 1), "11 kB"},
	{int((kb - 1) * kb), "999 kB"},
	{int((kb-1)*kb + 1), "1.0 MB"},
	{int(kb*kb - 1), "1.0 MB"},
	{int(kb * kb), "1.0 MB"},
	{int(kb*kb + 1), "1.1 MB"},
	{int(1.1 * kb * kb), "1.1 MB"},
	{int(1.1*kb*kb) + 1, "1.2 MB"},
	{int(1.9 * kb * kb), "1.9 MB"},
	{int(1.9*kb*kb) + 1, "2.0 MB"},
	{int(9 * kb * kb), "9.0 MB"},
	{int(9*kb*kb + 1), "9.1 MB"},
	{int(9.9 * kb * kb), "9.9 MB"},
	{int(9.9*kb*kb) + 1, "10 MB"},
	{int(10 * kb * kb), "10 MB"},
	{int(10*kb*kb + 1), "11 MB"},
	{int((kb - 1) * kb * kb), "999 MB"},
	{int((kb-1)*kb*kb + 1), "1.0 GB"},
}

func TestBase2LS(t *testing.T) {
	for _, c := range base2Cases {
		got := SizeLS(c.input)
		if got != c.output {
			t.Errorf("got '%s', but wanted '%s'", got, c.output)
		}

	}
}

func TestBase2IEC(t *testing.T) {
	for _, c := range base2Cases {
		got := SizeIEC(c.input)
		want := c.output
		if strings.Contains(want, " ") {
			want += "iB"
		}
		if got != want {
			t.Errorf("got '%s', but wanted '%s'", got, want)
		}

	}
}

func TestBase10SI(t *testing.T) {
	for _, c := range base10Cases {
		got := SizeSI(c.input)
		if got != c.output {
			t.Errorf("got '%s', but wanted '%s'", got, c.output)
		}

	}
}
