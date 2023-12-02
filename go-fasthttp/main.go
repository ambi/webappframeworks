package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"runtime"
	"sort"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sugawarayuuta/sonnet"
	"github.com/valyala/fasthttp"
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
	methodGET  = []byte("GET")
)

// Test 1: JSON serialization
func jsonHandler(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set(HeaderContentType, MIMEApplicationJSON)
	sonnet.NewEncoder(ctx).Encode(Message{"Hello, World!"})
}

// Test 2: Single database query
func dbHandler(ctx *fasthttp.RequestCtx) {
	world := new(World)
	if err := fetchRandomWorld(world); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	ctx.Response.Header.Set(HeaderContentType, MIMEApplicationJSON)
	sonnet.NewEncoder(ctx).Encode(world)
}

// Test 3: Multiple database queries
func queriesHandler(ctx *fasthttp.RequestCtx) {
	n := getQueryCount(ctx.URI().QueryArgs().Peek("n"))
	worlds := make([]World, n)
	for i := 0; i < n; i++ {
		if err := fetchRandomWorld(&worlds[i]); err != nil {
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			return
		}
	}

	ctx.Response.Header.Set(HeaderContentType, MIMEApplicationJSON)
	sonnet.NewEncoder(ctx).Encode(worlds)
}

// Test 4: Fortunes
func fortunesHandler(ctx *fasthttp.RequestCtx) {
	rows, err := fortuneSelectStmt.Query()
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	defer rows.Close()

	fortunes, err := fetchFortunes(rows)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	fortunes = append(fortunes, &Fortune{Message: "Additional fortune added at request time."})
	sort.Sort(FortunesByMessage{fortunes})

	ctx.Response.Header.Set(HeaderContentType, MIMETextHTML)
	fortuneTmpl.Execute(ctx, fortunes)
}

// Test 5: Database updates
func updatesHandler(ctx *fasthttp.RequestCtx) {
	n := getQueryCount(ctx.URI().QueryArgs().Peek("n"))
	worlds := make([]World, n)
	for i := 0; i < n; i++ {
		// Fetch and modify
		world := &worlds[i]
		if err := fetchRandomWorld(&worlds[i]); err != nil {
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			return
		}
		world.RandomNumber = uint16(randomWorldNum())

		// Update
		if _, err := worldUpdateStmt.Exec(world.RandomNumber, world.ID); err != nil {
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			return
		}
	}

	ctx.Response.Header.Set(HeaderContentType, MIMEApplicationJSON)
	sonnet.NewEncoder(ctx).Encode(worlds)
}

// Test 6: Plaintext
func plaintextHandler(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set(HeaderContentType, MIMETextPlain)
	ctx.Write(helloWorld)
}

func fetchRandomWorld(w *World) error {
	n := randomWorldNum()
	return worldSelectStmt.QueryRow(n).Scan(&w.ID, &w.RandomNumber)
}

func randomWorldNum() int {
	return rand.Intn(worldRowCount) + 1
}

func getQueryCount(q []byte) int {
	n, err := strconv.Atoi(string(q))
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

func handler(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Server", "Go net/http")

	if !bytes.Equal(ctx.Method(), methodGET) {
		ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		return
	}

	switch string(ctx.Path()) {
	case "/json":
		jsonHandler(ctx)
	case "/db":
		dbHandler(ctx)
	case "/queries":
		queriesHandler(ctx)
	case "/fortunes":
		fortunesHandler(ctx)
	case "/updates":
		updatesHandler(ctx)
	case "/plaintext":
		plaintextHandler(ctx)
	default:
		ctx.SetStatusCode(fasthttp.StatusNotFound)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	initDB()

	log.Fatal(fasthttp.ListenAndServe(":8080", handler))
}
