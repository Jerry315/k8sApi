package helpers

import "github.com/GehirnInc/crypt/apr1_crypt"

//模拟htpasswd生成 密码
func HashApr1(password string)  string  {
	s,err:= apr1_crypt.New().Generate([]byte(password), nil)
	if err!=nil{
		panic(err)
	}
	return s
}

