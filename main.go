package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/urfave/cli"
)

var data = []string{"Date Joined", "username", "Full Name", "E-mail"}

func main() {
	app := &cli.App{
		Name:  "chow",
		Usage: "Parse the JSON file",
		Action: func(c *cli.Context) error {
			path := c.Args().Get(0)

			if path == "" {
				log.Fatal("Please include a valid file")
			}

			jsonByte, err := ioutil.ReadFile(path)

			if err != nil {
				log.Fatal(err)
			}

			var format []UsersData
			_ = json.Unmarshal(jsonByte, &format)

			file, err := os.Create("./result.csv")

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
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
