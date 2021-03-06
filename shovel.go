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
	CFG := GetCFG()
	DBHost := flag.String("host", CFG.DBHost, "<hostname>:<port>")
	DBName := flag.String("database", CFG.DBName, "<dbname>")
	DBUser := flag.String("user", CFG.DBUser, "<dbuser>")
	DBPass := flag.String("pass", CFG.DBPass, "<dbpass>")
	Buffer := flag.Bool("buffer", true, "Enable async input from stdin (up to 100 lines buffered before blocking)")
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
	var buf chan string
	if *Buffer || CFG.Buffering {
		buf = make(chan string, 100)
	} else {
		buf = make(chan string)
	}
	go LogRot(buf, q)
	for {
		line, _, err := bio.ReadLine()
		if err != nil {
			logger.Fatalln("Failed to read from stdin")
			break
		}
		buf <- string(line)
	}
}

func LogRot(input chan string, query *sql.Stmt) {
	for line := range input {
		_, e := query.Exec(string(line))
		if e != nil {
			log.Fatalln("Failed to write to the DB")
			return
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
