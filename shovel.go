package main

import (
	"bufio"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func main() {
	fmt.Fprint(os.Stderr, "Connecting to DB\n")
	DBHost := flag.String("host", "localhost:3306", "<hostname>:<port>")
	DBName := flag.String("database", "Shovel", "<dbname>")
	DBUser := flag.String("user", "root", "<dbuser>")
	DBPass := flag.String("pass", "", "<dbpass>")
	Tablename := flag.String("tablename", "", "<tablename> else it will make a new one")

	con, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", *DBUser, *DBPass, *DBHost, *DBName))

	bio := bufio.NewReader(os.Stdin)
	var hasMoreInLine bool = true
	for hasMoreInLine {
		line, hasMoreInLine, err := bio.ReadLine()

	}
}

func InsertIntoDB(input string) {

}
