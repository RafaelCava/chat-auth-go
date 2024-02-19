package protocols

type Hasher interface {
	Hash(value string) (string, error)
}
