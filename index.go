package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"log"
	
)

func main() {
	count := 10
	for index := 0; index < count; index++ {
		println(index)
	}
	DataTypes()
	ConditionalStatements(17)
	CaseStateMents(18)
	//SimpleHttpServer()
	StartRestApiServices()

}

func CaseStateMents(age int) {
	switch age {
	case 10:
		println("Don't run after girls")
		break
	case 18:
		println("Harmones are running high")
		break
	default:
		break
	}
}

func ConditionalStatements(age int) {
	if age < 5 {
		println("you are a baby")
	} else if age < 12 {
		println("you are a toddler")
	} else if age < 18 {
		println("you are a teen")
	} else {
		println("you are the man")
	}
}

func DataTypes() {
	var name string = "Voidy Walker"
	const pi float64 = 3.14124
	win := true
	x := 5

	fmt.Printf("%.3f \n", pi)
	fmt.Printf("%T \n", name) // prints the DataType as string 'T'
	fmt.Printf("%s \n", name)
	fmt.Printf("%t \n", win)
	fmt.Printf("%d \n", x)
	fmt.Printf("%b \n", 25)
	fmt.Printf("%c \n", 34)
	fmt.Printf("%x \n", 15)
	fmt.Printf(" '5 mod 2' = %d \n", 5%2)
	fmt.Printf(" '4 mod  2' = %d \n", 4%2)
}

func SimpleHttpServer() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go on Now!</h1>")
}

func TestEndPoint(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "testing server - server is running !!!")
}

func StartRestApiServices(){
	fmt.Println("Strating application")
	log.Print("Starting Go Server at http://localhost:8011")
	router := mux.NewRouter()

	router.HandleFunc("/test", TestEndPoint).Methods("GET")
	//router.HandleFunc("/projects", GetProjectsDataEndPoint).Methods("GET")
	//router.HandleFunc("/project/{projectid}", GetProjectEndPoint).Methods("GET")
	//router.HandleFunc("/project/{projectid}", CreateProjectEndPoint).Methods("POST")
	//router.HandleFunc("/project/{projectid}", DeleteProjectEndPoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8011", router))
}