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
		if fw[0] == "start" {
			fmt.Println("Game has been started")
			fmt.Fprintf(con, "You've started the game\n\n")
			rand.Seed(time.Now().UnixNano())
			startgame(rand.Intn(4), con)
		} else if fw[0] == "tutorial" {
			fmt.Println("Tutorial has been requested")
			fmt.Fprintln(con, "The game works as follows, you begin with starting the game, then the game randomly selects a gamemode")
		} else {
			fmt.Println("Invalid command has been entered, command: ", fw[0])
			fmt.Fprintln(con, `Invalid command. Type "Start" to start the game! And "Tutorial" for the tutorial`)
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
