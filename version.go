package semver

type Version struct {
	Prefix              string
	Major, Minor, Patch uint16
	More                string
}
