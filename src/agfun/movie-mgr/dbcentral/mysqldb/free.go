package mysqldb

import (
	"agfun/agfun-service/util"
	"agfun/movie-mgr/entity"
	"fmt"
)

func AddFreeMovies(free *entity.FreeMovie) error {
	db := getSysDB().Create(free)
	return db.Error
}
func GetFreeMovies(free entity.FreeMovie, filter *util.PageFilter) ([]*entity.FreeMovie, int, error) {
	sql := ""
	var params []interface{}
	comma := ""
	if len(free.UID) > 0 {
		sql = fmt.Sprintf("%s %s uid = ?", sql, comma)
		params = append(params, free.UID)
		comma = "AND"
	}
	if len(free.Name) > 0 {
		sql = fmt.Sprintf("%s %s name = ?", sql, comma)
		params = append(params, free.Name)
		comma = "AND"
	}
	var frees []*entity.FreeMovie
	var total int
	db := getSysDB().Model(&entity.FreeMovie{}).Where(sql, params...).Count(&total)
	if db.Error != nil {
		return nil, -1, db.Error
	}
	sql = util.PageFilterSql(sql, "id", filter)
	db = getSysDB().Where(sql, params...).Find(&frees)
	return frees, total, db.Error
}

func UpdateFreeMovie(free entity.FreeMovie, querys map[string]interface{}) error {
	up := make(map[string]interface{}, 20)
	if len(free.UID) > 0 {
		up["uidl"] = free.UID
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
		db := getSysDB().Model(&entity.FreeMovie{}).Updates(up).Where(sql, params...)
		return db.Error
	}
}

func DelFreeMovie(free entity.FreeMovie) error {
	if free.ID > 0 {
		db := getSysDB().Delete(&free)
		return db.Error
	}
	sql := ""
	var params []interface{}
	comma := ""
	if len(free.UID) > 0 {
		sql = fmt.Sprintf("%s %s uid = ?", sql, comma)
		params = append(params, free.UID)
		comma = "AND"
	}
	if len(free.Name) > 0 {
		sql = fmt.Sprintf("%s %s name = ?", sql, comma)
		params = append(params, free.Name)
		comma = "AND"
	}
	db := getSysDB().Delete(&entity.FreeMovie{}).Where(sql, params...)
	return db.Error
}
