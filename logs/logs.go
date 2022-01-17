package logs

import (
	"log"
	"os"

	"Work.go/LPG-Bot/LPGBot/filescheck"
	"Work.go/LPG-Bot/LPGBot/print"
)

func CheckAndCreate() (err error) {

	if filescheck.Exists("./logs.txt") {
		print.InfoLog("[INFO] Logs file exists", "[SERVER]")
	} else {
		createlog, err := os.Create("./logs.txt")
		if err != nil {
			log.Fatal(err)
			return err
		}
		createlog.Close()
		print.InfoLog("[INFO] File logs.txt created", "[SERVER]")
	}
	return nil
}
