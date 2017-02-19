package stathat

// StatHat is a StatHat which does StatHat things
type StatHat string

// New returns a StatHat using the provided access token (see https://www.stathat.com/access)
func New(token string) StatHat {
	return StatHat(token)
}

func (s StatHat) String() string {
	return string(s)
}

func (s StatHat) urlPrefix() string {
	return `https://www.stathat.com/x/` + s.String()
}
