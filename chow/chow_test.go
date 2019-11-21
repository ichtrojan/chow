package chow

import (
	"flag"
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

	errorMessage := "No file specified: Please include a valid file"
	if err.Error() != errorMessage {
		t.Error("Path should be empty.")
	}
}

// func TestChowWithBadJson(t *testing.T) {
// 	testArgs := []string{"fixtures/fake.json"}
// 	app := &cli.App{Writer: ioutil.Discard}
// 	set := flag.NewFlagSet("test", 0)
// 	_ = set.Parse(testArgs)

// 	context := cli.NewContext(app, set, nil)

// 	err := HandleFunc(context)

// 	errorMessage := "No file specified: Please include a valid file"
// 	if err.Error() != errorMessage {
// 		t.Error("Path should be empty.")
// 	}
// }
