package model

import (
	"fmt"
	// "fmt"
	// "fmt"
	"gin-demo/initDB"
	"log"
)

// 此文件可以理解为操作数据库

type UserModel struct {
	Id            int    `form:"id"`
	Email         string `form:"email"`
	Password      string `form:"password"`
	PasswordAgain string `form:"password-again"`
	Avatar        string
}

func (user *UserModel) Save() int64 {
	result, e := initDB.Db.Exec("insert into lc_member (email, password) values (?,?);", user.Email, user.Password)
	if e != nil {
		log.Panicln("user insert error", e.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Panicln("user insert id error", err.Error())
	}
	return id
}

// 通过邮箱查询
func (user *UserModel) QueryByEmail() UserModel {
	u := UserModel{}
	row := initDB.Db.QueryRow("select * from lc_member where email = ?;", user.Email)
	e := row.Scan(&u.Id, &u.Email, &u.Password)
	if e != nil {
		log.Panicln(e)
	}
	return u
}

// 通过ID查询
func (user *UserModel) QueryById(id int) (UserModel, error) {
	u := UserModel{}
	fmt.Println(id)
	row := initDB.Db.QueryRow("select * from lc_member where id = ?;", id)
	e := row.Scan(&u.Id, &u.Email, &u.Password, &u.Avatar)
	if e != nil {
		log.Panicln(e)
	}
	return u, e
}

// 更新用户信息
func (user *UserModel) Update(id int) error {
	var stmt, e = initDB.Db.Prepare("update lc_member set password=?,avatar=?  where id=? ")
	if e != nil {
		log.Panicln("发生了错误", e.Error())
	}
	_, e = stmt.Exec(user.Password, user.Avatar, user.Id)
	if e != nil {
		log.Panicln("错误 e", e.Error())
	}

	return e
}
