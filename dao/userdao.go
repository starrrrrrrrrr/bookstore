package dao

import (
	"exercise/bookstore/model"
	"exercise/bookstore/utils"
)

func CheckUsernameAndPassword(username, password string) (*model.User, error) {
	sqlstr := "select * from users where username=? and password=?"

	row := utils.Db.QueryRow(sqlstr, username, password)

	user := &model.User{}

	row.Scan(&user.Id, &user.Username, &user.Password, &user.Email)

	return user, nil
}

func CheckUsername(username string) (*model.User, error) {
	sqlStr := "select * from users where username=?"

	row := utils.Db.QueryRow(sqlStr, username)

	user := &model.User{}

	row.Scan(&user.Id, &user.Username, &user.Password, &user.Email)

	return user, nil
}

func SaveUser(username, password, email string) error {
	sqlStr := "insert into users(username, password, email) values(?,?,?)"

	_, err := utils.Db.Exec(sqlStr, username, password, email)

	return err
}
