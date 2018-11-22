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
func GetFreeVideos(free entity.FreeMovie, filter *util.PageFilter) ([]*entity.FreeMovie, int, error) {
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
