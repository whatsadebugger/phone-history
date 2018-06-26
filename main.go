package main

func main() {
	app := NewApplication()

	router := setupRouter(app)

	router.Run(":8080")
}
