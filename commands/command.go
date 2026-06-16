package commands

// Command processes lines of code
type Command interface {
	// Execute retrieves a line of code, processes it and returns the resulting line of code
	Execute(loc string) string
}
