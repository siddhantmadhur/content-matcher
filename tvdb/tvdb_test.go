package tvdb

import (
	"fmt"
	"os"
	"testing"

	"github.com/siddhantmadhur/content/client"
)

func TestFetch(t *testing.T) {
	readToken := os.Getenv("TVDB_READ_TOKEN")
	if readToken == "" {
		fmt.Printf("[ERROR]: TVDB_READ_TOKEN env variable not provided.\n")
		t.FailNow()
	}

	var tvdb TVDB
	tvdb.ApiKey = readToken

	var result struct {
		Id    int    `json:"id"`
		Title string `json:"title"`
	}

	err := tvdb.Fetch(client.FetchParams{
		Method:   "GET",
		Endpoint: "/movie/11",
	}, &result)

	if err != nil {
		fmt.Printf("[ERROR]: %s\n", err.Error())
		t.FailNow()
	}

	if result.Title != "Star Wars" {
		fmt.Printf("[ERROR]: \nExpected: Star Wars \nGot: %s \n", result.Title)

		t.FailNow()
	}
}

func TestGetFromId(t *testing.T) {
	readToken := os.Getenv("TVDB_READ_TOKEN")
	if readToken == "" {
		fmt.Printf("[ERROR]: TVDB_READ_TOKEN env variable not provided.\n")
		t.FailNow()
	}

	var tvdb TVDB
	tvdb.ApiKey = readToken

	result, err := tvdb.GetFromId(11)

	if err != nil {
		fmt.Printf("[ERROR]: %s\n", err.Error())
		t.FailNow()
	}

	if result.Title != "Star Wars" {
		fmt.Printf("[ERROR]: Title does not match Star Wars\n")
		t.FailNow()
	}

}
