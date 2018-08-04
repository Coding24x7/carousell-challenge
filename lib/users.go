package lib

// User represents user account
type User struct {
	Name     string `json:"name"` // label of the healthRule
	password string
	Country  string `json:"country"` // healthcheck rule to run
}

// RegisterUser creates new user account
func RegisterUser(name, password, country string) (*User, error) {
	u := &User{
		Name:     name,
		password: password,
		Country:  country,
	}
	err := UserRegistry.Set(u)
	return u, err
}

// LoginUser logs in user account
func LoginUser(name, password string) (*User, error) {
	return UserRegistry.Login(name, password)
}

// DeleteUser deletes user account
func DeleteUser(name string) (*User, error) {
	return UserRegistry.Delete(name)
}
