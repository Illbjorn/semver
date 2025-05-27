package semver

// cint is a base-10 string to uint16 converter.
//
// This is used in place of `strconv` as it avoids about 4 allocations during
// the major/minor/patch parse which cuts the time of the entire semver parse in
// half.
func cint(s string) (uint16, error) {
	const (
		base uint16 = 10
		max  uint64 = 1<<16 - 1
	)

	// `n` is the uint16 we'll build up and return later
	var n uint16

	for _, c := range []byte(s) {
		// NOTE: We can skip the [0-9] check, as this is handled during the
		// `number()` parse
		//
		// Get the [0-9] offset
		offset := c - '0'

		// Each byte we progress through `s`, we multiply by base (10)
		//
		// Check for overflow
		if uint64(n)*uint64(base) > max {
			return uint16(max), ErrOverflow
		}
		n *= base

		// Add the offset
		//
		// Check for overflow
		if uint64(n)+uint64(offset) > max {
			return uint16(max), ErrOverflow
		}
		n += uint16(offset)
	}

	return n, nil
}
