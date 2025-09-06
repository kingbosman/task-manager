### Taskmaster: A CLI To-Do List

This is a simple command-line interface (CLI) to-do list application built to explore the **Cobra** framework in **Go**. It allows users to manage their daily tasks directly from the terminal.

---

### Key Features

* **Task Management**: Add, list, and complete tasks.
* **Simple Interface**: Interact with your to-do list using intuitive commands.
* **Cobra Integration**: A practical demonstration of how to structure a CLI application with commands, subcommands, and flags using the Cobra library.
* **Persistent Storage**: Tasks are saved to a file, so they persist between sessions.

---

### How to Use

1.  **Installation**
    ```sh
    go install github.com/kingbosman/task-manager
    ```

2.  **Adding a Task**
    ```sh
    task add "push to prod"
    ```

3.  **Listing Tasks**
    ```sh
    task list --page=2
    ```

4.  **Completing a Task**
    ```sh
    task complete [task_id]
    ```
    
4.  **Deleting a Task**
    ```sh
    task delete [task_id]
    ```

---

### Built With

* **Go**: The programming language
* **Cobra**: A powerful framework for building modern CLI applications
* **SQLite3**: Used for simple, file-based database storage

Feel free to explore the code to see how Cobra is used to build a robust and user-friendly CLI.
