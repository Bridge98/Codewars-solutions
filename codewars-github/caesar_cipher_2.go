package kata

import (
	"math"
	"strings"
	"unicode"
	"unicode/utf8"
)

func Chunk(s string, index int) string {
	if index < 0 || index > 4 {
		return ""
	}

	var baseSize, size int
	if len(s)%5 == 0 {
		baseSize = len(s) / 5
		size = baseSize
	} else {
		baseSize = int(math.Ceil(float64(len(s)) / 5))
		if index == 4 {
			size = len(s) % baseSize
		} else {
			size = baseSize
		}
	}
	res := s[index*baseSize : index*baseSize+size]
	return res
}

/*
Prefix builds the encoding prefix.

It assumes s is a string containing at least one rune, and its first rune is an
ASCII character, upper- or lower-case.
*/

func Prefix(s string, shift int) []rune {
	r := []rune(s)[0]
	r = unicode.ToLower(r)
	var off = r - 'a'
	if off < 0 || off >= 26 {
		panic("Invalid string")
	}
	off += int32(shift)
	off %= 26
	r2 := 'a' + off
	return []rune{r, r2}
}

func ApplyShift(r rune, shift int) rune {
	if r >= 'A' && r <= 'Z' {
		return 'A' + (r-'A'+rune(shift))%26
	}
	if r >= 'a' && r <= 'z' {
		return 'a' + (r-'a'+rune(shift))%26
	}
	return r
}

func Shift(s string) int {
	if utf8.RuneCountInString(s) < 2 {
		panic("String too short to decode")
	}
	rs := []rune(s)
	var initial, shifted int = int(rs[0]), int(rs[1])
	shift := (26 + shifted - initial)%26 // Avoid negatives
	return shift
}

func Encode(s string, shift int) []string {
	const prefixLen = 2
	src := []rune(s)
	dst := make([]rune, len(src)+prefixLen)
	prefix := Prefix(s, shift)
	n := copy(dst, prefix)
	if n != prefixLen {
		panic("copy error on prefix")
	}

	for i, r := range src {
		dst[i+prefixLen] = ApplyShift(r, shift)
	}
	s = string(dst)
	res := []string{Chunk(s, 0), Chunk(s, 1), Chunk(s, 2), Chunk(s, 3), Chunk(s, 4)}
	if len(res[4]) == 0 {
		res = res[:4]
	}
	return res
}

func Decode(arr []string) string {
	reverseShift := 26 - Shift(arr[0])
	s := strings.Join(arr, "")[2:]
	l := utf8.RuneCountInString(s)
	runes := make([]rune, l)
	for i, r := range s {
		runes[i] = ApplyShift(r, reverseShift)
	}
	s = string(runes)
	return s
}