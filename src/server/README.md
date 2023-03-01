# Backend Server
## How to run server:
### Choose Angular Handler 
Backend: Use angular_embedded.go <br>
Frontend: Use angular_live.go <br>
One of the two above files will be located in the src folder (outside the server folder) and the other inside the server folder. Drag the desired file into the server folder and the other into the src folder.  <br>
Basically you should only have one of the two files inside the server folder at a time <br>

### Navigate terminal to correct folder
Navigate the terminal to .\src\server <br>

### Load dependencies
If it has been a while or building is not working, dependencies may have been added that you don't have
To fix this, run "go mod tidy"

### Build (First time running after a pull or after every server code update)
Run "go build" in command line (make sure you are in the server folder) <br>

### Run .exe
Run "pomodoro_planner.exe" <br>

### Server should be running!
If the terminal shows no output and only has a blank line, the the server is working properly

### *Close server
To close the server just press Ctrl+C in the terminal
