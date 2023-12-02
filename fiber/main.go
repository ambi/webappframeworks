package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"sort"
	"strconv"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

const (
	// Database
	worldSelect        = "SELECT id, randomNumber FROM World WHERE id = ?"
	worldUpdate        = "UPDATE World SET randomNumber = ? WHERE id = ?"
	fortuneSelect      = "SELECT id, message FROM Fortune;"
	worldRowCount      = 10000
	maxConnectionCount = 256
)

type Message struct {
	Message string `json:"message"`
}

type World struct {
	Id           uint16 `json:"id"`
	RandomNumber uint16 `json:"randomNumber"`
}

type Fortune struct {
	Id      uint16 `json:"id"`
	Message string `json:"message"`
}

type Fortunes []*Fortune

func (s Fortunes) Len() int      { return len(s) }
func (s Fortunes) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type ByMessage struct{ Fortunes }

func (s ByMessage) Less(i, j int) bool { return s.Fortunes[i].Message < s.Fortunes[j].Message }

var (
	// Database
	worldStatement   *sql.Stmt
	fortuneStatement *sql.Stmt
	updateStatement  *sql.Stmt
)

func parseQueries(c *fiber.Ctx) int {
	n, err := strconv.Atoi(c.Query("n"))
	if err != nil {
		n = 1
	} else if n < 1 {
		n = 1
	} else if n > 500 {
		n = 500
	}
	return n
}

// / Test 1: JSON serialization
func json(c *fiber.Ctx) error {
	return c.JSON(Message{Message: "Hello, World!"})
}

// / Test 2: Single database query
func db(c *fiber.Ctx) error {
	var world World
	err := worldStatement.QueryRow(rand.Intn(worldRowCount)+1).Scan(&world.Id, &world.RandomNumber)
	if err != nil {
		return err
	}
	return c.JSON(&world)
}

// / Test 3: Multiple database queries
func dbs(c *fiber.Ctx) error {
	numQueries := parseQueries(c)

	worlds := make([]World, numQueries)
	for i := 0; i < numQueries; i++ {
		err := worldStatement.QueryRow(rand.Intn(worldRowCount)+1).Scan(&worlds[i].Id, &worlds[i].RandomNumber)
		if err != nil {
			return err
		}
	}
	return c.JSON(&worlds)
}

// / Test 4: Fortunes
func fortunes(c *fiber.Ctx) error {
	rows, err := fortuneStatement.Query()
	if err != nil {
		return err
	}

	fortunes := make(Fortunes, 0)
	for rows.Next() { //Fetch rows
		fortune := Fortune{}
		if err := rows.Scan(&fortune.Id, &fortune.Message); err != nil {
			return err
		}
		fortunes = append(fortunes, &fortune)
	}
	fortunes = append(fortunes, &Fortune{Message: "Additional fortune added at request time."})

	sort.Sort(ByMessage{fortunes})

	c.Response().Header.SetContentType(fiber.MIMETextHTMLCharsetUTF8)
	return c.Render("fortunes", fortunes)
}

// / Test 5: Database updates
func update(c *fiber.Ctx) error {
	numQueries := parseQueries(c)
	world := make([]World, numQueries)
	for i := 0; i < numQueries; i++ {
		if err := worldStatement.QueryRow(rand.Intn(worldRowCount)+1).Scan(&world[i].Id, &world[i].RandomNumber); err != nil {
			return err
		}
		world[i].RandomNumber = uint16(rand.Intn(worldRowCount) + 1)
		if _, err := updateStatement.Exec(world[i].RandomNumber, world[i].Id); err != nil {
			return err
		}
	}
	return c.JSON(world)
}

// / Test 6: plaintext
func plaintext(c *fiber.Ctx) error {
	_, err := c.WriteString("Hello, World!")
	return err
}

func main() {
	engine := html.New("./templates", ".html")

	cfg := fiber.Config{Views: engine}
	for i := range os.Args[1:] {
		if os.Args[1:][i] == "-prefork" {
			cfg.Prefork = true
		}
	}

	app := fiber.New(cfg)

	app.Use(func(c *fiber.Ctx) error {
		c.Response().Header.Set("Server", "fiber")
		return c.Next()
	})
	app.Get("/json", json)
	app.Get("/db", db)
	app.Get("/queries", dbs)
	app.Get("/fortunes", fortunes)
	app.Get("/updates", update)
	app.Get("/plaintext", plaintext)
	fmt.Println("Listening and serving HTTP")
	app.Listen(":8080")
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	dsn := "root@tcp(%s:3306)/hello_world"
	dbhost := "localhost"

	db, err := sql.Open("mysql", fmt.Sprintf(dsn, dbhost))
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	db.SetMaxIdleConns(maxConnectionCount)
	worldStatement, err = db.Prepare(worldSelect)
	if err != nil {
		log.Fatal(err)
	}
	fortuneStatement, err = db.Prepare(fortuneSelect)
	if err != nil {
		log.Fatal(err)
	}
	updateStatement, err = db.Prepare(worldUpdate)
	if err != nil {
		log.Fatal(err)
	}
}
