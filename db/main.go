//Пакет по работе с БД
package db

//Библиотеки с которыми работает данный пакет
import (
	"log"
	"os"
	"time"

	pg "github.com/go-pg/pg"
)

//Функция инициализации соединения с БД
func Connect() *pg.DB {

	//Параметры подсоединения к БД. Хорошо бы вынести в отдельный файл
	opts := &pg.Options{
		User:         "postgres",
		Password:     "17032016",
		Addr:         "localhost:5434",
		DialTimeout:  30 * time.Second,
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 1 * time.Minute,
		IdleTimeout:  30 * time.Minute,
	}
	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect to database.\n")
		os.Exit(100)
	}
	log.Printf("Connection to database successful.\n")

	CreateMyTable(db)
	return db

	//closeErr := db.Close()
	//if closeErr != nil {
	//	log.Printf("Error while closing the connection, Reason: %v\n", closeErr)
	//	os.Exit(100)
	//}
	//log.Printf("Connection closed successfully.\n")
}
