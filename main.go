package main

import "final_exam/routers"

func main() {
	r := routers.SetupRouter()
	r.Run(":2019")
}
