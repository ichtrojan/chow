package chow

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/urfave/cli"
)

var data = []string{"Date Joined", "username", "Full Name", "E-mail"}

// HandleFunc handles everything that pertains with this CLI application
func HandleFunc(c *cli.Context) error {
	path := c.Args().Get(0)

	if path == "" {
		errorMessage := "No path specified: Please include a path."
		return errors.New(errorMessage)
	}

	jsonByte, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	var format []UsersData
	err = json.Unmarshal(jsonByte, &format)

	if err != nil {
		return err
	}

	timestamp := time.Now()

	file, err := os.Create("./chow_output" + timestamp.String() + ".csv")

	if err != nil {
		return err
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
