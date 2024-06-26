package handlers

import (
    "context"
    "encoding/json"
    
    "net/http"
    "time"
    
    
    "github.com/imkarannn/todo-go/db"
    "github.com/imkarannn/todo-go/models"
    "go.mongodb.org/mongo-driver/bson"

)



// Handler to create a new user
//This function is the endpoint for creating a new user. It receives an HTTP request (r) and a response writer (w) as parameters.
func CreateUserEndpoint(w http.ResponseWriter, r *http.Request) {
    
    w.Header().Set("Content-Type", "application/json")
    //This line sets the response header to indicate that the response will be in JSON format.
    var user models.UserMaster
    
    //decoding the json data 
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    user.CreatedAt = time.Now()
    
    /*This line gets the MongoDB collection where you want to insert the user data. 
    GetCollection() is likely a function defined elsewhere in your code that returns the MongoDB collection.*/
    collection := db.GetCollection()
    

    /*This line inserts the user struct into the MongoDB collection.
     If there's an error inserting, it returns a Internal Server Error response with the error message.*/
    result, err := collection.InsertOne(context.Background(), user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }


    /*This line encodes the result  into JSON format and writes it to the response writer w.*/
    json.NewEncoder(w).Encode(result)
}

func AddUserRole(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var userRole models.UserRoleMaster
    if err := json.NewDecoder(r.Body).Decode(&userRole); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    collection :=db.GetCollection()
    result, err := collection.InsertOne(context.Background(), userRole)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(result)
}

// Handler to get all users
func GetUsersEndpoint(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var users []UserMaster
    collection := db.GetCollection()
    cursor, err := collection.Find(context.Background(), bson.D{})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer cursor.Close(context.Background())
    for cursor.Next(context.Background()) {
        var user models.UserMaster
        if err := cursor.Decode(&user); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        users = append(users, user)
    }
    if err := cursor.Err(); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(users)
}
