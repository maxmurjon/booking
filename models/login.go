package models

type Login struct {
	UserName string `json:"username"`
	Password    string `json:"password"`
}

type LoginResponse struct {
	UserData *User  `json:"user_data"`
	Token    string `json:"token"`
}
