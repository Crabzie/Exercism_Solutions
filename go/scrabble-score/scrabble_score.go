// package scrabble compute the scrabble score for a word.
package scrabble

// Score function takes a word and returns the scrabbled result.
// count variable holds the sum of numbers.
// temp variable holds the string value of the byte's.
func Score(word string) int {
	var count int = 0
	letters := map[string]int{"a": 1, "e": 1, "i": 1, "o": 1, "u": 1, "l": 1, "n": 1, "r": 1, "s": 1, "t": 1, "d": 2, "g": 2, "b": 3, "c": 3, "m": 3, "p": 3, "f": 4, "h": 4, "v": 4, "w": 4, "y": 4, "k": 5, "j": 8, "x": 8, "q": 10, "z": 10}
	for i := 0; i < len(word); i++ {
		temp := string(word[i])
		count += letters[temp]
	}
	return count
}
