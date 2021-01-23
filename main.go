package main
import (
	"fmt"
	"net/http"
)
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey, Go is beautifull")
}

func homepageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hey, Your are in home page</h1> ")
	fmt.Fprintf(w, "<p> go is fast</p> ")
	fmt.Fprintf(w, `
	<p> i agree with you that go is fast</p>
	<p>you %s even add %s</p> `, "can", "<strong>variables</strong>")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/home", homepageHandler)
	http.ListenAndServe("0.0.0.0:80", nil)
}
