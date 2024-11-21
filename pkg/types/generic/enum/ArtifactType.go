package artifact

type Type uint8

const (
	Service Type = iota
	External
	Binary
)
