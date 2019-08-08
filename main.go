package main

import (
	"fmt"
	"geoirb/checker/app"
	"geoirb/checker/handling"
	"log"
	"os"
	"time"
)

const VERSION string = "1.0.0"
const KEYWORDS string = "Проверка наличия пользовательских слов"
const SYSTEM string = "Проверка наличия системных слов"
const HASH string = "Проверка хеш сумм"

func main() {

	if len(os.Args) > 0 && os.Args[1] == "-v" {
		log.Println("Версия: ", VERSION)

		log.Print("Описание: ")
		switch app.GetNameApp() {
		case "hash":
			log.Println(HASH)
		case "system":
			log.Println(SYSTEM)
		case "keywords":
			log.Println(KEYWORDS)
		}
		return
	}

	step := app.TickInit(time.Duration(app.Load("time")["sleep"].(int)) * time.Second)
	for {
		select {
		case <-step.Next:
			conn := app.Init()

			startTime := time.Now()
			handling.Start(conn)

			fmt.Printf("\nTime to check %v\n\n", time.Now().Sub(startTime))
			conn.Pause()
		}
		step.Wait()
	}
}
