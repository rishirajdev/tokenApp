package controllers

import "github.com/revel/revel"
import "github.com/revel/revel/cache"
import "tokenApp/app/models"
import "fmt"
import "crypto/rand"
import "encoding/base64"
//import "golang.org/x/crypto/bcrypt"
//import "tokenApp/app/routes"
import "time"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	token := make(map[string]interface{})
	//id := c.Session["username"]
	fmt.Println(c.Request.Header.Get("Authorization"))
	if err := cache.Get(c.Request.Header.Get("Authorization"), &token); err != nil {
		
        fmt.Println("Token Expired")
		return c.Redirect(App.Notfound)
		
	}
	fmt.Println(token)
	/*if c.Session["username"] != "Rishiraj"{
		fmt.Println("Inside IF 2")
		
	}*/
	return c.Render("Success")
}

func (c App) Authentication() revel.Result {

	userlog := &models.User{
		Username:      c.Request.Form["username"][0],
		Password:   c.Request.Form["password"][0],
	}

	fmt.Println("c.Request.Form : ", c.Request.Form)
	fmt.Println(userlog.Username)

	data := make(map[string]interface{})
	if userlog.Username == "Rishiraj" {
		if userlog.Password == "pass123" {
			fmt.Println("Success")
			token, err := GenerateRandomString(32)
			if err != nil {

				fmt.Println(err)
			}
			//c.Session["username"] = userlog.Username
			//c.Session["token"] = token
			//c.Session.SetDefaultExpiration()
			//c.Session.SetNoExpiration()
			data["username"] = userlog.Username
			data["access_token"]="Bearer "+token
			go cache.Set("Bearer "+token, data, 20*time.Second)
			fmt.Println(20*time.Second)

		}

	} else {
		fmt.Println("Incorrect Password")
	}

	return c.RenderJson(data)
	//return apiResp(usersuccess)
}

func (c App) Notfound() revel.Result{
	
	return c.RenderText("404 session expired")
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}