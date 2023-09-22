package server

func Start() {
	router := setRouter()

	router.Run(":8080")
}
