package main

import (
	"bufio"
	// "flag"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func main() {
	fmt.Fprint(os.Stderr, "Connecting to DB\n")
	con, err := sql.Open("mysql", "root:@tcp(localhost:3306)/Shovel")
	bio := bufio.NewReader(os.Stdin)
	var hasMoreInLine bool = true
	for hasMoreInLine {
		line, hasMoreInLine, err := bio.ReadLine()

	}
}

func InsertIntoDB(input string) {

}
