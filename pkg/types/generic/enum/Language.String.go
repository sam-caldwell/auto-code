package enum

// String - return a string representation of our language
func (l *Language) String() string {
	return [...]string{"golang", "typeScript", "sql"}[*l]
}
