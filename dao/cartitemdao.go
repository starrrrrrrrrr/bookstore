package dao

import (
	"exercise/bookstore/model"
	"exercise/bookstore/utils"
)

//向数据库中插入购物项
func AddCartItem(cartItem *model.CartItem) error {
	//sql语句
	sqlStr := "insert into cart_items (count, amount, book_id, cart_id) values(?,?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, cartItem.Count, cartItem.GetAmount(), cartItem.Book.Id, cartItem.CartId)
	if err != nil {
		return err
	}
	return nil
}

//从数据库中删除购物项
func DeleteCartItem(id string) error {
	//sql语句
	sqlStr := "delete from cart_items where id = ?"
	//执行
	inStmt, _ := utils.Db.Prepare(sqlStr)
	_, err := inStmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

//使用bookId从数据库中获取购物项
func GetCartItemByBookIdAndCartId(bookId, CartId string) (*model.CartItem, error) {
	//sql语句
	sqlStr := "select id, count, amount, cart_id from cart_items where book_id=? and cart_id=?"
	row := utils.Db.QueryRow(sqlStr, bookId, CartId)
	//创建CartItem结构体
	cartItem := &model.CartItem{}

	err := row.Scan(&cartItem.CartItemId, &cartItem.Count, &cartItem.Amount, &cartItem.CartId)
	if err != nil {
		return nil, err
	}
	//获取book
	book, _ := GetBookById(bookId)
	cartItem.Book = book
	return cartItem, nil
}

//使用cartId从数据库中获取购物车的所有购物项
func GetCartItemsByCartId(cartId string) ([]*model.CartItem, error) {
	//sql语句
	sqlStr := "select id, count, amount, book_id, cart_id from cart_items where cart_id = ?"
	//执行
	rows, err := utils.Db.Query(sqlStr, cartId)
	if err != nil {
		return nil, err
	}

	//创建cartItems切片
	var cartItems []*model.CartItem
	//扫描rows中的内容
	for rows.Next() {
		//创建cartItem实例接收rows的内容
		var bookId string
		cartItem := &model.CartItem{}
		rows.Scan(&cartItem.CartItemId, &cartItem.Count, &cartItem.Amount, &bookId, &cartItem.CartId)
		//获取图书
		book, _ := GetBookById(bookId)
		cartItem.Book = book
		cartItems = append(cartItems, cartItem)
	}
	return cartItems, nil
}

//更新购物项里的count
func UpdateBookCount(cartItem *model.CartItem) error {
	//sql语句
	sqlStr := "update cart_items set count=?, amount=? where book_id=? and cart_id=?"
	//执行
	_, err := utils.Db.Exec(sqlStr, cartItem.Count, cartItem.GetAmount(), cartItem.Book.Id, cartItem.CartId)
	if err != nil {
		return err
	}
	return nil
}

//根据CartId删除所有购物项
func DeleteCartItemsByCartId(cartId string) error {
	//sql语句
	sqlStr := "delete from cart_items where cart_id=?"
	//执行
	_, err := utils.Db.Exec(sqlStr, cartId)
	if err != nil {
		return err
	}
	return nil
}

//根据id删除cartItem
func DeleteCartItemById(id string) error {
	//sql语句
	sqlStr := "delete from cart_items where id=?"
	//执行
	_, err := utils.Db.Exec(sqlStr, id)
	if err != nil {
		return err
	}
	return nil
}
