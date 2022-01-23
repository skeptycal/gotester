package gotester

import (
	"os"

	errorlogger "github.com/skeptycal/errorlogger"
)

var (
	log      = errorlogger.Log
	checkErr = errorlogger.Err
	pwd      string
)

func init() {
	var err error
	pwd, err = os.Getwd()
	die(err)
}

// check will log an error message if err is not nil
func check(err error) error {
	if err == nil {
		return err
	}
	return checkErr(err)
}

// die will exit with an error message if err is not nil
func die(err error) {
	if err == nil {
		return
	}
	log.Fatal(err.Error())
}
