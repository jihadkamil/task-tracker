## Task-Tracker

** Features **
- tasks management
- add new task
- update task
- update task status
- delete task
- view all tasks list
- view all tasks list filtered by status

** How to Run The Program **
> git clone https://github.com/jihadkamil/task-tracker.git
> cd task-tracker
<!-- 1. run on the fly -->
> go run main.go
<!-- 2. run on the executable file -->
> go build -o task-tracker
> ./task-tracker

* Add new task: automatically create id, status (todo) createdAt, updatedAt
> add "task-name"

* Get all tasks list
> list

* Update task name by id
> update 1 "task-name-2"

* Delete task by id
> delete 1

* Update task status to done
> mark-done 1

* Update task status to in-progress
> mark-in-progress 1


* Get tasks list filtered by status
> list todo
> list in-progress
> list done

* Exit program
> exit

* Clear terminal
> clear


@author: jihadkamil.dev
https://roadmap.sh/projects/task-tracker
