package main

import (
	"net/http"
	"html/template"
	"log"
	"io/ioutil"
	"strings"
	"time"
	// "path/filepath"
	// "fmt"
)


var templates = template.Must(template.ParseFiles(
								"tmpl/404.html",
								"tmpl/default.html",
								"tmpl/default_footer.html",
								"tmpl/default_header.html",
								"tmpl/post.html"))

type Post struct {
	Status  string
	Title   string
	Date    string
	Summary string
	Body    template.HTML
	File    string
	Year	string
}



func main() {

	// SET UP HANDLE FOR STATIC URLS
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fsd := http.FileServer(http.Dir("data"))
	http.Handle("/data/", http.StripPrefix("/data/", fsd))

	// BESPOKE FUNCTIONS FOR HANDLING OTHER URLS
	http.HandleFunc("/blog/", BlogView)
	http.HandleFunc("/", RootView)

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}


// Here we are implementing the NotImplemented handler. Whenever an API endpoint is hit
// we will simply return the message "Not Implemented"

var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
  w.Write([]byte("Not yet implemented"))
})





func BlogView(w http.ResponseWriter, r *http.Request) {

	// fmt.Println(r.URL.Path[1:])

	if r.URL.Path[1:] == "blog/" {
		// posts := getPosts()
		// t := template.New("default.html")
		// t, _ = t.ParseFiles("default.html")
		// t.Execute(w, posts)
		type PageVariables struct {
			Title string
			Year string
			Content template.HTML
		}


		content := template.HTML("Welcome to the blog")

		p := PageVariables{Title: "blog", Year: GetCurrentTime("2006"), Content: content}

		err := templates.ExecuteTemplate(w, "default.html", p)
	    if err != nil {
	        http.Error(w, err.Error(), http.StatusInternalServerError)
	    }


		return
	}



	// TEST FILE ACTUALLY EXISTS OTHERWISE APP PANICS

	postToLoad := r.URL.Path[1:][5:len(r.URL.Path[1:])]
	// f := "posts/" + r.URL.Path[1:] + ".txt"
	f := "posts/" + postToLoad + ".txt"
	fileread, _ := ioutil.ReadFile(f)
	lines := strings.Split(string(fileread), "\n")
	// fmt.Println(lines)
	// return
	status := string(lines[0])
	title := string(lines[1])
	date := string(lines[2])
	summary := string(lines[3])
	body := strings.Join(lines[4:len(lines)], "\n")
	htmlBody := template.HTML([]byte(body))
	post := Post{status, title, date, summary, htmlBody, r.URL.Path[1:], GetCurrentTime("2006")}


	err := templates.ExecuteTemplate(w, "post.html", post)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

}

func RootView(w http.ResponseWriter, r *http.Request) {

	type PageVariables struct {
		Title string
		Year string
		Content template.HTML
	}


    if r.URL.Path != "/" {
        ErrorHandler(w, r, http.StatusNotFound)
        return
    }


	f := "posts/home.html"
	fileread, _ := ioutil.ReadFile(f)
	lines := strings.Split(string(fileread), "\n")
	// fmt.Println(lines)
	// return
	status := string(lines[0])
	title := string(lines[1])
	date := string(lines[2])
	summary := string(lines[3])
	body := strings.Join(lines[4:len(lines)], "\n")
	htmlBody := template.HTML([]byte(body))
	p := Post{status, title, date, summary, htmlBody, r.URL.Path[1:], GetCurrentTime("2006")}







	// content := template.HTML("LastFuel. Canyoning. Welcome")

	// p := PageVariables{Title: "Home", Year: GetCurrentTime("2006"), Content: content}

	err := templates.ExecuteTemplate(w, "post.html", p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {

    type PageVariables struct {
        Title string
        Year string
        Content string
    }

    w.WriteHeader(status)
    if status == http.StatusNotFound {
        p := PageVariables{Title: "404 Not Found", Year: GetCurrentTime("2006")}

        err := templates.ExecuteTemplate(w, "404.html", p)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    }
    if status == http.StatusUnauthorized {
        redirect(w, r, "/login")
    }
}

func GetCurrentTime(timeFormat string) (string){
	time := time.Now().Local()
	return time.Format(timeFormat)
}

func redirect(w http.ResponseWriter, r *http.Request, url string) {

	var currentPath string = r.URL.Path
    http.Redirect(w, r, url + "?redirect=" + currentPath, 303)
}