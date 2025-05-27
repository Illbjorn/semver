package semver

func Parse(input string) (version Version, err error) {
	state := StatePrefix

	left := input
	for len(left) > 0 {
		next := left[0]

		switch {
		case isAlpha(next) && state == StatePrefix: // Prefix
			version.Prefix, left = word(left)
			state++

		case isNum(next) && state <= StateMajor: // Major
			if version.Major, left, err = number(left); err != nil {
				return
			}
			// We can land here from either Prefix/Major states, move directly to
			// StateMajor
			state = StateMinor

		case isNum(next) && state == StateMinor: // Minor
			if version.Minor, left, err = number(left); err != nil {
				return
			}
			state++

		case isNum(next) && state == StatePatch: // Patch
			if version.Patch, left, err = number(left); err != nil {
				return
			}
			state++

		case next == '-' && state > StatePrefix: // More (Consume the rest)
			version.More = left
			return

		case next == '.': // Discard
			left = left[1:]

		default:
			err = newErrIllFormed(input, len(input)-len(left))
			return
		}
	}

	return
}

/*------------------------------------------------------------------------------
 * Accumulators
 *----------------------------------------------------------------------------*/

func word(s string) (word string, left string) {
	i := 0
	for i = range len(s) {
		if !isAlpha(s[i]) {
			break
		}
	}

	if i < 0 || i >= len(s) {
		// `i` is out of bounds
		return "", ""
	} else if len(s[i:]) == 1 {
		// The word reaches the end of the input string
		return s, ""
	} else {
		// There's still more string ahead, reslice and move along
		return s[:i], s[i:]
	}
}

func number(s string) (n uint16, left string, err error) {
	i := 0
	for i = range len(s) {
		if !isNum(s[i]) {
			break
		}
	}

	if i < 0 || i >= len(s) {
		// `i` is out of bounds
		return 0, "", nil
	} else if len(s[i:]) == 1 {
		// The number reaches the end of the input string
		n, err = cint(s)
		left = ""
	} else {
		// There's still more string ahead, reslice and move along
		n, err = cint(s[:i])
		left = s[i:]
	}
	return
}
