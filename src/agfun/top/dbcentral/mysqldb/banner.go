package mysqldb

import (
	"agfun/top/entity"
	"fmt"
)

func GetBanners(banner entity.Banner) ([]entity.Banner, error) {
	sql := ""
	var params []interface{}
	comma := ""
	if !banner.CreateTime.IsZero() {
		sql = fmt.Sprintf("%s %s create_time > ?", sql, comma)
		params = append(params, banner.CreateTime)
		comma = "AND"
	}
	var banners []entity.Banner
	db := getSysDB().Where(sql, params...).Find(&banners)
	return banners, db.Error
}
