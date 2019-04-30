package mysqldb

import (
	"agfun/movie/entity"
	"agfun/util"
	"fmt"
)

func GetMovies(movie entity.Movie, filter *util.PageFilter) ([]*entity.Movie, int, error) {
	sql := ""
	var params []interface{}
	comma := ""
	if len(movie.ID) > 0 {
		sql = fmt.Sprintf("%s %s id = ?", sql, comma)
		params = append(params, movie.ID)
		comma = "AND"
	}
	var movies []*entity.Movie
	var total int
	db := getSysDB().Model(&entity.Movie{}).Where(sql, params...).Count(&total)
	if db.Error != nil {
		return nil, -1, db.Error
	}
	sql = util.PageFilterSql(sql, "id", filter)
	db = getSysDB().Where(sql, params...).Find(&movies)
	return movies, total, db.Error
}

func AddMovie(movie *entity.Movie) error {
	db := getSysDB().Create(movie)
	return db.Error
}
