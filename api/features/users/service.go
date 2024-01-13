package User

//Create contract or rule if we want to use our service
type UserService interface {
	RegisterUser(User *RegisterUser) error //Using RegisterUser Method or We need implement RegisterUser for using our service
}

//provide package or tools for user operation
type UserServiceImpl struct {
	userRepository UserRepository
}

//Create New Instance and return instance from userService
func NewUserService(userRepository UserRepository) UserService {
	return &UserServiceImpl{userRepository: userRepository}
}

func (s *UserServiceImpl) RegisterUser(user *RegisterUser) error {
	err := s.RegisterUser(user)

	if err != nil {
		return err
	}
	return nil
}
