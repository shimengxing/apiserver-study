package model

import (
	"apiserver-study/pkg/auth"
	"apiserver-study/pkg/constvar"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type UserModel struct {
	BaseModel
	Username string `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
}

func (u *UserModel) TableName() string {
	return "tb_users"
}

func (u *UserModel) Create() error {
	return DB.Docker.Create(&u).Error
}

//根据id删除用户
func DeleteUser(id uint64) error {
	user := UserModel{}
	user.BaseModel.Id = id
	return DB.Docker.Delete(&user).Error
}

//更新用户信息
func (u *UserModel) Update() error {
	return DB.Docker.Save(u).Error
}

//通过用户名获取用户信息
func GetUser(username string) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Docker.Where("username=?", username).First(&u)
	return u, d.Error
}

//模糊获取所有用户
func ListUser(username string, offset interface{}, limit int) ([]*UserModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	users := make([]*UserModel, 0)
	var count uint64

	where := fmt.Sprintf("username like '%%%s%%'", username)
	//得到数量总数
	if err := DB.Docker.Model(&UserModel{}).Where(where).Count(&count).Error; err != nil {
		return users, count, err
	}

	//根据数据分组
	if err := DB.Docker.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err
	}

	return users, count, nil
}

//比较密码
func (u *UserModel) Compare(pwd string) (err error) {
	err = auth.Compare(u.Password, pwd)
	return err
}

//密码加密
func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return err
}

//验证字段合法性
func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
