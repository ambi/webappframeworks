package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"runtime"
	"sort"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Message struct {
	Message string `json:"message"`
}

type World struct {
	ID           uint16 `json:"id"`
	RandomNumber uint16 `json:"randomNumber"`
}

type Fortune struct {
	ID      uint16 `json:"id"`
	Message string `json:"message"`
}

type Fortunes []*Fortune

func (f Fortunes) Len() int {
	return len(f)
}

func (f Fortunes) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

type FortunesByMessage struct {
	Fortunes
}

func (f FortunesByMessage) Less(i, j int) bool {
	return f.Fortunes[i].Message < f.Fortunes[j].Message
}

const (
	// Template
	fortuneHTML = `
    <!doctype html>
    <html>
    <head>
      <title>Fortunes</title>
    </head>
    <body>
      <table>
        <tr>
          <th>id</th>
          <th>message</th>
          </tr>
        {{range .}}
        <tr>
          <td>{{ .ID }}</td>
          <td>{{ .Message }}</td>
        </tr>
        {{end}}
      </table>
    </body>
    </html>
  `
	// Database
	connectionString = "root@tcp(%s)/hello_world"
	worldSelect      = "SELECT id, randomNumber FROM World WHERE id = ?"
	worldUpdate      = "UPDATE World SET randomNumber = ? WHERE id = ?"
	fortuneSelect    = "SELECT id, message FROM Fortune"
	worldRowCount    = 10000
	maxConnections   = 256
	// HTTP
	HeaderContentType   = "Content-Type"
	MIMEApplicationJSON = "application/json"
	MIMETextHTML        = "text/html"
	MIMETextPlain       = "text/plain"
)

var (
	// Database
	db                *sql.DB
	worldSelectStmt   *sql.Stmt
	worldUpdateStmt   *sql.Stmt
	fortuneSelectStmt *sql.Stmt

	// Template
	fortuneTmpl = template.Must(template.New("fortune").Parse(fortuneHTML))

	helloWorld = []byte("Hello, World!")
)

// Test 1: JSON serialization
func jsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(HeaderContentType, MIMEApplicationJSON)
	json.NewEncoder(w).Encode(Message{"Hello, World!"})
}

// Test 2: Single database query
func dbHandler(w http.ResponseWriter, r *http.Request) {
	world := new(World)
	if err := fetchRandomWorld(world); err != nil {
		w.WriteHeader(500)
		return
	}

	w.Header().Set(HeaderContentType, MIMEApplicationJSON)
	json.NewEncoder(w).Encode(world)
}

// Test 3: Multiple database queries
func queriesHandler(w http.ResponseWriter, r *http.Request) {
	n := getQueryCount(r.URL.Query().Get("n"))
	worlds := make([]World, n)
	for i := 0; i < n; i++ {
		if err := fetchRandomWorld(&worlds[i]); err != nil {
			w.WriteHeader(500)
			return
		}
	}

	w.Header().Set(HeaderContentType, MIMEApplicationJSON)
	json.NewEncoder(w).Encode(worlds)
}

// Test 4: Fortunes
func fortunesHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := fortuneSelectStmt.Query()
	if err != nil {
		w.WriteHeader(500)
		return
	}
	defer rows.Close()

	fortunes, err := fetchFortunes(rows)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	fortunes = append(fortunes, &Fortune{Message: "Additional fortune added at request time."})
	sort.Sort(FortunesByMessage{fortunes})

	w.Header().Set(HeaderContentType, MIMETextHTML)
	fortuneTmpl.Execute(w, fortunes)
}

// Test 5: Database updates
func updatesHandler(w http.ResponseWriter, r *http.Request) {
	n := getQueryCount(r.URL.Query().Get("n"))
	worlds := make([]World, n)
	for i := 0; i < n; i++ {
		// Fetch and modify
		world := &worlds[i]
		if err := fetchRandomWorld(&worlds[i]); err != nil {
			w.WriteHeader(500)
			return
		}
		world.RandomNumber = uint16(randomWorldNum())

		// Update
		if _, err := worldUpdateStmt.Exec(world.RandomNumber, world.ID); err != nil {
			w.WriteHeader(500)
			return
		}
	}

	w.Header().Set(HeaderContentType, MIMEApplicationJSON)
	json.NewEncoder(w).Encode(worlds)
}

// Test 6: Plaintext
func plaintextHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(HeaderContentType, MIMETextPlain)
	w.Write(helloWorld)
}

func fetchRandomWorld(w *World) error {
	n := randomWorldNum()
	return worldSelectStmt.QueryRow(n).Scan(&w.ID, &w.RandomNumber)
}

func randomWorldNum() int {
	return rand.Intn(worldRowCount) + 1
}

func getQueryCount(q string) int {
	n, err := strconv.Atoi(q)
	if err != nil || n < 1 {
		return 1
	}
	if n > 500 {
		return 500
	}
	return n
}

func fetchFortunes(rows *sql.Rows) (Fortunes, error) {
	fortunes := make(Fortunes, 0)
	for rows.Next() { // Fetch rows
		f := new(Fortune)
		if err := rows.Scan(&f.ID, &f.Message); err != nil {
			return nil, fmt.Errorf("Error scanning fortune row: %s", err.Error())
		}
		fortunes = append(fortunes, f)
	}
	return fortunes, nil
}

func initDB() {
	host := "localhost"

	var err error
	db, err = sql.Open("mysql", fmt.Sprintf(connectionString, host))
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	db.SetMaxIdleConns(maxConnections)
	db.SetMaxOpenConns(maxConnections)

	worldSelectStmt, err = db.Prepare(worldSelect)
	if err != nil {
		log.Fatal(err)
	}
	worldUpdateStmt, err = db.Prepare(worldUpdate)
	if err != nil {
		log.Fatal(err)
	}
	fortuneSelectStmt, err = db.Prepare(fortuneSelect)
	if err != nil {
		log.Fatal(err)
	}
}

type apiHandler struct{}

func (apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "Go net/http")
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	switch r.URL.Path {
	case "/json":
		jsonHandler(w, r)
	case "/db":
		dbHandler(w, r)
	case "/queries":
		queriesHandler(w, r)
	case "/fortunes":
		fortunesHandler(w, r)
	case "/updates":
		updatesHandler(w, r)
	case "/plaintext":
		plaintextHandler(w, r)
	default:
		w.WriteHeader(404)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	initDB()

	// http.HandleFunc("/json", jsonHandler)
	// http.HandleFunc("/db", dbHandler)
	// http.HandleFunc("/queries", queriesHandler)
	// http.HandleFunc("/fortunes", fortunesHandler)
	// http.HandleFunc("/updates", updatesHandler)
	// http.HandleFunc("/plaintext", plaintextHandler)
	// log.Fatal(http.ListenAndServe(":8080", nil))

	s := &http.Server{
		Addr:    ":8080",
		Handler: apiHandler{},
	}
	log.Fatal(s.ListenAndServe())
}
