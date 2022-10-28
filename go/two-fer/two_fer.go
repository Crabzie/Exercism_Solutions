// package twofer provide the utility to return
// a	string
package twofer

// ShareWith function is used to return a string
// depending on the provided name
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}
	return "One for " + name + ", one for me."
}
