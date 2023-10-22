package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var db *gorm.DB

// Task represents a task in the to-do list
type Task struct {
	gorm.Model
	Name      string
	Completed bool
}

func main() {
	// Initialize the database connection
	var err error
	db, err = gorm.Open("sqlite3", "tasklist.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	defer db.Close()

	// AutoMigrate the Task model (create the table)
	db.AutoMigrate(&Task{})

	// Initialize the router
	router := mux.NewRouter()

	// Serve the index.html file under the root path
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/index.html")
	}).Methods("GET")

	router.HandleFunc("/tasks", ListTasks).Methods("GET")
	router.HandleFunc("/create", CreateTask).Methods("POST")
	router.HandleFunc("/update/{id:[0-9]+}", UpdateTask).Methods("POST")
	router.HandleFunc("/delete/{id:[0-9]+}", DeleteTask).Methods("POST")
	router.HandleFunc("/clear-all", ClearAllTasks).Methods("POST")

	// Serve static files (Tailwind CSS and JS)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Log that we're starting the server
	fmt.Println("Starting server on port 8080")

	// Start the server
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)

}

// RenderTaskList renders the task list using the provided tasks and writes it to the response writer
func RenderTaskList(w http.ResponseWriter, tasks []Task) {
	// Define a template for rendering just the task list
	tmpl := template.Must(template.New("task-list").Parse(`
        {{range .}}
        <div class="task flex justify-between items-center py-6 px-4">
            <!-- Use hx-post to specify the URL for toggling the task completion status -->
            <input
                type="checkbox"
                {{if .Completed}}checked{{end}}
                hx-post="/update/{{.ID}}"
                hx-target="#task-list"
				class="form-checkbox h-4 w-4 text-indigo-600 transition duration-150 ease-in-out"
            />
            {{.Name}}
			<button hx-post="/delete/{{.ID}}" hx-target="#task-list"
			type="button" class="btn btn-close" aria-label="Close"></button>
        </div>
        {{end}}
    `))

	// Execute the template, passing the provided tasks data
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.Execute(w, tasks); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// ListTasks lists all tasks
func ListTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ListTasks")
	var tasks []Task
	db.Find(&tasks)
	RenderTaskList(w, tasks)
}

// CreateTask creates a new task
func CreateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateTask" + r.Form.Get("taskName"))

	// Parse the form data from the request
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Extract the task name from the form data
	taskName := r.Form.Get("taskName")

	// Create a new task
	newTask := Task{
		Name:      taskName,
		Completed: false, // Initial status is incomplete
	}

	// Save the new task to the database
	db.Create(&newTask)

	// Retrieve the updated list of tasks from the database
	var tasks []Task
	db.Find(&tasks)
	RenderTaskList(w, tasks)
}

// UpdateTask marks a task as completed or incomplete
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	// Get the task ID from the URL parameter
	vars := mux.Vars(r)
	taskID := vars["id"]

	fmt.Println("UpdateTask" + taskID)

	// Find the task in the database by its ID
	var task Task
	if err := db.First(&task, taskID).Error; err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	// Toggle the completion status of the task
	task.Completed = !task.Completed
	db.Save(&task)

	// Retrieve the updated list of tasks from the database
	var tasks []Task
	db.Find(&tasks)
	RenderTaskList(w, tasks)
}

// DeleteTask deletes a task
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	// Delete the task from the database
	vars := mux.Vars(r)
	taskID := vars["id"]

	fmt.Println("DeleteTask" + taskID)

	var task Task
	if err := db.First(&task, taskID).Error; err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	db.Delete(&task)

	// Retrieve the updated list of tasks from the database
	var tasks []Task
	db.Find(&tasks)
	RenderTaskList(w, tasks)
}

// ClearAllTasks clears all tasks
func ClearAllTasks(w http.ResponseWriter, r *http.Request) {
	// Delete all tasks from the database
	db.Delete(&Task{})

	// Retrieve the updated list of tasks from the database
	var tasks []Task
	db.Find(&tasks)
	RenderTaskList(w, tasks)
}
