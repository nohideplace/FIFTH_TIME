package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

type Login struct {
	User map[string]string `json:"user,omitempty"`
}

// 定义数组结构体数组存储数据
func readFile(filename string) (Login, error) {

	filePtr, err := os.Open(filename)
	if err != nil {
		fmt.Println("Open file failed [Err:%s]", err.Error())
		return Login{}, nil
	}
	defer filePtr.Close()

	var person Login

	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&person)
	if err != nil {
		fmt.Println("Decoder failed", err.Error())
	} else {
		fmt.Println("Decoder success")
	}
	return person, nil

}
func writeFile(filename string, persons Login) error {
	personInfo := persons
	// 创建文件
	filePtr, err := os.Create(filename)
	if err != nil {
		fmt.Println("Create file failed", err.Error())
		return err
	}
	defer filePtr.Close()

	// 创建Json编码器
	//encoder := json.NewEncoder(filePtr)
	//err = encoder.Encode(personInfo)
	//if err != nil {
	//	fmt.Println("Encoder failed", err.Error())
	//
	//} else {
	//	fmt.Println("Encoder success")
	//}
	//return nil
	//带JSON缩进格式写文件
	data, err := json.MarshalIndent(personInfo, "", "  ")
	if err != nil {
		fmt.Println("Encoder failed", err.Error())

	} else {
		fmt.Println("Encoder success")
	}

	filePtr.Write(data)
	return nil
}

func main() {
	router := gin.Default()
	//定义结构体数组
	router.GET("/cookie", func(c *gin.Context) {

		cuname, err1 := c.Cookie("username")
		cpass, err2 := c.Cookie("password")
		username := c.DefaultQuery("username", "null")
		password := c.DefaultQuery("password", "null")
		//如果没有cookie就设置cookie
		if err1 != nil && err2 != nil {

			//如果输入了用户名和密码,且不存在 或者输入了存在且密码正确
			v, err := readFile("./data.json")
			if err != nil {
				fmt.Println("读取文件失败", err)
			}
			//获取本地的对应账号的密码
			u, ok := v.User[username]
			if username != "null" && password != "null" {
				if !ok || ok && u == password {
					if !ok {
						v.User[username] = password
						c.SetCookie("password", password, 3600, "/", "127.0.0.1", false, true)
						c.SetCookie("username", username, 3600, "/", "127.0.0.1", false, true)
						c.JSON(200, "登录成功,"+username+"你的密码是:"+password+"请牢记")
						writeFile("./data.json", v)
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
			//如果
			if username == cuname && password == cpass {
				c.JSON(200, "你已登录，"+cuname+cpass)
			} else {
				if username != cuname && password != "null" {
					v, err := readFile("./data.json")
					if err != nil {
						fmt.Println("读取文件失败", err)
					}
					v.User[username] = password
					c.SetCookie("password", password, 3600, "/", "127.0.0.1", false, true)
					c.SetCookie("username", username, 3600, "/", "127.0.0.1", false, true)
					c.JSON(200, "登录成功,"+username+"你的密码是:"+password+"请牢记")
					writeFile("./data.json", v)
					c.JSON(200, "更换账号")
				}
				if username == cuname && password != cpass {
					c.JSON(200, "密码错误")
				}
			}

		}

	})

	router.Run(":8000")
}
