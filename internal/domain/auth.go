package domain

type AuthContext struct {
	ID   int64
	Type int16
	Name string
}

type SignUpRequest struct {
	Name     string
	Phone    string
	Password string
	Type     int16

	Surname *string
	Email   *string
}

func (r SignUpRequest) Validate() error {
	return nil
}

type SignUpResponse struct {
	Token string
}

type SignInRequest struct {
	Phone    string
	Password string
}

func (r SignInRequest) Validate() error {
	return nil
}

type SignInResponse struct {
	ID         int64
	Token      string
	Name       string
	Type       int16
	Surname    *string
	Patronymic *string
}
