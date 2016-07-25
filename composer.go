package binary

import "fmt"

// Compose to compose template
func Compose(t *Tree) string {
	for _, asset := range t.Iter() {
		for _, part := range asset.Iter() {
			fmt.Println(part.Hash())
		}
	}
	return ""
}
