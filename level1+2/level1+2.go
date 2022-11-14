package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	users := make(map[string]string)
	router.GET("/cookie", func(c *gin.Context) {
		cuname, err1 := c.Cookie("username")
		cpass, err2 := c.Cookie("password")
		//如果没有cookie就设置cookie
		if err1 != nil && err2 != nil {
			username := c.DefaultQuery("username", "null")
			password := c.DefaultQuery("password", "null")
			//如果输入了用户名和密码,且不存在 或者输入了存在且密码正确
			v, ok := users[username]
			if username != "null" && password != "null" {
				if !ok || ok && v == password {
					if !ok {
						users[username] = password
						c.SetCookie("password", password, 3600, "/", "127.0.0.1", false, true)
						c.SetCookie("username", username, 3600, "/", "127.0.0.1", false, true)
						c.JSON(200, "登录成功,"+username+"你的密码是:"+password+"请牢记")
					} else {
						c.SetCookie("password", password, 3600, "/", "127.0.0.1", false, true)
						c.SetCookie("username", username, 3600, "/", "127.0.0.1", false, true)
						c.JSON(200, "登录成功,"+username+"您是已经注册的账户")
					}

				} else { //如果存在且密码错误
					c.JSON(200, "密码输入错误")
				}
			} else {
				c.JSON(200, "您未登录，且未输入账号或密码")
			}
			//输入了用户名和密码，且存在（又分为密码输对和输错）

		} else {
			c.JSON(200, "你已登录，"+cuname+cpass)
		}

	})

	router.Run(":8000")
}
