# HUNGRY_LOGIN

## My first go at Go!

This is a login webapp created with a Mongo-backed Go web server. Authentication is supported by AES encryption/decryption. User objects are stored in a MongoDB Atlas database. An example of the current database is shown below in a screenshot. Logging is also implemented - so login attempts and new account creation events are recorded. The instructions to start the server are below, along with a few accounts that are already created. 

### Local Workspace Setup
1. Clone this repo into your local workspace
2. Set the GOPATH environment variable using `export GOPATH=<this project's path>` in your terminal
3. From the `src` folder, run `go run main.go` to start the server!

### Trying it out!
1. After starting the server, navigate to `localhost:8080` in any browser
2. Log in with one of the users displayed below
3. You'll see the user details displayed on the screen
4. Check the logins.log file in your local workspace to see the event logged

### Create a new user
1. Click the `New User?` button on the main login page
2. Fill out the new user form and click `Sign me up!` (should redirect to login page)
3. Try logging in with the new account
4. You can also check the new.log file to see the details of the added user



#### User 1
  * username: bhavikbhatt
  * password: bhattbhavik23
  
#### User 2
  * username: jayz
  * password: zjay50

#### User 3
  * username: kevinhart
  * password: hartkevin41


![User model shown in MongoDB Atlas UI](https://github.com/BhavikBhatt/HUNGRY_LOGIN/blob/master/pkg/Screen%20Shot%202020-08-15%20at%2010.30.52%20PM.png)
