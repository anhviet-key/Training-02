package utils

import (
	"encoding/json"
	db "go-main/package/data"
	m "go-main/package/todo"

	"github.com/go-redis/redis"
)

func GetTodos() (todos []m.Todo) {
	keys, err := db.RedisClient.Keys("*").Result()

	if err != nil {
		panic(err)
	}

	for _, key := range keys {
		todo := GetTodo(key)
		todos = append(todos, todo)
	}
	return todos
}

func GetTodo(key string) (todo m.Todo) {
	val, err := db.RedisClient.Get(key).Result()

	if err != nil {
		panic(err)
	}

	if err != redis.Nil {
		err = json.Unmarshal([]byte(val), &todo)
	}

	return todo
}
