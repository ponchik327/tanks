package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Настройка обработчиков
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/game", gameHandler)

	// Настройка статических файлов (HTML, JS, CSS)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Запуск сервера
	port := ":8080"
	fmt.Printf("Сервер запущен на http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, "static/index.html")
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	// Здесь будет логика WebSocket соединения для игры
	fmt.Fprintf(w, "game endpoint для IO танков")
}
