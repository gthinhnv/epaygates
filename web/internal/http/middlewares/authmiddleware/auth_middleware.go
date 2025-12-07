package authmiddleware

import (
	"web/internal/bootstrap"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		menuItems := bootstrap.Config.MenuItems
		// userRole := user.Role.ToInt()
		userRole := 1

		for i := 0; i < len(menuItems); i++ {
			for _, role := range menuItems[i].Roles {
				if role == userRole || true {
					menuItems[i].IsShow = true
					break
				}
			}
			if menuItems[i].IsShow {
				for j := 0; j < len(menuItems[i].Children); j++ {
					for _, childRole := range menuItems[i].Children[j].Roles {
						if childRole == userRole || true {
							menuItems[i].Children[j].IsShow = true
							break
						}
					}
				}

			}
		}

		c.Set("menuItems", menuItems)

		c.Next()
	}
}

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
