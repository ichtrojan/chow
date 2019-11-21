package chow

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/urfave/cli"
)

var data = []string{"Date Joined", "username", "Full Name", "E-mail"}

// HandleFunc handles everything that pertains with this CLI application
func HandleFunc(c *cli.Context) error {
	path := c.Args().Get(0)

	if path == "" {
		errorMessage := "No file specified: Please include a valid file"
		log.Fatal(errorMessage)
		return errors.New(errorMessage)
	}

	jsonByte, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
		return err
	}

	var format []UsersData
	_ = json.Unmarshal(jsonByte, &format)

	timestamp := time.Now()

	file, err := os.Create("./chow_output" + timestamp.String() + ".csv")

	if err != nil {
		log.Fatal(err)
	}

	w := csv.NewWriter(file)

	_ = w.Write(data)

	for _, data := range format {
		dateCreated := time.Unix(int64(data.DateCreate), 0)
		fmt.Printf("%s %s %s %s\n", time.Unix(int64(data.DateCreate), 0), data.User.Name, data.User.RealName, data.Email)
		_ = w.Write([]string{dateCreated.String(), data.User.Name, data.User.RealName, data.Email})
	}

	return nil
}
