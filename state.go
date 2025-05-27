package semver

type State uint8

const (
	StatePrefix State = 1 + iota
	StateMajor
	StateMinor
	StatePatch
	StateMore
)

func (self State) String() string {
	switch self {
	case StatePrefix:
		return "prefix"
	case StateMajor:
		return "major"
	case StateMinor:
		return "minor"
	case StatePatch:
		return "patch"
	case StateMore:
		return "more"
	default:
		return "none"
	}
}
