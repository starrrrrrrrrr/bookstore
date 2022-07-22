package dao

import (
	"exercise/bookstore/model"
	"exercise/bookstore/utils"
	"net/http"
)

//向数据库中添加session
func AddSession(sess *model.Session) error {
	//sql语句
	sqlStr := "insert into sessions values(?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, sess.SessionId, sess.UserName, sess.UserId)
	if err != nil {
		return err
	}
	return nil
}

//从数据库中删除session
func DelSession(id string) error {
	//sql语句
	sqlStr := "delete from sessions where session_id = ?"
	//执行
	_, err := utils.Db.Exec(sqlStr, id)
	if err != nil {
		return err
	}
	return nil
}

//从数据库中获取session
func GetSession(id string) (*model.Session, error) {
	//sql语句
	sqlStr := "select * from sessions where session_id = ?"
	//执行
	row := utils.Db.QueryRow(sqlStr, id)
	//创建session实例
	session := &model.Session{}
	//获取session元素
	err := row.Scan(&session.SessionId, &session.UserName, &session.UserId)
	if err != nil {
		return nil, err
	}
	return session, nil
}

//IsLogin 判断用户是否登录
func IsLogin(r *http.Request) (bool, *model.Session) {
	//获取cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		//获取cookie的value
		value := cookie.Value
		//判断对应的session是否存在
		session, _ := GetSession(value)
		if session.UserId > 0 {
			return true, session
		}
	}
	return false, nil
}
