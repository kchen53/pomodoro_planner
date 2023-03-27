# Backend Server
## How to run server:

### Navigate to the server exe file location

Normally located in ./src/server

### Run .exe

Run "pomodoro_planner_FE.exe"

### Server should be running!

If the terminal shows no output and only has a blank line, the the server is working properly

### *Stop server

To stop the server just press Ctrl+C in the terminal

## How to build project

#### Prerequisites:

* Installed gcc
* (If working in VS code) install C/C++ extension

### Set CGO_ENABLED

At the beginning of every work session this environment variable must be set <br>
Type "set CGO_ENABLED=1" into terminal

### Build

If testing build as normal with "go build" <br> <br>
To push changes to the front end, build to the front end server exe: pomodoro_planner_FE.exe <br>
Type "go build -o pomodoro_planner_FE.exe
