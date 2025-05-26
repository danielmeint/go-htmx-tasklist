# Simple Go Web Application

This is a simple interactive web application built with Go, htmx, and Bootstrap. It serves as a to-do list application for learning purposes. The application allows you to create tasks, mark tasks as completed, delete tasks, and clear all tasks.

## Features

- Create new tasks.
- Mark tasks as completed or incomplete.
- Delete individual tasks.
- Clear all tasks at once.

## Technologies Used

- Go (Golang) for the backend.
- [htmx](https://htmx.org/) for dynamic HTML updates.
- [Bootstrap](https://getbootstrap.com/) for styling.
- SQLite for the database.

## Installation and Usage

1. Clone this repository:

   ```bash
   git clone https://github.com/danielmeint/go-htmx-tasklist
   ```

2. Navigate to the project directory:

   ```bash
   cd go-htmx-tasklist
   ```

3. Install Go dependencies:

   ```bash
   go get
   ```

4. Run the application:

   ```bash
   go run main.go
   ```

5. Access the application in your web browser at [http://localhost:8080](http://localhost:8080).

## Troubleshooting

**Database connection error?** The SQLite driver requires CGO to be enabled:
```bash
go env -w CGO_ENABLED=1
```

If you're missing build tools:
- Ubuntu/Debian: `sudo apt-get install build-essential`
- macOS: `xcode-select --install`

The database file (`tasklist.db`) is created automatically on first run.

## Usage

- Create a new task: Enter a task name in the input field and click "Add Task."
- Mark a task as completed: Click the checkbox next to the task.
- Delete a task: Click the "Delete" button next to the task.
- Clear all tasks: Click the "Clear All Tasks" button.

## Contributing

Contributions are welcome! If you'd like to contribute to this project, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them.
4. Push your changes to your fork.
5. Create a pull request to the original repository.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
