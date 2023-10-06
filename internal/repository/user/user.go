package user

type UserRepository struct {
}

func NewUserRepository() (*UserRepository, error) {
	return &UserRepository{}, nil
}
