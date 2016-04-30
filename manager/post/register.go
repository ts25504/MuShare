package post

import (
	"MuShare/datatype"
	"MuShare/db/models"
	"net/http"
	"time"
	"regexp"
)

func (this *Post) Register(body *datatype.RegisterBody)  datatype.Response{
	var res datatype.Response
	reg := regexp.MustCompile(`^(\w)+(\.\w+)*@(\w)+((\.\w+)+)$`)
	sel := make([]bool, 3)
	u := models.User{}
	flag := 0
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
    this.DB.Where("name=?", body.Name).First(&u)
		sel[2] = checkUser(u)
  }
	if body.Mail != "" {
    this.DB.Where("mail=?", body.Mail).First(&u)
		sel[0] = checkUser(u)
  }
	if body.Phone != "" {
    this.DB.Where("phone=?", body.Phone).First(&u)
		sel[1] = checkUser(u)
  }

	for i, v := range sel{
		if v {
			flag = i
			goto Forbidden
		}
	}

	CreateUser(&u, body)
	this.DB.Create(&u)

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
		resText = "Mail existed"
	case 1:
		resText = "Phone existed"
	case 2:
		resText = "Name existed"
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

func CreateUser(u *models.User,body *datatype.RegisterBody){
	u.Mail = body.Mail
	u.Name = body.Name
	u.Phone = body.Phone
	u.CreatedAt = time.Now().Unix()
	u.UpdateAt = time.Now().Unix()
	u.Password = body.Password
}