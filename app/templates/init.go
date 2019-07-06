package main

import (
	"log"
	"os"

	"<%=repoUrl%>/helpr"
)

func init() {
	helpr.SetupLogger(log.New(os.Stderr, "LOGGER:\n", log.Llongfile))

}
