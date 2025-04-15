# Project from https://roadmap.sh/projects/task-tracker

Run `$ go build` to build the project or download the .exe file from releases page

`$ task-tracker-cli list` - Listing all tasks  
`$ task-tracker-cli list [todo, in-progress, done]` - Listing tasks by status  

`$ task-tracker-cli add "Buy groceries"` - Adding a new task  
`$ task-tracker-cli update 1 "Feed the cat"` - Updating a task  
`$ task-tracker-cli delete 1` - Deleting a task  

`$ task-tracker-cli mark-in-progress 1` - Marking a task as in progress  
`$ task-tracker-cli mark-done 1` - Marking a task as done  

All tasks are outputted in the format of: `- [id] description - status (createdAt | updatedAt)` in your time zone  
Example: `- [5] Write README.md - todo (Apr 7, 2025 23:21 | Updated: Apr 8, 2025 18:53)`  

