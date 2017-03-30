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
	id := c.Session["username"]
	
	if err := cache.Get("token_"+id, &token); err != nil {
		
        fmt.Println(err)
		fmt.Println(cache.Get("token_"+id, &token))
	}else if token["access_token"] != c.Session["token"]{
		
		return c.Render("Token Expire")
	}
	
	/*if c.Session["username"] != "Rishiraj"{
		fmt.Println("Inside IF 2")
		
	}*/
	return c.Render()
}

func (c App) Authentication() revel.Result {

	userlog := &models.User{
		Username:      c.Request.Form["username"][0],
		Password:   c.Request.Form["password"][0],
	}

	fmt.Println("c.Request.Form : ", c.Request.Form)

	data := make(map[string]interface{})
	if userlog.Username == "Rishiraj" {
		if userlog.Password == "pass123" {
			fmt.Println("Success")
			token, err := GenerateRandomString(32)
			if err != nil {

				fmt.Println(err)
			}
			c.Session["username"] = userlog.Username
			c.Session["token"] = token
			//c.Session.SetDefaultExpiration()
			c.Session.SetNoExpiration()
			data["access_token"] = token
			go cache.Set("token_"+userlog.Username, data, 20*time.Second)
			fmt.Println(1*time.Second)

		}

	} else {
		fmt.Println("Incorrect Password")
	}

	return c.RenderJson(data)
	//return apiResp(usersuccess)
}

func (c App) Notfound() revel.Result{
	
	return c.Render("404 session expired")
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