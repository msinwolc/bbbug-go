package models

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/msinwolc/util"
)

type User struct {
	UserId        int `gorm:"primarykey"`
	UserIcon      int
	UserSex       int
	UserAccount   string
	UserPassword  string
	UserSalt      string
	UserName      string
	UserHead      string
	UserRemark    string
	UserGroup     int
	UserIpreg     string
	UserOpenid    string
	UserExtra     string
	UserApp       int
	UserDevice    string
	UserTouchtip  string
	UserVip       string
	UserStatus    int
	UserCreatedAt time.Time `gorm:"type:datetime; not null" json:"user_created_at"`
	UserUpdatedAt time.Time `gorm:"type:datetime; not null" json:"user_updated_at"`
}

func TableName() string {
	return "sa_user"
}

func Login(account, password string) (User, error) {
	user := User{}
	result := GetDB().Table(TableName()).Where("user_account = ?", account).Find(&user)
	if result.Error != nil {
		err := fmt.Errorf("该账号不存在！")
		return user, err
	}
	if user.UserStatus == 1 {
		err := fmt.Errorf("账号已被锁定！")
		return user, err
	}
	if util.EncodeBySalt(password, user.UserSalt) != user.UserPassword {
		err := fmt.Errorf("密码错误！请重新输入")
		return user, err
	}
	return user, nil
}

func RegByLogin(account, password, ip, device string) (User, error) {
	user := User{}
	salt := Salt()
	result := GetDB().Table(TableName()).Where("user_account = ?", account).Attrs(User{
		UserAccount:   account,
		UserSalt:      salt,
		UserPassword:  util.EncodeBySalt(password, salt),
		UserName:      account,
		UserHead:      "",
		UserGroup:     0,
		UserRemark:    "",
		UserIpreg:     ip,
		UserDevice:    device,
		UserCreatedAt: time.Now(),
		UserUpdatedAt: time.Now(),
	}).FirstOrCreate(&user)
	if result.Error != nil {
		return user, result.Error
	}
	if user.UserStatus == 1 {
		return user, fmt.Errorf("账号已被锁定！")
	}
	return user, nil
}

func GetUserStatusById(id int) (int, error) {
	user := User{}
	result := GetDB().Table(TableName()).Where("user_id = ?", id).Select("user_status").First(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	return user.UserStatus, nil
}

func GetUserInfoById(id int) (User, error) {
	user := User{}
	result := GetDB().Table(TableName()).Where("uid = ?", id).First(&user)
	return user, result.Error
}

func UpdateUser(id int, attr map[string]interface{}) error {
	result := GetDB().Model(&User{}).Where("uid = ?", id).Updates(attr)
	return result.Error
}

func Salt() string {
	return fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
}
