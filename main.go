package main

import (
    "database/sql"
    "fmt"
    "log"
   

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // Подключение к удаленному серверу MySQL
    dsn := "evn-test:L7#dXdDSiGEGRz&@tcp(192.168.49.42:3306)/TEST"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Проверка соединения с базой данных
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Successfully connected to the remote database!")

    // Выполнение запроса
    rows, err := db.Query("SELECT id, name FROM users")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    // Вывод результатов
    for rows.Next() {
        var id int
        var name string
        if err := rows.Scan(&id, &name); err != nil {
            log.Fatal(err)
        }
        fmt.Printf("%d: %s\n", id, name)
    }

    if err := rows.Err(); err != nil {
        log.Fatal(err)
    }
}
