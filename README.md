# f1-telemetry
Telemetry app for EA/Codemasters [F1 games](https://www.ea.com/games/f1). The telemetry receiver is written in Go, with web based data visualization.
This is a fun project with no roadmap for future features.

## Run application
`go run main.go`

## Build application
`go build main.go`

# Web Interface
While the application is running, the web interface can be accessed through http://localhost:8080/
The IP and port number is displayed at the end of the page. These values should be used in F1 21+ Telemetry settings if the game is run in the same network.