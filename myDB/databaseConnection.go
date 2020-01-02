package myDB

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Entry struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
}

type Config struct {

	DBProperties struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Name string `yaml:"name"`
	Address string `yaml:"address"`
	Port string `yaml:"port"`
	} `yaml:"databaseConfig"`
}

func (i Entry) String() string {
	return fmt.Sprintf("Id: %v \n Name: %v \n Age: %v \n", i.Id, i.Name, i.Age)
}


func GetDatabase() * sql.DB {

	f, err := os.Open("properties.yml")
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	//log.Printf("%T", f)
	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Panic(err)
	}
	db, err := sql.Open("mysql", cfg.DBProperties.Username + ":" + cfg.DBProperties.Password + "@tcp(" + cfg.DBProperties.Address + ":" + cfg.DBProperties.Port + ")/" + cfg.DBProperties.Name )
	if err != nil {
		log.Panic(err.Error())
	}
	log.Println( "DB Connection Successful")
	return db
}

func GetAllEntry () (entries []Entry) {

	db := GetDatabase()
	res, err := db.Query("Select * FROM go_demo")
	if err != nil {
		log.Panic(err.Error())
	}

	for res.Next() {
		var i Entry
		err = res.Scan(&i.Id, &i.Name, &i.Age)
		if err != nil {
			log.Panic(err.Error())
		}
		entries = append(entries, i)
		log.Print(i)
	}
	defer db.Close()
	return
}

func GetEntry (id int64) (entry Entry) {

	db := GetDatabase()
	selDB, err := db.Query("SELECT * FROM go_demo WHERE Id=?", id)
	if err != nil {
		panic(err.Error())
	}
	for selDB.Next() {
		err = selDB.Scan(&entry.Id, &entry.Name, &entry.Age)
		if err != nil {
			panic(err.Error())
		}
	}
	defer db.Close()
	return
}

func AddEntry (name string, age int) (err error){

	db := GetDatabase()
	insForm, err := db.Prepare("INSERT INTO go_demo(Name, Age) VALUES(?,?)")
	if err != nil {
		log.Panic(err.Error())
	}
	insForm.Exec(name, age)
	log.Printf("INSERT: Name: %v | Age: %v",name , age)
	defer db.Close()
	return nil
}