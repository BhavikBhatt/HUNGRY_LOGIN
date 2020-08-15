package main


import (
    "fmt"
    "log"
    "net/http"
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
	"services"
	"reflect"
)


type MongoFields struct {
    Username string `json:"Field Str"`
	Ciphertext string `json:"Field Str"`
	Email string `json:"Field Str"`
	Name string `json:"Field Str"`
	Age int `json:"Field Int"`
}

var ctx context.Context
var usersCollection *mongo.Collection

func formHandler(w http.ResponseWriter, r *http.Request) {

    if r.URL.Path != "/form" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err)
        return
    }
    //fmt.Fprintln(w, "POST request successful")
    username := r.FormValue("username")
    password := r.FormValue("password")

	authenticated := services.Authenticate(ctx, usersCollection, username, password)

	if authenticated {
		//fmt.Fprintf(w, "AUTHENTICATED! :)")
		http.Redirect(w, r, "google.com", http.StatusSeeOther)
	} else {
		fmt.Fprintf(w, "WRONG PASSWORD :(")
	}

}

func main() {

    client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:admin@hungrycluster.pebho.mongodb.net/HungryUsers?retryWrites=true&w=majority"))
    if err != nil {
        log.Fatal(err)
    }
    ctx = context.Background()
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }
    //defer client.Disconnect(ctx)

    hungryUsersDatabase := client.Database("HungryUsers")
    usersCollection = hungryUsersDatabase.Collection("users")

    _ = hungryUsersDatabase
    _ = usersCollection
	fmt.Println(reflect.TypeOf(ctx))
    fileServer := http.FileServer(http.Dir("./site"))
    http.Handle("/", fileServer)
    http.HandleFunc("/form", formHandler)

    fmt.Printf("Starting server at port 8080\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}