package calculator

// private to the package
var opCount map[string]int

func init() {
	opCount = make(map[string]int)
}

// public
func OpCount() map[string]int {
	return opCount
}
