package main

import (
	"fmt"
	"server"
	"video"
)

func main() {
	server.Init()
	video2, _ := video.RequestVideo("TestVideo")
	fmt.Println(string(video2.Media))
}
