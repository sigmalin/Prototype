package main

// Main will set up an http server and three endpoints
func main() {
	initResponseProvider()
	initSessionProvider()

	go dbConnect()

	apiService()
}
