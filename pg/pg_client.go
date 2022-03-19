package pg

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host                 = "minikube"
	port                 = 30432
	user                 = "postgres"
	password             = "thalesdigital"
	dbname               = "blog"
	sqlInsertTpl         = `INSERT INTO blog_item (id,author_id,title,content) VALUES($1, $2, $3, $4) RETURNING id;`
	sqlDeleteTpl         = `DELETE FROM blog_item where id=$1;`
	sqlSelectTpl         = `SELECT id,author_id,title,content FROM blog_item WHERE id=$1;`
	sqlUpdateTpl         = `UPDATE blog_item SET content=$1 WHERE id=$2 AND author_id=$3 RETURNING content;`
	sqlUpdateNoReturnTpl = `UPDATE blog_item SET content=$1 WHERE id=$2 AND author_id=$3;`
)

var (
	db      *sql.DB
	errConn error
)

type BlogItem struct {
	BlogId  string
	Author  string
	Title   string
	Content string
}

func init() {
	pgInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	if db, errConn = sql.Open("postgres", pgInfo); errConn != nil {
		panic(errConn)
	}

}

func PingDb() string {
	if errPing := db.Ping(); errPing != nil {
		panic(errPing)
	}
	//defer db.Close()
	log.Println("Connect to db successfully")
	return "OK"
}
func InsertBlog(bl BlogItem) int {
	var lastInserted string
	if errInsert := db.QueryRow(sqlInsertTpl, bl.BlogId, bl.Author, bl.Title, bl.Content).Scan(&lastInserted); errInsert != nil {

		log.Printf("error inserting record %v", errInsert)
	} else {
		log.Printf("Last inserted %s", lastInserted)
		return 1
	}
	return 0
}
func DeleteSingleBlog(blogId string) int {
	if r, errDel := db.Exec(sqlDeleteTpl, blogId); errDel != nil {
		log.Printf("error deleting record %v\n", errDel)
	} else {
		if recordNum, errAccess := r.RowsAffected(); errAccess != nil {
			log.Printf("error inserting record %v\n", errAccess)
		} else {
			return int(recordNum)
		}
	}
	return 0
}
func SelectBlogById(blogId string) BlogItem {
	row := db.QueryRow(sqlSelectTpl, blogId)
	var id, author, title, content string
	var bl BlogItem
	switch err := row.Scan(&id, &author, &title, &content); err {
	case sql.ErrNoRows:
		log.Printf("No record returned with id %s\n", blogId)
		return bl
	case nil:
		bl = BlogItem{id, author, title, content}
		log.Printf("Got %v ", bl)
		return bl
	default:
		log.Fatalf("Error %v", err)

	}

	return bl
}
func UpdateBlogWithReturn(id string, author string, content string) int {
	var newContent string
	if errUpdate := db.QueryRow(sqlUpdateTpl, content, id, author).Scan(&newContent); errUpdate != nil {
		panic(errUpdate)
	} else {
		if content == newContent {
			return 1
		}
	}
	return 0
}

func UpdateBlogWithAffectedRecords(id string, author string, content string) int {
	if r, updateErr := db.Exec(sqlUpdateNoReturnTpl, content, id, author); updateErr != nil {
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
