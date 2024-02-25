package main

import (
	"GoGPT/bot"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(len(bot.Bots))

	for name := range bot.Bots {
		go func(n string) {
			defer wg.Done()
			bot.Algalon(bot.Bots[n])
		}(name)
	}

	wg.Wait()
}
