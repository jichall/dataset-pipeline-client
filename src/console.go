package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func Init() {
	exit := false

	fmt.Println("DPS CLI")
	fmt.Println("=======")
	fmt.Println("Type help for more information")

	reader := bufio.NewReader(os.Stdin)

	for !exit {
		var command string
		//var argument string

		fmt.Print("> ")
		command, _ = reader.ReadString('\n')
		command = strings.Replace(command, "\n", "", -1)

		if command == "help" {
			fmt.Printf("\nAvailable commands are: \n")
			fmt.Println("dir <dir>\t\tSets the directory to read the files ")
			fmt.Println("url <url>\t\tSets the url to send the loaded files to")
			fmt.Println("send\t\t\tLoad and send the files to the server")
			fmt.Printf("get <pk>\t\tGets the data associated with the pk " +
				"attribute\n")
			fmt.Printf("exit\t\t\tExits the terminal\n")

		} else if strings.HasPrefix(command, "url") {
			c, err := checkCommand(command)

			if err != nil {
				fmt.Println(url)
			} else {
				url = c[1]
			}
		} else if strings.HasPrefix(command, "dir") {
			c, err := checkCommand(command)

			if err != nil {
				fmt.Println(dir)
			} else {
				dir = c[1]
			}
		} else if command == "send" {
			var rows [][]JSONData
			var filenames []string

			filenames, rows = load()

			for i := range rows {
				send(filenames[i], rows)
			}
		} else if strings.HasPrefix(command, "get") {
			fmt.Println(strings.Split(command, " "))

			c, err := checkCommand(command)

			if err != nil {
				fmt.Println("[!] Missing pk value on the get command")
			} else {
				fmt.Printf("%s", get(c[1]))
			}
		} else if command == "exit" {
			os.Exit(0)
		} else {
			fmt.Printf("[!] Invalid command %s. Type help to see the available "+
				"commands\n", command)
		}

	}
}

func checkCommand(command string) ([]string, error) {
	var c []string

	c = strings.Split(command, " ")
	if len(c) == 1 {
		err := errors.New("")
		return nil, err
	}

	return c, nil
}

func get(pk string) []byte {
	res, err := http.Get("http://" + url + "/v1/get/" + pk)

	if err != nil {
		fmt.Printf("[!] Error trying to get. Cause %s\n", err.Error())
		defer func() {
			if recover() != nil {
				fmt.Printf("[!] Check the server status.\n")
			}
		}()
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Printf("Error trying to read response. Cause %s", err.Error())
	}

	return body
}

func load() ([]string, [][]JSONData) {
	files, err := Load(dir, ext)
	var filenames []string = make([]string, len(files))

	if err != nil {
		log.Fatalf("[!] Some error occurred on directory. Cause: %v",
			err.Error())
	}

	fmt.Println("[+] Files loaded. Getting the JSON and sending it to the " +
		"database.")

	var data [][]JSONData = make([][]JSONData, len(files))

	// For each file, unmarshall its JSON content and send it to the server
	for i := range files {
		filenames[i] = files[i].Filename
		rows := JSONGet(dir + files[i].Filename)
		data[i] = rows
	}

	return filenames, data
}

func send(filename string, rows [][]JSONData) {
	for i := range rows {
		for j := range rows[i] {
			_, err := http.Get("http://" + url + "/v1/new/?filename=" +
				filename + "&pk=" + rows[i][j].Pk + "&score=" +
				rows[i][j].Score)

			if err != nil {
				fmt.Printf("[!] Send data failed. Cause %s\n", err.Error())
				defer func() {
					if recover() != nil {
						fmt.Printf("[!] Check the server status.\n")
					}
				}()
			}
		}
	}
}
