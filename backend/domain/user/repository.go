package user

type UserRepository interface {
	GetByID(id int) (*User, error)
    GetAll() ([]User, error)
    Create(user *User) error
    Update(user *User) error
    Delete(id int) error
}
