package todo

import(
	"net/http"
	"github.com/go-chi/go-chi"
	"github.com/go-chi/render"
)

type Todo struct {
	Slug string `json: "slug"`
	Title string `json: "title"`
	Body string `json: "body"`
}

func Routes() *chi.Mux {
	router := chi.Newrouter()
	router.Get("/{todoId}", GetATodo)
	router.Delete("/{todoId}", DeleteTodo)
	router.Post("/", CreateTodo)
	router.Get("/", GetAllTodos)
	return router
}


func GetATodo(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	todos := Todo{
		Slug: todoID,
		Title: "Hello world",
		Body: "Heloo world from planet earth",
	}
	render.JSON(w,r,todos)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request){
	response := make(map[string]string)
	response["message"] = "Deleted TODO success"
	response.JSON(w,r, response)
}

func CreateTodo( w http.ResponseWriter, r *http.Request){
	todos := []Todo{
		{
			Slug: "slug",
			Title: "Hello World",
			Body: "Hello world from planet earth"
		},
	}
	render.JSON(w,r,todos)
}

