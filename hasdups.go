package hasdups

import (
	"sort"
)

// partitionOn takes a list of strings assumed to be sorted and
// separates them out into partitions based on the letterth
// letter.
func partitionOn(offsets []string, letter int) [][]string {
	var (
		part  int
		parts [][]string
		r     = offsets[0][letter]
	)
	for i, text := range offsets {
		if text[letter] != r {
			parts = append(parts, offsets[part:i])
			part = i
			r = text[letter]
		}
	}
	parts = append(parts, offsets[part:])
	return parts
}

// largePartitions takes a group of string partitions and throws away
// any that have fewer than minCount members.
func largePartitions(offsets [][]string, minCount int) [][]string {
	var part int
	for i, p := range offsets {
		if len(p) < minCount {
			offsets[part], offsets[i] = offsets[i], offsets[part]
			part++
		}
	}
	return offsets[part:]
}

// hasDups offsets is an array of sorted strings that are assumed all to
// have a common prefix of length letter-1. If letter = minLength, then
// it declines to check further and returns true. Otherwise, it breaks
// offsets into partitions based on the letterth letter, discards any
// that have fewer than minCount members, and recursively calls hasDups
// on those sub-partitions.
func hasDups(offsets []string, minLength, minCount, letter int) bool {
	partitions := largePartitions(
		partitionOn(offsets, letter),
		minCount,
	)
	if len(partitions) == 0 {
		return false
	}
	letter++
	if letter >= minLength {
		return true
	}
	for _, part := range partitions {
		if hasDups(part, minLength, minCount, letter) {
			return true
		}
	}
	return false
}

// HasDups determines whether text has any substrings of minimum length
// minLength that occur minCount or more times.
func HasDups(text string, minLength, minCount int) bool {
	var l = len(text)
	if l == 0 {
		return false
	}
	var offsets = make([]string, l-minLength+1)
	for i := range offsets {
		offsets[i] = text[i:]
	}
	sort.Strings(offsets)
	return hasDups(offsets, minLength, minCount, 0)
}
