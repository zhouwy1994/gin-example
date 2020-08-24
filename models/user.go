package models

type Users struct {
	ID int `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func ExistUserByName(username string) bool {
	var user Users
	db.Select("id").Where("username = ?", username).First(&user)
	if user.ID > 0 {
		return true
	}

	return false
}

func AddUser(username string, password string) error {
	return db.Create(&Users {
		Username : username,
		Password : password,
	}).Error
}

func GetUser(maps interface {}) (Users, error) {
	user := Users{}
	err := db.Where(maps).First(&user).Error
	return user, err
}

func DeleteUser(id int) error {
	return db.Where("id = ?", id).Delete(&Users{}).Error
}

func UpdateUser(id int, data interface {}) error {
	return db.Model(&Users{}).Where("id = ?", id).Updates(data).Error
}

func CheckUser(username, password string) (bool, error) {
	var user Users
	err := db.Select("id").Where(Users{Username : username,
		Password : password}).First(&user).Error
	if user.ID > 0 {
		return true, err
	}

	return false, err
}