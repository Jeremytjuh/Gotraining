package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

var correct int
var incorrect int

func main() {
	li, lierr := net.Listen("tcp", ":8080")
	if lierr != nil {
		log.Fatalln(lierr)
	}

	defer li.Close()

	for {
		con, conerr := li.Accept()
		if conerr != nil {
			log.Println(conerr)
			continue
		}
		fmt.Fprintln(con, `Welcome to my math game! To start the game you have to type "Start"`)
		go handle(con)
	}
}

func handle(con net.Conn) {
	scanner := bufio.NewScanner(con)
	for scanner.Scan() {
		li := strings.ToLower(scanner.Text())
		fw := strings.Fields(li)
		if len(fw) != 0 {
			if fw[0] == "start" {
				fmt.Println("Game has been started")
				fmt.Fprintf(con, "You've started the game\n\n")
				rand.Seed(time.Now().UnixNano())
				startgame(rand.Intn(4), con)
			} else if fw[0] == "tutorial" {
				fmt.Println("Tutorial has been requested")
				fmt.Fprintf(con, "The game works as follows, you begin with starting the game, then the game randomly selects a gamemode. The gamemodes being Add, Subtract, Multiply, Divide.\nNONE OF THE COMMANDS ARE CASE SENSITIVE")
			} else if fw[0] == "startmode" {
				if len(fw) >= 2 {
					switch fw[1] {
					case "add":
						fmt.Println("Gamemode Add has been started")
						fmt.Fprintf(con, "You've started gamemode Add\n\n")
						startgame(1, con)
					case "subtract":
						fmt.Println("Gamemode Subtract has been started")
						fmt.Fprintf(con, "You've started gamemode Substract\n\n")
						startgame(2, con)
					case "multiply":
						fmt.Println("Gamemode Multiply has been started")
						fmt.Fprintf(con, "You've started gamemode Multiply\n\n")
						startgame(3, con)
					case "divide":
						fmt.Println("Gamemode Divide has been started")
						fmt.Fprintf(con, "You've started gamemode Divide\n\n")
						startgame(4, con)
					}
				} else {
					fmt.Println("Invalid gamemode has been entered, command: ", fw[0])
					fmt.Fprintf(con, "Invalid gamemode, use help StartMode for more information\n\n")
				}
			} else if fw[0] == "help" {
				if len(fw) >= 2 {
					switch fw[1] {
					case "start":
						fmt.Println("Help start has been requested")
						fmt.Fprintf(con, `You can use the command "Start" to start a random gamemode`)
						fmt.Fprintf(con, "\n\n")
					case "tutorial":
						fmt.Println("Help tutorial has been requested")
						fmt.Fprintf(con, `You can use the command "Tutorial" to request the tutorial, this will explain the basics of the game`)
						fmt.Fprintf(con, "\n\n")
					case "startmode":
						fmt.Println("Help startmode has been requested")
						fmt.Fprintf(con, `You can use the command "StartMode" together with either "Add", "Substract", "Multiply" or "Divide" to start the corresponding gamemode`)
						fmt.Fprintf(con, "\n\n")
					case "information":
						fmt.Println("Help information has been requested")
						fmt.Fprintf(con, `You can use the command "Information" to request some basic information on why this game was created\`)
						fmt.Fprintf(con, "\n\n")
					case "help":
						fmt.Println("Help help has been requested")
						fmt.Fprintf(con, `All help commands:
				"Help Start"
				"Help Tutorial"
				"Help StartMode"
				"Help Information"
				"Help Help"`)
						fmt.Fprintf(con, "\n\n")
					}
				} else {
					fmt.Println("Invalid gamemode has been entered, command: ", fw[0])
					fmt.Fprintf(con, "Invalid help, use help help for more information\n\n")
				}
			} else if fw[0] == "information" {
				fmt.Println("Information has been requested")
				fmt.Fprintf(con, "This game was made to further play with tcp shit -Jeremytjuh\n\n")
			} else {
				fmt.Println("Invalid command has been entered, command: ", fw[0])
				fmt.Fprintln(con, `Invalid command. Type "Help Help" for a list of all the commands`)
				fmt.Fprintf(con, "\n\n")
			}
		}
	}
}

func startgame(gamemode int, con net.Conn) {
	switch gamemode {
	case 1:
		add(con)
	case 2:
		subtract(con)
	case 3:
		multiply(con)
	case 4:
		divide(con)
	}
}

func add(con net.Conn) {
	fmt.Println("Gamemode: Add")
	fmt.Fprintf(con, "Gamemode: Add\n\n")
	for somcount := 0; somcount < 5; somcount++ {
		num1 := rand.Intn(20)
		num2 := rand.Intn(20)
		answer := num1 + num2
		fmt.Fprintf(con, "Question: %d + %d\n", num1, num2)
		scanner := bufio.NewScanner(con)
		for scanner.Scan() {
			typedvalue := scanner.Text()
			if typedvalue != "" {
				val, valerr := strconv.ParseInt(typedvalue, 10, 64)
				if valerr != nil {
					log.Println("Scanner error: ", valerr)
				} else {
					if val == int64(answer) {
						fmt.Println("Correct")
						fmt.Fprintf(con, "Correct answer!!!\n\n")
						correct++
						break
					} else {
						fmt.Println("Incorrect")
						fmt.Fprintf(con, "Incorrect answer...\n\n")
						incorrect++
						break
					}
				}
			}
		}
	}
	fmt.Printf("Final score: %d / 5\n", correct)
	fmt.Println("Game has ended")
	fmt.Fprintf(con, "Your score is: %d / 5\nPress enter to quit", correct)
}

func subtract(con net.Conn) {
	fmt.Println("Gamemode: Subtract")
	fmt.Fprintln(con, "Gamemode: Subtract")
	for somcount := 0; somcount < 5; somcount++ {
		num1 := rand.Intn(20)
		num2 := rand.Intn(20)
		if num1 < num2 {
			num1, num2 = num2, num1
		}
		answer := num1 - num2
		fmt.Fprintf(con, "Question: %d - %d\n", num1, num2)
		scanner := bufio.NewScanner(con)
		for scanner.Scan() {
			typedvalue := scanner.Text()
			if typedvalue != "" {
				val, valerr := strconv.ParseInt(typedvalue, 10, 64)
				if valerr != nil {
					log.Println("Scanner error: ", valerr)
				} else {
					if val == int64(answer) {
						fmt.Println("Correct")
						fmt.Fprintf(con, "Correct answer!!!\n\n")
						correct++
						break
					} else {
						fmt.Println("Incorrect")
						fmt.Fprintf(con, "Incorrect answer...\n\n")
						incorrect++
						break
					}
				}
			}
		}
	}
	fmt.Printf("Final score: %d / 5\n", correct)
	fmt.Println("Game has ended")
	fmt.Fprintf(con, "Your score is: %d / 5\nPress enter to quit", correct)
}

func multiply(con net.Conn) {
	fmt.Println("Gamemode: Multiply")
	fmt.Fprintln(con, "Gamemode: Multiply")
	for somcount := 0; somcount < 5; somcount++ {
		num1 := rand.Intn(10)
		num2 := rand.Intn(10)
		answer := num1 * num2
		fmt.Fprintf(con, "Question: %d x %d\n", num1, num2)
		scanner := bufio.NewScanner(con)
		for scanner.Scan() {
			typedvalue := scanner.Text()
			if typedvalue != "" {
				val, valerr := strconv.ParseInt(typedvalue, 10, 64)
				if valerr != nil {
					log.Println("Scanner error: ", valerr)
				} else {
					if val == int64(answer) {
						fmt.Println("Correct")
						fmt.Fprintf(con, "Correct answer!!!\n\n")
						correct++
						break
					} else {
						fmt.Println("Incorrect")
						fmt.Fprintf(con, "Incorrect answer...\n\n")
						incorrect++
						break
					}
				}
			}
		}
	}
	fmt.Printf("Final score: %d / 5\n", correct)
	fmt.Println("Game has ended")
	fmt.Fprintf(con, "Your score is: %d / 5\nPress enter to quit", correct)
}

func divide(con net.Conn) {
	fmt.Println("Gamemode: Divide")
	fmt.Fprintln(con, "Gamemode: Divide")
	for somcount := 0; somcount < 5; somcount++ {
		num1 := rand.Intn(10)
		num2 := rand.Intn(10)
		num1 = num1 * num2
		answer := num1 / num2
		fmt.Fprintf(con, "Question: %d / %d\n", num1, num2)
		scanner := bufio.NewScanner(con)
		for scanner.Scan() {
			typedvalue := scanner.Text()
			if typedvalue != "" {
				val, valerr := strconv.ParseInt(typedvalue, 10, 64)
				if valerr != nil {
					log.Println("Scanner error: ", valerr)
				} else {
					if val == int64(answer) {
						fmt.Println("Correct")
						fmt.Fprintf(con, "Correct answer!!!\n\n")
						correct++
						break
					} else {
						fmt.Println("Incorrect")
						fmt.Fprintf(con, "Incorrect answer...\n\n")
						incorrect++
						break
					}
				}
			}
		}
	}
	fmt.Printf("Final score: %d / 5\n", correct)
	fmt.Println("Game has ended")
	fmt.Fprintf(con, "Your score is: %d / 5\nPress enter to quit", correct)
}
