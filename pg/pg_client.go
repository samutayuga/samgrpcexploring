package pg

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/samutayuga/samgrpcexploring/blog/cfg"
	"log"
	"os"
)

const (
	sqlInsertTpl         = `INSERT INTO blog_item (id,author_id,title,content) VALUES($1, $2, $3, $4) RETURNING id;`
	sqlDeleteTpl         = `DELETE FROM blog_item where id=$1;`
	sqlSelectTpl         = `SELECT id,author_id,title,content FROM blog_item WHERE id=$1;`
	sqlUpdateTpl         = `UPDATE blog_item SET content=$1 WHERE id=$2 AND author_id=$3 RETURNING content;`
	sqlUpdateNoReturnTpl = `UPDATE blog_item SET content=$1 WHERE id=$2 AND author_id=$3;`
	sqlSelectAllTpl      = `SELECT id,author_id,title,content FROM blog_item LIMIT 1000;`
)

var (
	//Db ...
	Db      *sql.DB
	errConn error
)

type BlogItem struct {
	BlogId  string
	Author  string
	Title   string
	Content string
}

func Init() {
	//retrieve password from env variable which is injected through a secrets
	dbPassword := os.Getenv("DB_PASSWORD")
	cfg := cfg.LoadConfig()
	pgInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DbHost, cfg.DbPort, cfg.DbUser, dbPassword, cfg.DbName)
	if Db, errConn = sql.Open("postgres", pgInfo); errConn != nil {
		panic(errConn)
	}

}

func PingDb() string {
	if errPing := Db.Ping(); errPing != nil {
		panic(errPing)
	}
	//defer db.Close()
	log.Println("Connect to db successfully")
	return "OK"
}
func InsertBlog(bl *BlogItem) int {
	var lastInserted string
	if errInsert := Db.QueryRow(sqlInsertTpl, bl.BlogId, bl.Author, bl.Title, bl.Content).Scan(&lastInserted); errInsert != nil {

		log.Printf("error inserting record %v", errInsert)
	} else {
		log.Printf("Last inserted %s", lastInserted)
		return 1
	}
	return 0
}
func DeleteSingleBlog(blogId string) int {
	if r, errDel := Db.Exec(sqlDeleteTpl, blogId); errDel != nil {
		log.Printf("error deleting record %v\n", errDel)
	} else {
		if recordNum, errAccess := r.RowsAffected(); errAccess != nil {
			log.Printf("error deleting record %v\n", errAccess)
		} else {
			return int(recordNum)
		}
	}
	return 0
}
func SelectBlogById(blogId string) *BlogItem {
	row := Db.QueryRow(sqlSelectTpl, blogId)
	var id, author, title, content string
	switch err := row.Scan(&id, &author, &title, &content); err {
	case sql.ErrNoRows:
		log.Printf("No record returned with id %s\n", blogId)
		return nil
	case nil:
		bl := BlogItem{id, author, title, content}
		log.Printf("Got %v ", bl)
		return &bl
	default:
		log.Fatalf("Error %v", err)

	}

	return nil
}
func UpdateBlogWithReturn(id string, author string, content string) int {
	var newContent string
	if errUpdate := Db.QueryRow(sqlUpdateTpl, content, id, author).Scan(&newContent); errUpdate != nil {
		panic(errUpdate)
	} else {
		if content == newContent {
			return 1
		}
	}
	return 0
}

func UpdateBlogWithAffectedRecords(id string, author string, content string) int {
	if r, updateErr := Db.Exec(sqlUpdateNoReturnTpl, content, id, author); updateErr != nil {
		panic(updateErr)
	} else {
		if rCount, errFetch := r.RowsAffected(); errFetch != nil {
			panic(errFetch)
		} else {
			return int(rCount)
		}
	}
	return 0
}
func SelectAll() []BlogItem {
	if rows, err := Db.Query(sqlSelectAllTpl); err != nil {
		panic(err)
	} else {
		defer rows.Close()
		allBls := make([]BlogItem, 0)
		for rows.Next() {
			var id, author, title, content string
			if errSc := rows.Scan(&id, &author, &title, &content); errSc != nil {
				panic(errSc)
			} else {
				allBls = append(allBls, BlogItem{id, author, title, content})
			}

		}
		return allBls
	}
	return nil
}
func CloseDb() {
	Db.Close()
}
