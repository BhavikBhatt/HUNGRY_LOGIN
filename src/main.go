package main


import (
    "fmt"
    "log"
    "net/http"
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
	"services"
	"strconv"
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

    username := r.FormValue("username")
    password := r.FormValue("password")

	authenticated := services.Authenticate(ctx, usersCollection, username, password)

	if len(authenticated.Username) > 0 {
		fmt.Fprintf(w, "Welcome %s!\n", authenticated.Name)
		fmt.Fprintf(w, "Check logs file for user and authentication information.")
	} else {
		fmt.Fprintf(w, "Wrong username and/or password.\n")
		fmt.Fprintf(w, "Check logs file for user and authentication information.")
	}
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {

    if r.URL.Path != "/createUser" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err)
        return
	}
	
    username := r.FormValue("username")
	password := r.FormValue("password")
	name := r.FormValue("name")
	age := r.FormValue("age")
	email := r.FormValue("email")

	ciphertext := services.Encrypt([]byte(password), "so hungry")
	intage, _ := strconv.Atoi(age)
    newUser := MongoFields{
        Username: username,
		Ciphertext: string(ciphertext),
		Email: email,
		Name: name,
		Age: intage,
	}
	
	res, err := usersCollection.InsertOne(ctx, newUser)
	if err != nil {
		fmt.Println("InsertOne ERROR:", err)
	}
	_ = res

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
    fileServer := http.FileServer(http.Dir("./site"))
    http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/createUser", createUserHandler)

    fmt.Printf("Starting server at port 8080\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}