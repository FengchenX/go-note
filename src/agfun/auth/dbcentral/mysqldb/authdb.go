package mysqldb

import (
	auth "agfun/agfun-service/dbcentral/mysqldb"
	"agfun/auth/entity"
	"github.com/jinzhu/gorm"
	"log"
	"strings"
)

func authDB() *gorm.DB {
	return auth.GetAuthDB()
}

func CreateTable() {
	if db := authDB().AutoMigrate(&entity.User{}); db.Error != nil {
		log.Fatal(db.Error)
	}
}

func CreateUser(user entity.User) (*entity.User, error) {
	db := authDB().Create(&user)
	if db.Error != nil {
		return nil, db.Error
	}
	return &user, nil
}

func GetUsers(user entity.User) ([]*entity.User, error) {
	args := ""
	var params []interface{}
	if user.UserName != "" {
		args += " AND user_name = ?"
		params = append(params, user.UserName)
	}

	args = strings.TrimPrefix(args, " AND")
	var users []*entity.User
	db := authDB().Where(args, params...).Find(&users)
	if db.Error != nil {
		return nil, db.Error
	}
	return users, nil
}
