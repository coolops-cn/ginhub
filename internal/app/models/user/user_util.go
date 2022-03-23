package user

import "github.com/coolops-cn/ginhub/pkg/database"

// 判断用户名是否存在
func IsUsernameExist(name string) bool {
	var count int64
	database.DB.Model(User{}).Where("username = ?", name).Count(&count)
	return count > 0
}

// GetByMulti 通过 手机号/Email/用户名 来获取用户
func GetByMulti(loginID string) (userModel User) {
	database.DB.
		Where("username = ?", loginID).First(&userModel)
	return
}

func Get(idstr string) (userModel User) {
	database.DB.Where("id", idstr).First(&userModel)
	return
}
