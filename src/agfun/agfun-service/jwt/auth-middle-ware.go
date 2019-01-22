package jwt

import (
	"agfun/agfun-service/dbcentral/etcddb"
	"agfun/agfun-service/dbcentral/mysqldb"
	"agfun/agfun-service/util"
	"agfun/auth/entity"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strings"
)

func AuthMiddleWare(authDB *gorm.DB, cli *etcddb.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开关
		if true {
			c.Next()
			return
		}
		session := c.GetHeader("session")
		var ID string
		e := cli.Get(session, &ID)
		if e != nil {
			util.Fail(c, e)
			c.Abort()
			return
		}
		var userRoles []*entity.UserRole
		db := authDB.Find(&userRoles, "user_id = ?", ID)
		if db.Error != nil {
			util.Fail(c, db.Error)
			c.Abort()
			return
		}
		verb := entity.Verb{
			ID:   "",
			Name: "GET",
		}
		switch c.Request.Method {
		case "GET":
			verb.Name = "GET"
		case "POST":
			verb.Name = "POST"
		case "UPDATE":
			verb.Name = "UPDATE"
		case "DELETE":
			verb.Name = "DELETE"
		}
		db = authDB.Where("name=?", verb.Name).Find(&verb)
		if db.Error != nil {
			util.Fail(c, db.Error)
			c.Abort()
			return
		}

		currentParts := splitPath(strings.Split(c.Request.RequestURI, "?")[0])

		for _, userRole := range userRoles {
			rule := entity.Rule{
				RoleID: userRole.RoleID,
				VerbID: verb.ID,
			}
			var resources []*entity.Resource
			db := authDB.Select("resources.name, resources.layer").
				Table("rules").
				Joins("INNER JOIN resources ON rules.resource_id = resources.id").
				Where("rules.role_id =? AND rules.verb_id=?", rule.RoleID, rule.VerbID).
				Find(&resources)
			if db.Error != nil {
				continue
			}

			for _, resource := range resources {
				resourceParts := splitLayer(resource.ParentID)
				resourceParts = append(resourceParts, resource.Name)
				var length int
				if len(resourceParts) == 1 {
					if resourceParts[0] == "authMgr" {
						if len(currentParts) >= 1 {
							if currentParts[0] == "users" ||
								currentParts[0] == "roles" ||
								currentParts[0] == "resources" {
								c.Next()
								return
							}
						}
					}
				}

				for _, resourcePart := range resourceParts[1:] {
					for j := 0; j < len(currentParts); j++ {
						if resourcePart != currentParts[j] {
							continue
						}
						currentParts = append(currentParts[:j], currentParts[j+1:]...)
						j--
						length++
						if length == len(resourceParts)-1 {
							c.Next()
							return
						}
						break
					}
				}
			}
		}
		util.Fail(c, fmt.Errorf("no auth"))
		c.Abort()
	}
}



// splitPath returns the segments for a URL path.
func splitPath(path string) []string {
	path = strings.Trim(path, "/")
	if path == "" {
		return []string{}
	}
	split := strings.Split(path, "/")
	return split
}

func splitLayer(layer string) []string {
	layer = strings.Trim(layer, "-")
	if layer == "" {
		return []string{}
	}
	return strings.Split(layer, "-")
}

func GetResourceParts(parentID string) (string, error) {
	if parentID == "0" {
		return "", nil
	}
	resource := entity.Resource{}
	db := mysqldb.GetAuthDB().Where("id = ?", parentID).First(&resource)
	if db.Error != nil {
		return "", db.Error
	}
	url := ""
	if len(resource.Name) > 0 {
		url = fmt.Sprintf("%s/%s", resource.Name, url)
	}
	if len(resource.Type) > 0 {
		url = fmt.Sprintf("%s/%s", resource.Type, url)
	}
	s, e := GetResourceParts(resource.ParentID)
	if e != nil {
		return "", e
	}
	if len(s) > 0 {
		url = fmt.Sprintf("%s/%s", s, url)
	}
	return url, nil
}

func IsResources(path string) (bool, error) {

	//sql := fmt.Sprintf(`SELECT resource, parent_id FROM sys_resource WHERE "name" = ''`)
	//rows, e := db.Conn.Query(sql)
	//if e != nil {
	//	return false, e
	//}
	//defer rows.Close()
	//var resources []types.Resource
	//for rows.Next() {
	//	var resource types.Resource
	//	e = rows.Scan(&resource.Resource, &resource.ParentId)
	//	if e != nil {
	//		return false, e
	//	}
	//	resources = append(resources, resource)
	//}
	//if e = rows.Err(); e != nil {
	//	return false, e
	//}
	//b, e := IsResursIncludePath(db, resources, path)
	//return b, e
	return false, nil
}

func IsResursIncludePath(authDB *pg.AuthCentralDB, resources []types.Resource, path string) (bool, error) {
//Label:
//	for _, resource := range resources {
//		resourceParts, e := GetResourceParts(authDB, resource.ParentId)
//		if e != nil {
//			continue
//		}
//		if len(resource.Resource) > 0 {
//			resourceParts = fmt.Sprintf("%s/%s", resourceParts, resource.Resource)
//		}
//		if len(resource.Name) > 0 {
//			resourceParts = fmt.Sprintf("%s/%s", resourceParts, resource.Name)
//		}
//		list := strings.Split(resourceParts, "/")
//		if len(list) <= 1 {
//			return false, nil
//		}
//		for i := 1; i < len(list); i++ {
//			if !strings.Contains(path, list[i]) {
//				continue Label
//			}
//		}
//		return true, nil
//	}
//	return false, nil
return false, nil
}
