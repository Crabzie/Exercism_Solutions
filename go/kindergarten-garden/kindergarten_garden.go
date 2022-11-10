// BenchmarkNewGarden-2               21458             46684 ns/op            9676 B/op        154 allocs/op
// BenchmarkGarden_Plants-2         1683150               741.9 ns/op             0 B/op          0 allocs/op
package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

// Garden and plant types are used for creating variable in order to solve this problem.
type plant []string
type Garden map[string]plant

// codeChars constant holds a string containing key runes for checking diagram errors.
const codeChars string = "VRCG\n"

// fixDiagramAndCheckGarden function checks diagram errors, splits diagram rows, and put each row in a variable.
// temp variable in fixDiagram function holds the split result on diagram.
func fixDiagramAndCheckGarden(children []string, diagram string) (err error, secondRow []string, firstRow []string) {
	if (len(diagram)-2)/len(children) != 4 {
		return errors.New("odd number of cups"), nil, nil
	}
	for i, char := range diagram {
		//validate first char
		if i == 0 && char != '\n' {
			return errors.New("wrong diagram format"), nil, nil
		}
		if !strings.Contains(codeChars, string(char)) {
			return errors.New("unsupported charracter:" + string(char)), nil, nil
		}

	}
	temp := strings.Split(strings.TrimPrefix(diagram, "\n"), "\n")
	return nil, strings.Split(temp[1], ""), strings.Split(temp[0], "")
}

// plants variable holds a map with plants names
// result variable is a map that holds children as keys and their flowers as values.
// keys variable holds the sorted slice of children slice.
// firstRow variables holds the first row of the diagram.
// secondRow variable holds the second row of the diagram.
func NewGarden(diagram string, children []string) (*Garden, error) {
	plants := map[string]string{"V": "violets", "C": "clover", "G": "grass", "R": "radishes"}
	result := make(Garden)
	err, secondRow, firstRow := fixDiagramAndCheckGarden(children, diagram)
	if err != nil {
		return nil, err
	}
	keys := make([]string, 0, len(children))
	keys = append(keys, children...)
	sort.Strings(keys)
	for _, c := range keys {
		_, ok := result[c]
		if !ok {
			for i, j := firstRow, 0; j < 2; firstRow, j = append(i[:0], i[1:]...), j+1 {
				result[c] = append(result[c], plants[i[0]])
			}
			for v, k := secondRow, 0; k < 2; secondRow, k = append(v[:0], v[1:]...), k+1 {
				result[c] = append(result[c], plants[v[0]])
			}
		} else {
			return nil, errors.New("duplicated names")
		}

	}
	return &result, nil
}
func (g Garden) Plants(child string) ([]string, bool) {
	_, ok := g[child]
	if !ok {
		return []string{}, false
	}
	return g[child], true
}
