package User

import (

)

//Create contract or rule if we want to use our service
type UserService interface{
	RegisterUser(User *RegisterUser) //Using RegisterUser Method or We need implement RegisterUser for using our service
}

//provide package or tools for user operation
type userService struct{
	userRepository U
}

func (userRepository repository.UserRepository)

