package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError error

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodderAmount, err := weightFodder.FodderAmount()
	if cows == 0 {
		return 0, errors.New("division by zero")
	}
	switch fodderAmount >= 0 {
	case true:
		if cows < 0 {
			var nephewErr SillyNephewError = fmt.Errorf("silly nephew, there cannot be %d cows", cows)
			return 0, nephewErr
		}
		switch err {
		case nil:
			return fodderAmount / float64(cows), nil
		case ErrScaleMalfunction:
			return (fodderAmount * 2) / float64(cows), nil
		default:
			return 0, err
		}
	case false:
		if errors.Is(err, ErrScaleMalfunction) || errors.Is(err, nil) {
			return 0, errors.New("negative fodder")
		} else {
			return 0, errors.New("non-scale error")
		}
	}
	return 69, errors.New("what kind of nehpew is this")
}
