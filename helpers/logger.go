package helpers

import (
	"log/slog"
	"os"
)

func CheckFolder() {
	if _, err := os.Stat("./logs"); os.IsNotExist(err) {
		err := os.Mkdir("./logs", 0755)
		if err != nil {
			panic(err)
		}
	}
}

var Slogger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
