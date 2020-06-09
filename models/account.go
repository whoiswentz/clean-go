package models

type AddAccountModel struct {
	Name                 string `json:"name"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

func (a AddAccountModel) IsNameEmpty() bool {
	return a.Name == ""
}

func (a AddAccountModel) IsPasswordEmpty() bool {
	return a.Password == ""
}

func (a AddAccountModel) IsEmailEmpty() bool {
	return a.Email == ""
}

func (a AddAccountModel) IsPasswordConfirmationEmpty() bool {
	return a.PasswordConfirmation == ""
}

type AccountModel struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
