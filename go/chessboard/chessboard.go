package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
// count variable holds how many squares are occupied.
// ok variable holds the bool value of existance of a file.
// x variable holds the bool value of the slice index
// within the file map key.
func CountInFile(cb Chessboard, file string) int {
	var count int
	_, ok := cb[file]
	if !ok {
		return 0
	} else {
		for _, x := range cb[file] {
			if x {
				count++
			}
		}
	}
	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
// count variable holds the sum of occupied squares.
// x variable holds the slice values depending on the map key.
func CountInRank(cb Chessboard, rank int) int {
	count := 0
	if rank >= 1 && rank <= 8 {
		for _, x := range cb {
			if x[rank-1] {
				count++
			}
		}
	} else {
		return 0
	}
	return count
}

// CountAll should count how many squares are present in the chessboard.
// count variable holds the sum of squares that are present.
// x variable holds the slice values depending on the map key.
func CountAll(cb Chessboard) int {
	var count int
	for _, x := range cb {
		count += len(x)
	}
	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
// count holds the sum of occupied squares.
// x holds the map value (slice).
// z holds the slice value (bool).
func CountOccupied(cb Chessboard) int {
	var count int
	for _, x := range cb {
		for _, z := range x {
			if z {
				count++
			}
		}
	}
	return count
}
