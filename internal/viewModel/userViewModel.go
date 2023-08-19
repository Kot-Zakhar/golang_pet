package viewModel

type UserViewModel struct {
	Id    uint64 `json:"id" validate: "required"`
	Name  string `json:"name" validate: "required"`
	Login string `json:"login" validate: "required"`
	Email string `json:"email" validate: "required,email"`
}

type CreateOrUpdateUserViewModel struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
