package account

import (
	"MuShare/datatype"
	"MuShare/db/models"
	"net/http"
	"time"
	"regexp"
  "fmt"
  "MuShare/datatype/request/user"
)

func (this *Account) Register(body *user.Account)  datatype.Response{
	var res datatype.Response
	//check mail
	reg := regexp.MustCompile(`^(\w)+(\.\w+)*@(\w)+((\.\w+)+)$`)
	sel := [...]bool{true, true, true}
  fmt.Println(sel)
	u := models.User{}
	flag := 0
  // begin transaction
  tx := this.DB.Begin()
  if (body.Password == "") {
    goto BadRequest
  }

  if body.Name == "" && body.Mail == "" && body.Phone == "" {
    goto BadRequest
  }

  if reg.FindAllString(body.Mail, -1) == nil {
    goto BadRequest
  }
	if body.Name != "" {
    tx.Where("name=?", body.Name).First(&u)
		sel[0] = checkUser(u)
  }
	if body.Mail != "" {
    tx.Where("mail=?", body.Mail).First(&u)
		sel[1] = checkUser(u)
  }
	if body.Phone != "" {
    tx.Where("phone=?", body.Phone).First(&u)
		sel[2] = checkUser(u)
  }

	for i, v := range sel{
		if v {
			flag = i
			goto Forbidden
		}
	}

	CreateUser(&u, body)
	tx.Create(&u)
  // transaction commit
  tx.Commit()

	res = datatype.Response{
		Status: http.StatusOK,
		Body:u,
	}
	return res

	BadRequest:
	res = datatype.Response{
    Status: http.StatusBadRequest,
  }
  return res

	Forbidden:
	var resText string
	switch flag {
	case 0:
		resText = "Name existed"
	case 1:
		resText = "Mail existed"
	case 2:
		resText = "Phone existed"
	}
	res = datatype.Response{
		Status:http.StatusForbidden,
		ResponseText: resText,
	}
	return res


}

func checkUser(user models.User) bool{
	if user.ID == 0{
		return false
	}
	return true

}

func CreateUser(u *models.User,body *user.Account){
	u.Mail = body.Mail
	u.Name = body.Name
	u.Phone = body.Phone
	u.CreatedAt = time.Now().Unix()
	u.UpdatedAt = time.Now().Unix()
	u.Password = body.Password
}