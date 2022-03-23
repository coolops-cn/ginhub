package user

import (
	"github.com/coolops-cn/ginhub/internal/app/models"
	"github.com/coolops-cn/ginhub/pkg/database"
	"github.com/coolops-cn/ginhub/pkg/hash"
	"gorm.io/gorm"
)

// UserModel 用户模型
type User struct {
	models.BaseModel

	Username string `json:"username,omitempty"`
	Password string `json:"-"`

	models.CommonTimeStampFiled
}

func (u *User) Create() {
	database.DB.Create(&u)
}

// BeforeSave GORM 的模型钩子，在创建和更新模型前调用
func (u *User) BeforeSave(tx *gorm.DB) (err error) {

	if !hash.BcryptIsHashed(u.Password) {
		u.Password = hash.BcryptHash(u.Password)
	}
	return
}

// ComparePassword 密码是否正确
func (u *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, u.Password)
}
