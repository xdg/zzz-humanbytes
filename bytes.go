// Copyright 2022 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

// Package humanbytes converts bytes to human-readable strings with ls -lh
// semantics. For bytes >= 1 KiB/kB, human-readable strings are in the form
// "X.Y <suffix>" where X and Y are single digits, or "Z <suffix>" where Z is
// two or three digits.
//
// All sizes are aggressively rounded up.
//
//  - 1025 rounds to 1.1 K
//  - 9.9 * 1024**2 + 1 rounds to 10 M
//
// Base-2 and Base-10 are supported, including modern Base-2 suffixes.
package humanbytes

import (
	"fmt"
	"math"
)

type format struct {
	base     float64
	logBase  float64
	suffixes []string
}

var formatLS = format{
	base:     1024,
	logBase:  math.Log(1024),
	suffixes: []string{"", " K", " M", " G", " T", " P", " E", " Z", " Y"},
}

var formatIEC = format{
	base:     1024,
	logBase:  math.Log(1024),
	suffixes: []string{"", " KiB", " MiB", " GiB", " TiB", " PiB", " EiB", " ZiB", " YiB"},
}

var formatSI = format{
	base:     1000,
	logBase:  math.Log(1000),
	suffixes: []string{"", " kB", " MB", " GB", " TB", " PB", " EB", " ZB", " YB"},
}

func SizeLS(size int) string {
	return humanSize(float64(size), formatLS)
}

func SizeIEC(size int) string {
	return humanSize(float64(size), formatIEC)
}

func SizeSI(size int) string {
	return humanSize(float64(size), formatSI)
}

func humanSize(size float64, f format) string {
	if size == 0 {
		return "0"
	}

	mag := math.Floor(math.Log(size) / f.logBase)
	size /= math.Pow(f.base, mag)

	switch {
	case mag == 0:
		// do nothing
	case size < 10:
		size = math.Ceil(size*10) / 10
	default:
		size = math.Ceil(size)
	}

	if size >= f.base {
		size /= f.base
		mag++
	}

	format := "%.1f%s"
	if mag == 0 || size >= 10 {
		format = "%.0f%s"
	}

	return fmt.Sprintf(format, size, f.suffixes[int(mag)])
}
