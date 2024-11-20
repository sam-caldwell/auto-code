package url

// FromString - store a given string as the Scheme
func (u *Scheme) FromString(s string) {

	*u = Scheme(s)

}
