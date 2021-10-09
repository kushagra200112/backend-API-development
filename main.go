package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type user struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type post struct{
	ID string `json:"id"`
	Caption string `json:"caption"`
	Imageurl string `json:"imageurl"`
	Timestamp string `json"timestamp"`
}

var users []user
var post [] Post

func Createuser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json") // to present in the json format and not in the text format
	json.NewEncoder(w).Encode(users)
}
func Getuser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //get paramaters
	for _, items := range users{
		if items.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&user{})
}
func postid(w http.ResponseWriter, r *http.Request){
	for _, i := range post{
		if i.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Post{})
}
func postsuser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json") 
	json.NewEncoder(w).Encode(post)
	json.NewEncoder(w).Encode(users)
}

func main() {
	//init router
	r:= mux.NewRouter()
	
	// mock data
	
	users = append(users, user{ID: "1", Name: "john doe", Email: "johndoe@xyzmail.com", Password: "johndoe!123"})
	users = append(users, user{ID: "2", Name: "steve smith", Email: "stevesmith@xyzmail.com", Password: "**stevesmith123**"}
	users = append(users, user{ID: "3", Name: "harrison wells", Email: "harrisonwells@xyzmail.com", Password: "harrison12wells3"}
	users = append(users, user{ID: "4", Name: "barry allen", Email: "barryallen@xyzmail.com", Password: "barry123allen"}
	
	post = append(post, Post{ID: "1", Caption: "hello", Imageurl: "xyz.com", Timestamp: time.Now()})
	post = append(post, Post{ID: "2", Caption: "hello world", Imageurl: "xyz.com", Timestamp: time.Now()})
	post = append(post, Post{ID: "3", Caption: "good day", Imageurl: "xyz.com", Timestamp: time.Now()})
	post = append(post, Post{ID: "4", Caption: "work hard", Imageurl: "xyz.com", Timestamp: time.Now()})
	//route handlers

	r.HandleFunc("/users", Createuser).Methods("POST")
	r.HandleFunc("/users/{id}", Getuser).Methods("GET")
	r.HandleFunc("/posts", posts).Methods("POST")
	r.HandleFunc("/posts/{id}", postid).Methods("GET")
	r.HandleFunc("/posts/users/{id}", postsuser).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
func connect(){
	const uri = "mongodb+srv://<username>:<password>@cluster0.fboyk.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")
}
}
