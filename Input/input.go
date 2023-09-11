package Input

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func Input() {
	scanner := bufio.NewScanner(os.Stdin)

	usernameSli := []string{}
	userID := []int{}

	db, err := sql.Open("sqlite3", "security.db")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Not Your Enter Real ID Just Enter Like 1, 4, 6, 7")
	fmt.Print("Enter Your ID:")
	scanner.Scan()
	id, err := strconv.ParseInt(scanner.Text(), 10, 32)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Enter Your Username:")
	scanner.Scan()
	username := scanner.Text()

	userID = append(userID, int(id))
	usernameSli = append(usernameSli, username)

	insertStmt := `
INSERT INTO user (id, username) VALUES (?, ?);
`
	changeStmt := `
UPDATE user
SET username = ? 
WHERE id = ?;
`

	fmt.Print("Enter 'first' to insert or 'still' to update: ")
	scanner.Scan()
	secim := strings.ToLower(scanner.Text())

	if secim == "first" {
		for i := 0; i < len(userID); i++ {
			_, err = db.Exec(insertStmt, userID[i], usernameSli[i])
			if err != nil {
				log.Fatal(err)
			}
		}
	} else if secim == "still" {
		for i := 0; i < len(userID); i++ {
			_, err = db.Exec(changeStmt, usernameSli[i], userID[i])
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
