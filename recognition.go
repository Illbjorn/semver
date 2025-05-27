package semver

/*------------------------------------------------------------------------------
 * Byte Recognition
 *----------------------------------------------------------------------------*/

func isAlpha(b byte) bool {
	return b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z'
}

func isNum(b byte) bool {
	return b >= '0' && b <= '9'
}
