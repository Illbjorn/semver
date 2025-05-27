package semver

import "fmt"

var (
	ErrIllFormedNumber error = e{"input is not a well-formed integer"}
	ErrOverflow        error = e{"input overflows 16-bit unsigned integer"}
)

func newErrIllFormed(input string, badStart int) e {
	return e{
		fmt.Sprintf(
			"[1:%d]: received ill-formed version string (%s)",
			badStart, input,
		),
	}
}

type e struct{ m string }

func (self e) Error() string { return self.m }
