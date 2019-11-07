package main

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"strings"
)

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
    // Since we'll be getting the hashed password from the DB it
    // will be a string so we'll need to convert it to a byte slice
    byteHash := []byte(hashedPwd)
    err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
    if err != nil {
        log.Println("ERR")
        return false
    } else {
		fmt.Println("OK")
		return true
	}
    
    
}

func loadUserFile(userFile string) map[string]string {
	file, err := os.Open(userFile)
	if err != nul {
		log.Fatal(err)
	}

	reader := bufio.NewScanner(file)
	reader.Split(bufio.ScanLines)

	userMap := make(map[string]string)

	for reader.Scan() {
		u := strings.Split(reader.Text(), ":")
		userMap[u[0]] = u[1]
	}

	return userMap
}

func main() {
	userFile := os.Args[1]


	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		users := loadUserFile(userFile)

		s := strings.Split(scanner.Text(), " ")
		decoded_pass, err  := url.QueryUnescape(s[1])
		decoded_user, err  := url.QueryUnescape(s[0])
		if err != nil {
			log.Fatal(err)
			return
		}
		comparePasswords(users[decoded_user], []byte(decoded_pass))
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}