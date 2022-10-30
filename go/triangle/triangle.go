package triangle

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind string

const (
	// Pick values for the following identifiers used by the test program.
	NaT Kind = "Nat" // not a triangle
	Equ Kind = "Equ" // equilateral
	Iso Kind = "Iso" // isosceles
	Sca Kind = "Sca" // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	switch v := a + b + c; {
	case v == 0 || (a+b < c || a+c < b || b+c < a):
		return NaT
	case a == b && b == c:
		return Equ
	case a != b && b != c && c != a:
		return Sca
	default:
		return Iso
	}
}
