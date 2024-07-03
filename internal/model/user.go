package models

type User struct {
	*Model
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// AddUser FatModel ThinView
func AddUser(username string, password string) error {
	user := User{
		Model:    &Model{},
		Username: username,
		Password: password,
	}

	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func CheckUser(username string, password string) bool {
	user := User{
		Model:    &Model{},
		Username: username,
		Password: password,
	}

	result := db.Take(&user)
	if result.RecordNotFound() {
		return false
	}
	return true
}
