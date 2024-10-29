package arguments

// NewConfiguration - Return a new and initialized Configuration object by reference.
func NewConfiguration() *Configuration {
	return &Configuration{
		data: make(map[string]any),
	}
}
