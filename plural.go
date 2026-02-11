// Package plrl provides functions for handling language plurals.
package plrl

// Any returns string `single` when length is 1 or -1, otherwise returns `multiple`.
func Any(length int, single, multiple string) string {
	if length == 1 || length == -1 {
		return single
	}
	return multiple
}

// Many returns string `single` when length is 1 or -1,
// otherwise returns `multiple`, or an "s" if `multiple` is empty.
func Many(length int, single, multiple string) string {
	if length == 1 || length == -1 {
		return single
	}
	if multiple != "" {
		return multiple
	}
	return "s"
}

// S returns an "s" when length is not 1 or -1.
func S(length int) string {
	if length == 1 || length == -1 {
		return ""
	}
	return "s"
}
