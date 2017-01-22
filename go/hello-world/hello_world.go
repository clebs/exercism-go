// Package greeting provides means of saying hello to people
package greeting

import "fmt"

const testVersion = 3

// HelloWorld greets the given name or the whole world if none is given
func HelloWorld(name string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("Hello, %s!", name)
}
