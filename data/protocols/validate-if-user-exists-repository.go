package protocols

type ValidateIfUserExistsRepository interface {
	HasUser(email string) (bool, error)
}
