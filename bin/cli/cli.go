package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/clarch/handler/cli"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	//all technology init here
	db, err := sqlx.Connect("postgres", "user=postgres dbname=election sslmode=disable password=postgres host=localhost")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	controller := cli.CliController{
		DB: db,
		//any dependency for controller injected here
	}

	for true {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("============= ELECTION MENU =============")
		fmt.Println("PICK MENU BY ENTER THE MENU NUMBER")
		fmt.Println("<1> Join Election")
		selected, _ := reader.ReadString('\n')
		selected = strings.TrimSuffix(selected, "\n")
		var err error
		optionSelected, err := strconv.Atoi(selected)
		if err != nil {
			fmt.Println(err)
		}

		print("\033[H\033[2J")

		switch optionSelected {
		case 1:
			fmt.Println("Input your VoterID : ")
			voterID, _ := reader.ReadString('\n')
			voterID = strings.TrimSuffix(voterID, "\n")
			fmt.Println("Input your ElectionID : ")
			electionID, _ := reader.ReadString('\n')
			electionID = strings.TrimSuffix(electionID, "\n")

			controller.CJoinElection(cli.JoinElectionRequest{
				VoterID:    voterID,
				ElectionID: electionID,
			})
		}
	}
}
