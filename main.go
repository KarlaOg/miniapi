package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", tellMeTime)
	http.HandleFunc("/hello", addEntries)
	http.HandleFunc("/entries", getEntries)
	http.ListenAndServe(":4567", nil)
}

func tellMeTime(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		current_time := time.Now()
		fmt.Fprintf(w, current_time.Format("15h04"))
	}
}
func addEntries(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:

		m := make(map[string]string)

		var datas []string

		if err := req.ParseForm(); err != nil {
			fmt.Println("Something went bad")
			fmt.Fprintln(w, "Something went bad")

			return
		}
		for key, value := range req.PostForm {

			switch key {
			case "entry":
				datas = append(datas, value[0])
			}

			m[key] = value[0]

		}
		fmt.Fprintln(w, m["author"], ":", m["entry"])

		f, err := os.OpenFile("myfile.txt",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		defer f.Close()

		sep := "\n"

		for _, value := range datas {

			if _, err := f.WriteString(value + sep); err != nil {
				log.Println(err)

			}

		}

	}

}

func getEntries(w http.ResponseWriter, req *http.Request) {

	switch req.Method {

	case http.MethodGet:

		f, err := os.Open("myfile.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)

		for scanner.Scan() {

			fmt.Fprintln(w, scanner.Text())

		}

		if err := scanner.Err(); err != nil {

			log.Fatal(err)
		}
	}
}
