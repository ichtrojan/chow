package chow

import (
	"flag"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/urfave/cli"
)

func TestDummy(t *testing.T) {
	total := 5 + 5

	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

func TestChowWithEmptyPath(t *testing.T) {
	testArgs := []string{""}
	app := &cli.App{Writer: ioutil.Discard}
	set := flag.NewFlagSet("test", 0)
	_ = set.Parse(testArgs)

	context := cli.NewContext(app, set, nil)

	err := HandleFunc(context)

	errorMessage := "No path specified: Please include a path."
	if err.Error() != errorMessage {
		t.Error("Should return an error saying no file specified.")
	}
}

func TestChowWithNonExistentFile(t *testing.T) {
	nonExistentFilePath := "fixtures/fakefake.json"
	testArgs := []string{nonExistentFilePath}
	app := &cli.App{Writer: ioutil.Discard}
	set := flag.NewFlagSet("test", 0)
	_ = set.Parse(testArgs)

	context := cli.NewContext(app, set, nil)

	err := HandleFunc(context)

	errorMessage := fmt.Sprintf("open %s: no such file or directory", nonExistentFilePath)
	if err.Error() != errorMessage {
		t.Error("File shouldn't exist.")
	}
}

func TestChowWithBadFile(t *testing.T) {
	testArgs := []string{"fixtures/fake.json"}
	app := &cli.App{Writer: ioutil.Discard}
	set := flag.NewFlagSet("test", 0)
	_ = set.Parse(testArgs)

	context := cli.NewContext(app, set, nil)

	err := HandleFunc(context)

	errorMessage := "json: cannot unmarshal object into Go value of type []chow.UsersData"
	if err.Error() != errorMessage {
		t.Error("Shoudln't unmarshal bad json into UsersData.")
	}
}
