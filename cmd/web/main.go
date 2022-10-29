package main

import (
	"context"
	"flag"
	_ "github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/pgxpool"
	"log"
	"net/http"
	"os"
	"snippetbox.KazbekTutinbay.net/internal/models"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *models.SnippetModel
}

func main() {
	addr := flag.String("addr", "localhost:4000", "HTTP network address")
	dsn := flag.String("dsn", "postgres://web_user:123@localhost:5432/snippetbox", "connection login to database")
	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	conn, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatalf("\nUnable to connect to database: %v\n", err)
		return
	}
	infoLog.Print("Connected to database")
	defer conn.Close()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &models.SnippetModel{Conn: conn},
	}
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("running server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)

}

func openDB(dsn string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	if err = pool.Ping(context.Background()); err != nil {
		return nil, err
	}

	return pool, err
}

//перенаправить потоки из stdout и stderr в файлы на диске при запуске приложения из терминала
//go run ./cmd/web >>/tmp/info.log 2>>/tmp/error.log
//go run ./cmd/web -addr=":5000" изменить среду запуска программы в терминале
//mv ui/html/pages/home.tmpl ui/html/pages/home.bak исскуственный internal error
//go get -u github.com/jackc/pgx обновить до последней версии пакет с pgx
//psql -d snippetbox -U web_user -p 5432 вход в базу через терминал
//"github.com/jackc/pgx/v5/pgconn"

//task list
//изучить глубже транзакции https://golangify.com/transactions-and-other-details
