package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/user"
)

type Defaults struct {
	DBHost    string
	DBName    string
	DBUser    string
	DBPass    string
	Buffering bool
}

func GetCFG() Defaults {
	Ret := Defaults{
		DBHost:    "localhost:3306",
		DBName:    "shovel",
		DBUser:    "root",
		DBPass:    "",
		Buffering: false,
	}
	cu, e := user.Current()
	if e != nil {
		// Uhh just go with defaults
		return Ret
	}
	FilePath := fmt.Sprintf("%s/.shovel", cu.HomeDir)

	b, e := ioutil.ReadFile(FilePath)

	if e != nil {
		outjson, _ := json.Marshal(Ret)
		ioutil.WriteFile(FilePath, outjson, 600)
		return Ret
	}

	e = json.Unmarshal(b, &Ret)
	if e != nil {
		Ret = Defaults{
			DBHost: "localhost:3306",
			DBName: "shovel",
			DBUser: "root",
			DBPass: "",
		}
		return Ret
	}
	return Ret
}
