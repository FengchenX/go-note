package mysqldb

import (
	"agfun/tv/entity"
	"agfun/util"
	"fmt"
)

func AddFreeTVs(free *entity.FreeTV) error {
	db := getSysDB().Create(free)
	return db.Error
}
func GetFreeTVs(free entity.FreeTV, filter *util.PageFilter) ([]*entity.FreeTV, int, error) {
	sql := ""
	var params []interface{}
	comma := ""
	if len(free.ID) > 0 {
		sql = fmt.Sprintf("%s %s uid = ?", sql, comma)
		params = append(params, free.ID)
		comma = "AND"
	}
	if len(free.Name) > 0 {
		sql = fmt.Sprintf("%s %s name = ?", sql, comma)
		params = append(params, free.Name)
		comma = "AND"
	}
	var frees []*entity.FreeTV
	var total int
	db := getSysDB().Model(&entity.FreeTV{}).Where(sql, params...).Count(&total)
	if db.Error != nil {
		return nil, -1, db.Error
	}
	sql = util.PageFilterSql(sql, "id", filter)
	db = getSysDB().Where(sql, params...).Find(&frees)
	return frees, total, db.Error
}

func UpdateFreeTV(free entity.FreeTV, querys map[string]interface{}) error {
	up := make(map[string]interface{}, 20)
	if len(free.ID) > 0 {
		up["uidl"] = free.ID
	}
	if len(free.Name) > 0 {
		up["name"] = free.Name
	}
	if len(free.Desc) > 0 {
		up["desc"] = free.Desc
	}
	if len(free.URL) > 0 {
		up["url"] = free.URL
	}

	if free.ID > 0 {
		db := getSysDB().Model(&free).Updates(up)
		return db.Error
	} else {
		sql := ""
		var params []interface{}
		comma := ""
		for k, v := range querys {
			sql = fmt.Sprintf("%s %s %s =?", sql, comma, k)
			params = append(params, v)
			comma = "AND"
		}
		db := getSysDB().Model(&entity.FreeTV{}).Updates(up).Where(sql, params...)
		return db.Error
	}
}

func DelFreeTV(free entity.FreeTV) error {
	if free.ID > 0 {
		db := getSysDB().Delete(&free)
		return db.Error
	}
	sql := ""
	var params []interface{}
	comma := ""
	if len(free.ID) > 0 {
		sql = fmt.Sprintf("%s %s uid = ?", sql, comma)
		params = append(params, free.ID)
		comma = "AND"
	}
	if len(free.Name) > 0 {
		sql = fmt.Sprintf("%s %s name = ?", sql, comma)
		params = append(params, free.Name)
		comma = "AND"
	}
	db := getSysDB().Delete(&entity.FreeTV{}).Where(sql, params...)
	return db.Error
}
