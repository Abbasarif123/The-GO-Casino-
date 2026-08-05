# 🎰 Go Casino - Web Slot Machine

A full-stack web-based slot machine game. This project features a robust backend built in Go and a modern, responsive frontend built with vanilla HTML, CSS, and JavaScript.

It was transitioned from a terminal-based Go application into a stateless web server, utilizing a RESTful API to handle game logic, bets, and balances.

Since Github pages cannot run Go code. Please follow the steps to run this project below.

## Features

* **Full-Stack Architecture:** Go handles the math, state validation, and API routing, while the browser renders the UI.
* **Stateless API:** Game spins and balance calculations are processed securely on the backend via a `/api/spin` endpoint.


* **Dynamic Win Calculation:** Payouts are calculated based on symbol frequencies and multipliers.


* **Game Over & Quit Logic:** The game locks the UI gracefully when the user bets `$0` or runs out of money.
* **Cloud-Ready:** Designed to easily deploy to platforms like Render or Railway using dynamic port mapping.



## 📂 Project Structure

```text
go-casino/
├── go.mod                 # Go module file for dependencies and cloud compilation
├── main.go                # Server entry point, winning logic, and global configurations
├── api.go                 # HTTP handlers and JSON struct definitions
├── spin.go                # Core slot machine mathematics and grid generation
└── static/                
    └── index.html         # The frontend UI, styles, and client-side JavaScript

```

* `main.go`: Initializes the game variables, calculates winning combinations across the grid, and starts the HTTP server.


* `api.go`: Replaces traditional terminal inputs with web structs (`SpinRequest` & `SpinResponse`) to parse incoming JSON bets and return the spin results.


* `spin.go`: Contains the mathematical engine that flattens the probability map and generates a random 3x3 2D slice.



## How to Run Locally

### Prerequisites

* You must have [Go installed](https://www.google.com/search?q=https://go.dev/doc/install) on your machine.

### Installation & Execution

1. Clone this repository or download the files into a single directory.
2. Open your terminal and navigate to the project folder.
3. If you haven't initialized the module yet, run:
```bash
go mod init go-casino

```


4. Run the Go server:
```bash
go run .

```


5. Open your web browser and navigate to: **http://localhost:8080**

## How to Deploy (Render)

This project is configured to read the `PORT` environment variable, making it immediately compatible with modern cloud hosts.

1. Push your code to a GitHub repository.
2. Create a free account on [Render](https://www.google.com/search?q=https://render.com).
3. Click **New +** -> **Web Service** -> **Build and deploy from a Git repository**.
4. Connect your repository.
5. Use the following deployment settings:
* **Runtime:** Go
* **Build Command:** `go build -o app .`
* **Start Command:** `./app`


6. Deploy! Your app will be live on an `.onrender.com` domain.

## How to Play

1. Enter your desired bet amount in the input field.
2. Click **SPIN!**
3. If a horizontal line matches completely, you win based on that symbol's multiplier!


4. **To quit:** Enter `0` in the bet field and hit spin.
5. **Game Over:** If your balance drops to `$0`, the slot machine will lock.
