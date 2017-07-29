package main

import (
	"bufio"
	"fmt"
	toml "github.com/BurntSushi/toml"
	"github.com/davecgh/go-spew/spew"
	color "github.com/fatih/color"
	errors "github.com/pkg/errors"
	"os"
)

type Config struct {
	Provider string
	Box      string
}

type Configs struct {
	Config []Config
}

type Network struct {
	Type  string
	Guest string
	Host  string
	Ip    string
}

type Networks struct {
	Network []Network
}

type Bookshelf struct {
	Dir string
}

type Bookshelves struct {
	Bookshelf []Bookshelf
}

func main() {
	var config Configs
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		switch errors.Cause(err).(type) {
		case *os.PathError:
			color.Red("Not found config.toml file!")
			if Question("Would you like to create a config.toml? (y/n) ") {
				createConfigToml()
			} else {
				fmt.Println("no")
			}
		default:
			spew.Dump(err)
		}

	}
	for _, s := range config.Config {
		if s.Provider == "vagrant" {
		}
		fmt.Printf("%s (%s)\n", s.Provider, s.Box)
	}
}

func createConfigToml() {
	fp, err := os.Create("./config.toml")
	if err != nil {
		color.Red("err")
	}
	defer fp.Close()
	writer := bufio.NewWriter(fp)
	bw := bufio.NewWriter(writer)
	bw.WriteString("[[Config]]\nprovider = \"vagrant\"\nbox = \"bento/ubuntu-16.04\"")
	bw.Flush()
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
