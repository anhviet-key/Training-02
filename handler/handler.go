package handler

import (
	"encoding/json"
	db "go-main/package/data"
	m "go-main/package/todo"
	"html/template"
	"strconv"

	"net/http"

	"github.com/gorilla/mux"
)

func GetAllTodo(writer http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/html/layout.html"))
	datas := m.Todos
	tmpl.Execute(writer, datas)
}

func GetApiTodo(writer http.ResponseWriter, request *http.Request) {

	var newtodos []m.Todo
	var todo m.Todo
	vals, err := db.RedisClient.LRange("Items", 0, -1).Result()

	if err != nil {
		panic(err)
	}

	for _, value := range vals {
		json.Unmarshal([]byte(value), &todo)
		newtodos = append(newtodos, todo)

	}
	responseWithJson(writer, http.StatusOK, newtodos)
}

func CreateTodo(writer http.ResponseWriter, request *http.Request) {
	var newTodo m.Todo
	if err := json.NewDecoder(request.Body).Decode(&newTodo); err != nil {
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}
	newTodo.ID = generateId(m.Todos)

	m.Todos = append(m.Todos, newTodo)

	bytes, _ := json.Marshal(newTodo)

	_, err := db.RedisClient.RPush("Items", bytes).Result()

	if err != nil {
		panic(err)
	}

	responseWithJson(writer, http.StatusCreated, newTodo)
}

func UpdateTodo(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid todo id"})
		return
	}

	var updateTodo m.Todo
	if err := json.NewDecoder(request.Body).Decode(&updateTodo); err != nil {
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}

	getId := db.RedisClient.LIndex("Items", int64(id)).Val()

	sec := map[string]interface{}{}
	if err := json.Unmarshal([]byte(getId), &sec); err != nil {
		panic(err)
	}

	sec["id"] = int(sec["id"].(float64))
	sec["title"] = updateTodo.Title
	sec["status"] = updateTodo.Status
	bytes, _ := json.Marshal(sec)

	_, err = db.RedisClient.LSet("Items", int64(id), bytes).Result()

	if err != nil {
		panic(err)
	}
	responseWithJson(writer, http.StatusOK, map[string]string{"message": "Todo was update"})
}

func DeleteTodo(writer http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid todo id"})
		return
	}
	bytes := db.RedisClient.LIndex("Items", int64(id)).Val()

	_, err = db.RedisClient.LRem("Items", 1, bytes).Result()
	if err != nil {
		panic(err)
	}
	responseWithJson(writer, http.StatusOK, map[string]string{"message": "Todo was deleted"})
}

func responseWithJson(writer http.ResponseWriter, status int, object interface{}) {
	// writer.Header().Set("Content-Type", "application/json")
	// writer.Header().Set("Access-Control-Allow-Origin", "*")
	// writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	// writer.Header().Set("Access-Control-Allow-Headers", "Authentication ,Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	writer.WriteHeader(status)
	json.NewEncoder(writer).Encode(object)
}

func generateId(todos []m.Todo) int {

	var maxId int
	for _, todo := range todos {
		if todo.ID > maxId {
			maxId = todo.ID
		}
	}
	return maxId + 1
}
