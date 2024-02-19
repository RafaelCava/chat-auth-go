package protocols

type HasherCompare interface {
	Compare(plaintext string, digest string) error
}
