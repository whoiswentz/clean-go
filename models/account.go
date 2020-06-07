package models

type AccountModel struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

func (a AccountModel) IsNameEmpty() bool {
	return a.Name == ""
}

func (a AccountModel) IsPasswordEmpty() bool {
	return a.Password == ""
}
