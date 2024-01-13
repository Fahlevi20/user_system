//structure data
package User

import{

}

type User struct{
	ID string 
	First_name string
	Last_name string
	Password string
	Email [] string
	// Phone string
	// Token
	// user_type
	// Refresh_token
	// Created_at
	// Update_at
	// User_id

}

type RegisterUser struct{
	type User struct{
		ID string 
		First_name string
		Last_name string
		Password string
		Email  string
		// Phone string
		// Token
		// user_type
		// Refresh_token
		// Created_at
		// Update_at
		// User_id
	
	}

}

type UserLogin struct{
	ID string 
	Password string
	Email  string
	// Token
	// user_type
	// Refresh_token
	// Created_at
	// Update_at
	// User_id

}
