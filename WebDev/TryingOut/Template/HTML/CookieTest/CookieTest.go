package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", counter)
	http.HandleFunc("/favicon.ico", http.NotFound)
	http.ListenAndServe(":8080", nil)
}

func counter(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("counter")
	if err == http.ErrNoCookie {
		fmt.Println("No such cookie exists a neef")
		cookie = &http.Cookie{
			Name:  "my-counter-cookie",
			Value: "0",
		}
	}
	counter, err := strconv.Atoi(cookie.Value)
	if err != nil {
		fmt.Fprintln(w, "No cookie found man,", err)
	}
	counter++
	cookie.Value = strconv.Itoa(counter)
	http.SetCookie(w, cookie)
	fmt.Fprintln(w, cookie)
}
