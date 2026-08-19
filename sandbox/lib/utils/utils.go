// Package utils holds the small computations the rest of the sandbox is
// written with: rendering an amount, doing arithmetic on a calendar, and
// packing several strings into one storage key.
//
// Nothing here knows what a registry is. Every function takes plain values
// and hands plain values back, which is what lets a task, a visualization and
// the ledger all reach for the same one and always render a figure the same
// way.
//
// This package declares no types and no factories, so no specification
// governs it.
package utils

import "strings"

// Separator splits the parts of a packed key. It is a character no name, date
// or amount can hold, so unpacking is never ambiguous — a description
// carrying one is rejected when the task validates its fields.
const Separator = "|"

// Pack composes one key out of several parts, so free text that is not unique
// can be stored beside a part that is. The first part is what makes the whole
// key unique, and the last part is the only one allowed to contain the
// separator — Unpack hands it back whole.
func Pack(parts ...string) string {
	return strings.Join(parts, Separator)
}

// Unpack splits a key composed by Pack back into count parts, padding with
// empty strings when the stored key carries fewer. The last part keeps
// everything that is left, separators included.
func Unpack(key string, count int) []string {
	parts := strings.SplitN(key, Separator, count)
	for len(parts) < count {
		parts = append(parts, "")
	}
	return parts
}

// Part reads one part of a packed key, returning "" when the position asked
// for is outside it.
func Part(key string, count int, index int) string {
	if index < 0 || index >= count {
		return ""
	}
	return Unpack(key, count)[index]
}
