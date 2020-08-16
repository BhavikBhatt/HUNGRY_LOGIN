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
	"time"
	"os"
	"models"
)

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
	
	f, err := os.OpenFile("./logs/logins.log",
	os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	logger := log.New(f, "prefix", log.LstdFlags)

    username := r.FormValue("username")
    password := r.FormValue("password")

	authenticated := services.Authenticate(ctx, usersCollection, username, password)

	if len(authenticated.Username) > 0 {
		fmt.Fprintf(w, "Welcome %s! You are authenticated!\n", authenticated.Name)
		logger.Println("Successful login by ", username)

	} else {
		fmt.Fprintf(w, "Wrong username and/or password.\n")
		logger.Println("Unsuccessful login attempt by ", username)
	}
	fmt.Fprintf(w, "Check logs file for user and authentication information.\n")
	fmt.Fprintf(w, "Press back to go back to the login page.")

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
    newUser := models.User{
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

	f, err := os.OpenFile("./logs/new_users.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, "prefix", log.LstdFlags)
	logger.Println("New User added at ", time.Now(), res.InsertedID,  newUser.Username, newUser.Name, newUser.Email)

    f, err2 := os.Open("site/index.html")
    if err2 != nil {
		// handle error
        return
	}
	time.Sleep(4 * time.Second)
    http.ServeContent(w, r, "site/index.html", time.Now(), f)

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