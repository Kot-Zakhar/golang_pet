package user

type RegisterUserRequestViewModel struct {
	Name  string
	Login string
}

type RegisterUserResponseViewModel struct {
	Error string
	Id    uint64
}
