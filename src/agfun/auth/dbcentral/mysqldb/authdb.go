package mysqldb

import (
	auth "agfun/agfun-service/dbcentral/mysqldb"
	"agfun/auth/entity"
	"fmt"
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
	if db := authDB().AutoMigrate(&entity.VipUser{}); db.Error != nil {
		log.Fatal(db.Error)
	}
}

func CreateUser(user *entity.User) (*entity.User, error) {
	if user == nil {
		return nil, fmt.Errorf("user is nil")
	}
	db := authDB().Create(user)
	if db.Error != nil {
		return nil, db.Error
	}
	return user, nil
}

func GetUsers(user entity.User) ([]*entity.User, int, error) {
	args := ""
	var params []interface{}
	if user.UserName != "" {
		args += " AND user_name = ?"
		params = append(params, user.UserName)
	}
	if user.Pwd != "" {
		args += " AND pwd = ?"
		params = append(params, user.Pwd)
	}

	args = strings.TrimPrefix(args, " AND")
	var users []*entity.User
	db := authDB().Where(args, params...).Find(&users)
	if db.Error != nil {
		return nil, -1, db.Error
	}
	var total int
	db = authDB().Model(&entity.User{}).Where(args, params...).Count(&total)
	if db.Error != nil {
		return nil, -1, db.Error
	}
	return users, -1, nil
}

func UpdateUser(user *entity.User) error {
	db := authDB().Model(&entity.User{}).Updates(user)
	if db.Error != nil {
		return db.Error
	}
	return nil
}
func AddVip(vip *entity.VipUser) error {
	if vip == nil {
		return fmt.Errorf("vip is nil")
	}
	db := auth.GetAuthDB().Create(vip)
	return db.Error
}
func GetVips(vip entity.VipUser) ([]*entity.VipUser, error) {
	sql := ""
	var param []interface{}
	comma := ""
	if vip.ID > 0 {
		sql = fmt.Sprintf("%s %s id = ?", sql, comma)
		param = append(param, vip.ID)
		comma = "AND"
	}
	if vip.UserID > 0 {
		sql = fmt.Sprintf("%s %s user_id = ?", sql, comma)
		param = append(param, vip.UserID)
		comma = "AND"
	}
	if vip.Level > 0 {
		sql = fmt.Sprintf("%s %s level = ?", sql, comma)
		param = append(param, vip.Level)
		comma = "AND"
	}
	if vip.Expire.Unix() > 0 {
		sql = fmt.Sprintf("%s %s expire >= ?", sql, comma)
		param = append(param, vip.Expire)
		comma = "AND"
	}

	var vips []*entity.VipUser
	db := auth.GetAuthDB().Where(sql, param...).Find(&vips)
	if db.Error != nil {
		return nil, db.Error
	}
	return vips, nil
}
func UpdateVip(vip *entity.VipUser, src entity.VipUser) error {
	if vip == nil {
		return fmt.Errorf("vip is nil")
	}
	newVip := make(map[string]interface{}, 20)
	if vip.Level > 0 {
		newVip["level"] = vip.Level
	}
	if vip.Expire.Unix() > 0 {
		newVip["expire"] = vip.Expire
	}

	sql := ""
	var params []interface{}
	comma := ""
	if src.ID > 0 {
		sql = fmt.Sprintf("%s %s id = ?", sql, comma)
		params = append(params, src.ID)
		comma = "AND"
	}
	if src.UserID > 0 {
		sql = fmt.Sprintf("%s %s user_id = ?", sql, comma)
		params = append(params, src.UserID)
		comma = "AND"
	}
	if src.Expire.Unix() > 0 {
		sql = fmt.Sprintf("%s %s expire > ?", sql, comma)
		params = append(params, src.Expire)
		comma = "AND"
	}
	db := auth.GetAuthDB().Model(vip).Where(sql, params...).Updates(newVip)
	return db.Error
}
func DelVip(vip entity.VipUser) error {
	sql := ""
	var params []interface{}
	comma := ""
	if vip.ID > 0 {
		sql = fmt.Sprintf("%s %s id = ?", sql, comma)
		params = append(params, vip.ID)
		comma = "AND"
	}
	if vip.UserID > 0 {
		sql = fmt.Sprintf("%s %s user_id = ?", sql, comma)
		params = append(params, vip.UserID)
		comma = "AND"
	}
	if vip.Expire.Unix() > 0 {
		sql = fmt.Sprintf("%s %s expire < ?", sql, comma)
		params = append(params, vip.Expire)
		comma = "AND"
	}
	db := auth.GetAuthDB().Where(sql, params...).Delete(entity.VipUser{})
	return db.Error
}
