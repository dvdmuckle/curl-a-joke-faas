package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"
)

type Joke struct {
	ID   uint   `gorm:"primary_key"`
	Joke string `gorm:"not null;unique"`
}
type Jsonjoke struct {
	Jokes []string `json:"jokes"`
}

func randjoke(dbFile string) (joke string) {
	db, err := gorm.Open("sqlite3", dbFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return "Error!"
	}
	defer db.Close()
	jokes := []Joke{}
	db.Find(&jokes)
	j := jokes[rand.Intn(len(jokes))]
	return j.Joke
}

func parsejson(jsnFile string) (jokes Jsonjoke) {
	fmt.Println("Parsing json...")
	file, err := ioutil.ReadFile(jsnFile)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(file, &jokes)
	return jokes
}

func setup(db *string, jsn *string) (dbFile string, jsnFile string) {
	dbFile = *db
	jsnFile = *jsn
	rand.Seed(time.Now().UTC().UnixNano())
	return dbFile, jsnFile
}

func dbSetup(dbFile string, newJokes Jsonjoke) {
	fmt.Println("Migrating to database...")
	db, err := gorm.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&Joke{})
	for _, elm := range newJokes.Jokes {
		newrec := Joke{Joke: elm}
		db.Create(&newrec)
	}
	db.Close()
}

func main() {
	dbPtr := flag.String("jokesdb", "jokes.db", "Location to the jokes database")
	jsnPtr := flag.String("jokesjsn", "jokes.json", "Location to the jokes json file")
	flag.Parse()
	dbFile, jsnFile := setup(dbPtr, jsnPtr)
	newJokes := parsejson(jsnFile)
	dbSetup(dbFile, newJokes)
	fmt.Println(randjoke(dbFile))
}
