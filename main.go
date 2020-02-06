package main

import (
	"fmt"
	"os"

	"github.com/ahsu1230/GolangCITest/arithmetic"
	"github.com/ahsu1230/GolangCITest/strings"

	log "github.com/sirupsen/logrus"
)

func initLogger() {
  // Log as JSON instead of the default ASCII formatter.
  log.SetFormatter(&log.JSONFormatter{})

  // Output to stdout instead of the default stderr
  // Can be any io.Writer, see below for File example
  log.SetOutput(os.Stdout)

  // Only log the warning severity or above.
  log.SetLevel(log.WarnLevel)
}

func main() {
	fmt.Println("Hello")
	fmt.Println(arithmetic.Double(4))
	fmt.Println(strings.Append("jackson", "5"))

	// Log something
	log.WithFields(log.Fields{
    "animal": "walrus",
    "size":   10,
  }).Info("A group of walrus emerges from the ocean")
}
