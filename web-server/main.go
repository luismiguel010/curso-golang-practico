package main

func main() {
	server := NewServer(":3000")
	server.Handle("GET", "/", HandlerRoot)
	server.Handle("POST", "/api", server.AddMiddleware(HandleHome, ChackAuth(), Logging()))
	server.Listen()
}