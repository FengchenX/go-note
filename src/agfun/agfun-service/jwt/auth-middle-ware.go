package jwt

import (
	"agfun/agfun-service/dbcentral/etcddb"
	"agfun/agfun-service/util"
	"agfun/auth/entity"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strings"
)

func AuthMiddleWare(authDB *gorm.DB, cli *etcddb.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 不用检查的url直接放行
		if strings.Contains(c.Request.RequestURI, "/users?user-name") {
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

		currentParts := splitPath(c.Request.RequestURI)

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
				resourceParts := splitLayer(resource.Layer)
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
						currentParts = append(currentParts[:j], currentParts[:j+1]...)
						j--
						length++
						if length == len(resourceParts)-1 {
							c.Next()
							return
						}
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
	return strings.Split(path, "/")
}

func splitLayer(layer string) []string {
	layer = strings.Trim(layer, "-")
	if layer == "" {
		return []string{}
	}
	return strings.Split(layer, "-")
}
