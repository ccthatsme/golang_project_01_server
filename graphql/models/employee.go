package models

type Employee struct {
	Id          string `json:"id"`
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
}

func NewEmployee(emp *Employee) *Employee {

	return &Employee{
		Id:          emp.Id,
		DisplayName: emp.DisplayName,
		Email:       emp.Email,
	}
}

// func NewUserByName(name string) (*User, error) {
// 	u, err := user.Lookup(name)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return NewUser(u)
// }
//
// func NewUser(u *user.User) (*User, error) {
// 	gids, err := u.GroupIds()
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return &User{
// 		UID:    u.Uid,
// 		GID:    u.Gid,
// 		Name:   u.Name,
// 		Home:   u.HomeDir,
// 		Groups: getGroups(gids),
// 	}, nil
// }
