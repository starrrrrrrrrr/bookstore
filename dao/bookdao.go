package dao

import (
	"exercise/bookstore/model"
	"exercise/bookstore/utils"
	"strconv"
)

//获取所有图书
func GetBooks() ([]*model.Book, error) {
	//sql语句
	sqlStr := "select * from books"
	//执行
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}

	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		//给book赋值
		rows.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		//将book添加到books中
		books = append(books, book)
	}

	return books, nil
}

//添加图书
func AddBook(b *model.Book) error {
	//sql语句
	sqlStr := "insert into books(title, author, price, sales, stock, img_path) values(?,?,?,?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ImgPath)
	if err != nil {
		return err
	}
	return nil
}

//删除图书
func DeleteBook(id string) error {
	//sql语句
	sqlStr := "delete from books where id = ?"
	//执行
	_, err := utils.Db.Exec(sqlStr, id)
	if err != nil {
		return err
	}
	return nil
}

//通过id获取图书信息
func GetBookById(id string) (*model.Book, error) {
	//sql语句
	sqlStr := "select * from books where id = ?"

	book := &model.Book{}
	row := utils.Db.QueryRow(sqlStr, id)
	//给book赋值
	err := row.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
	if err != nil {
		return nil, err
	}
	return book, nil
}

//更改数据库中图书信息
func UpdateBook(b *model.Book) error {
	//sql语句
	sqlStr := "update books set title=?,author=?,price=?,sales=?,stock=? where id=?"
	//执行
	_, err := utils.Db.Exec(sqlStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.Id)
	if err != nil {
		return err
	}
	return nil
}

//获取图书分页
func GetPage(pageno string) (*model.Page, error) {
	pageNo, _ := strconv.ParseInt(pageno, 10, 64)
	//获取总记录数
	var totalRecord int64
	sqlStr := "select count(*) from books"
	//执行
	row := utils.Db.QueryRow(sqlStr)
	row.Scan(&totalRecord)
	//设置页面大小
	var pageSize int64 = 4
	//计算总页数
	var totalPage int64
	if totalRecord%pageSize == 0 {
		totalPage = totalRecord / pageSize
	} else {
		totalPage = totalRecord/pageSize + 1
	}
	//从数据库中获取图书
	sqlStr = "select * from books limit ?,?"
	rows, err := utils.Db.Query(sqlStr, (pageNo-1)*4, pageSize)
	if err != nil {
		return nil, err
	}
	//创建图书切片
	var books []*model.Book
	//获取图书
	for rows.Next() {
		book := &model.Book{}
		err = rows.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	page := &model.Page{
		Books:       books,
		PageNo:      pageNo,
		PageSize:    pageSize,
		TotalPage:   totalPage,
		TotalRecord: totalRecord,
	}
	return page, nil
}

//获取带价格范围图书分页
func GetPageBooksByPrice(pageno, min, max string) (*model.Page, error) {
	pageNo, _ := strconv.ParseInt(pageno, 10, 64)
	//获取总记录数
	var totalRecord int64
	sqlStr := "select count(*) from books where price between ? and ?"
	//执行
	row := utils.Db.QueryRow(sqlStr, min, max)
	row.Scan(&totalRecord)
	//设置页面大小
	var pageSize int64 = 4
	//计算总页数
	var totalPage int64
	if totalRecord%pageSize == 0 {
		totalPage = totalRecord / pageSize
	} else {
		totalPage = totalRecord/pageSize + 1
	}
	//从数据库中获取图书
	sqlStr = "select * from books where price between ? and ? limit ?,?"
	rows, err := utils.Db.Query(sqlStr, min, max, (pageNo-1)*4, pageSize)
	if err != nil {
		return nil, err
	}
	//创建图书切片
	var books []*model.Book
	//获取图书
	for rows.Next() {
		book := &model.Book{}
		err = rows.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	page := &model.Page{
		Books:       books,
		PageNo:      pageNo,
		PageSize:    pageSize,
		TotalPage:   totalPage,
		TotalRecord: totalRecord,
	}
	return page, nil
}
