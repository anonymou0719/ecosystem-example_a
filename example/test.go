package example

import "fmt"

// ExampleAPI provides an example API function.
func ExampleAPI(name string) string {
	return fmt.Sprintf("Hello, %s! Welcome to the API.", name)
}
