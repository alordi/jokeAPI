package main

import (
	"context"
    "fmt"
    "log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
	"strings"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
)

var gmLambda *gorillamux.GorillaMuxAdapter

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Welcome to Austin's Joke API")
	fmt.Fprintf(w, "For more info go to: https://github.com/alordi/jokeAPI")
    fmt.Println("Endpoint Hit: homePage")
}


func getJokes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var jokes []Joke

	//Connection mongoDB with helper class
	collection := ConnectDB()

	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		GetError(err, w)
		return
	}

	// Close the cursor once finished
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var joke Joke
		err := cur.Decode(&joke)
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		jokes = append(jokes, joke)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(jokes)
}

func getJokesByType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var jokes []Joke
	var params = mux.Vars(r)
	temp := params["types"]
	types := strings.Split(temp, ",")
	
	collection := ConnectDB()

	match := bson.D{
		{"$match", bson.D{
			{"type", bson.D{
				{"$in", types},
			}},
		}},
	}

	cur, err := collection.Aggregate(context.TODO(), mongo.Pipeline{match})

	if err != nil {
		GetError(err, w)
		return
	}

	// Close the cursor once finished
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		var joke Joke
		err := cur.Decode(&joke)
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		jokes = append(jokes, joke)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(jokes)
}

func getJokesNotInType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var jokes []Joke
	var params = mux.Vars(r)
	temp := params["types"]
	types := strings.Split(temp, ",")
	
	collection := ConnectDB()

	match := bson.D{
		{"$match", bson.D{
			{"type", bson.D{
				{"$nin", types},
			}},
		}},
	}

	cur, err := collection.Aggregate(context.TODO(), mongo.Pipeline{match})

	if err != nil {
		GetError(err, w)
		return
	}

	// Close the cursor once finished
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		var joke Joke
		err := cur.Decode(&joke)
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		jokes = append(jokes, joke)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(jokes)
}

func getRandomJoke(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var joke Joke

	collection := ConnectDB()

	sample := bson.D{
		{"$sample", bson.D{
			{"size", 1},
		}},
	}

	cur, err := collection.Aggregate(context.TODO(), mongo.Pipeline{sample})

	if err != nil {
		GetError(err, w)
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		er := cur.Decode(&joke) 
		if er != nil {
			log.Fatal(er)
		}
	}
		
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(joke) 
}

func getRandomJokeByType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var joke Joke

	var params = mux.Vars(r)
	temp := params["types"]
	types := strings.Split(temp, ",")

	//Connection mongoDB with helper class
	collection := ConnectDB()

	match := bson.D{
		{"$match", bson.D{
			{"type", bson.D{
				{"$in", types},
			}},
		}},
	}

	sample := bson.D{
		{"$sample", bson.D{
			{"size", 1},
		}},
	}

	cur, err := collection.Aggregate(context.TODO(), mongo.Pipeline{match,sample})

	if err != nil {
		GetError(err, w)
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		er := cur.Decode(&joke)
		if er != nil {
			log.Fatal(er)
		}
	}
		

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(joke)
}

func getRandomJokeNotInType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var joke Joke

	var params = mux.Vars(r)
	temp := params["types"]
	types := strings.Split(temp, ",")

	//Connection mongoDB with helper class
	collection := ConnectDB()

	match := bson.D{
		{"$match", bson.D{
			{"type", bson.D{
				{"$nin", types},
			}},
		}},
	}

	sample := bson.D{
		{"$sample", bson.D{
			{"size", 1},
		}},
	}

	cur, err := collection.Aggregate(context.TODO(), mongo.Pipeline{match,sample})

	if err != nil {
		GetError(err, w)
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		er := cur.Decode(&joke)
		if er != nil {
			log.Fatal(er)
		}
	}
		

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(joke)
}

func getJokeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var joke Joke
	var params = mux.Vars(r)
	id, e := strconv.Atoi(params["id"])

	collection := ConnectDB()

	filter := bson.M{"jokeId": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&joke)

	if err != nil || e != nil {
		GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(joke)
}


func init() {
	//Init Router
	r := mux.NewRouter()

	// arrange our routes
	r.HandleFunc("/", homePage).Methods("GET")
	r.HandleFunc("/jokes", getJokes).Methods("GET")
	r.HandleFunc("/jokes/type={types}", getJokesByType).Methods("GET")
	r.HandleFunc("/jokes/type!={types}", getJokesNotInType).Methods("GET")
	r.HandleFunc("/jokes/random", getRandomJoke).Methods("GET")
	r.HandleFunc("/jokes/random/type={types}", getRandomJokeByType).Methods("GET")
	r.HandleFunc("/jokes/random/type!={types}", getRandomJokeNotInType).Methods("GET")
	r.HandleFunc("/jokes/{id}", getJokeByID).Methods("GET")
	

  	// set our port address
	//log.Fatal(http.ListenAndServe(":8000", r))

	gmLambda = gorillamux.New(r)

}

func HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return gmLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(HandleRequest)
}