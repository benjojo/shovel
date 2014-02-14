package main

import (
	"bufio"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"time"
)

func main() {
	logger := log.New(os.Stderr, "[Shovel]", log.Ltime)
	logger.Println("Connecting to DB\n")
	DBHost := flag.String("host", "localhost:3306", "<hostname>:<port>")
	DBName := flag.String("database", "Shovel", "<dbname>")
	DBUser := flag.String("user", "root", "<dbuser>")
	DBPass := flag.String("pass", "", "<dbpass>")
	Tablename := flag.String("tablename", "", "<tablename> else it will make a new one")
	flag.Parse()
	con, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", *DBUser, *DBPass, *DBHost, *DBName))
	defer con.Close()
	if err != nil {
		logger.Fatalln("Unable to connect to the database, Aborting.")
	}
	TN := ""
	if *Tablename == "" {
		TN = MakeNewTable()
	} else {
		TN = *Tablename
	}
	bio := bufio.NewReader(os.Stdin)
	var hasMoreInLine bool = true
	for hasMoreInLine {
		line, hasMoreInLine, err := bio.ReadLine()
		InsertIntoDB(string(line), TN)
	}
}

func InsertIntoDB(input string, tablename string) {

}

/*
CREATE TABLE `asdasd` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`line` TEXT NOT NULL,
	`logtime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`)
)
COLLATE='utf8_general_ci';
*/
func MakeNewTable(DB *sql.DB, logger *log.Logger) string {
	NewTableName := fmt.Sprintf("Shovel_%d", int32(time.Now().Unix()))
	_, e := DB.Exec("CREATE TABLE ? (`id` INT NOT NULL AUTO_INCREMENT,`line` TEXT NOT NULL,`logtime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,PRIMARY KEY (`id`))COLLATE='utf8_general_ci';", NewTableName)
	if e != nil {
		logger.Fatalln("Could not make a table to store the data in.")
	}
	return NewTableName
}
