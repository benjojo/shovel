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
	DBHost := flag.String("host", "localhost:3306", "<hostname>:<port>")
	DBName := flag.String("database", "shovel", "<dbname>")
	DBUser := flag.String("user", "root", "<dbuser>")
	DBPass := flag.String("pass", "", "<dbpass>")
	Tablename := flag.String("tablename", "", "<tablename> else it will make a new one")
	flag.Parse()
	logger := log.New(os.Stderr, "[Shovel]", log.Ltime)
	logger.Println("Connecting to DB")
	con, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", *DBUser, *DBPass, *DBHost, *DBName))
	defer con.Close()
	if err != nil {
		logger.Fatalln("Unable to connect to the database, Aborting.")
	}
	TN := ""
	if *Tablename == "" {
		TN = MakeNewTable(con, logger)
	} else {
		TN = *Tablename
	}
	logger.Printf("Logging line by line into %s.", TN)
	bio := bufio.NewReader(os.Stdin)
	q, e := con.Prepare(fmt.Sprintf("INSERT INTO %s (`line`) VALUES (?);", TN))
	if e != nil {
		logger.Fatalln("Could not prepare the insert query")
		return
	}
	for {
		line, _, err := bio.ReadLine()
		if err != nil {
			logger.Fatalln("Failed to read from stdin")
			break
		}
		_, e = q.Exec(string(line))
		if e != nil {
			logger.Fatalln("Failed to write to the DB")
			break
		}
	}
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
	_, e := DB.Exec(fmt.Sprintf("CREATE TABLE %s (`id` INT NOT NULL AUTO_INCREMENT,`line` TEXT NOT NULL,`logtime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,PRIMARY KEY (`id`))COLLATE='utf8_general_ci';", NewTableName))
	if e != nil {
		logger.Fatalln("Could not make a table to store the data in.")
	}
	return NewTableName
}
