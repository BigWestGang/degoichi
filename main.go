package main

import (
	"bufio"
	"fmt"
	toml "github.com/BurntSushi/toml"
	"github.com/davecgh/go-spew/spew"
	colorable "github.com/mattn/go-colorable"
	errors "github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"os"
)

type Config struct {
	Provider string
	Box      string
}

type Configs struct {
	Config []Config
}

func main() {
	log.SetFormatter(&log.TextFormatter{ForceColors: true})
	log.SetOutput(colorable.NewColorableStdout())
	var config Configs
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		switch errors.Cause(err).(type) {
		case *os.PathError:
			log.Error(err)
			if Question("Would you like to create a config.toml? (y/n) default: y") {
				fmt.Println("yes")
			} else {
				fmt.Println("no")
			}
		default:
			log.Error(err)
			spew.Dump(err)
		}

	}
	for _, s := range config.Config {
		if s.Provider == "vagrant" {

		}
		fmt.Printf("%s (%s)\n", s.Provider, s.Box)
	}
}

func Question(q string) bool {
	result := true
	fmt.Print(q)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		i := scanner.Text()
		if i == "Y" || i == "y" || i == "yes" || i == "" {
			break
		} else if i == "N" || i == "n" || i == "no" {
			result = false
			break
		} else {
			fmt.Println("Please enter y(yes) or n(no)")
			fmt.Print(q)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return result
}
