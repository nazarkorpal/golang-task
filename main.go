package main

import (
	"context"
	"fmt"
	"golang-task/config"
	"log"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

type Data struct {
	Timestamp time.Time
	Message   string
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	isClosed := false

	go func() {
		time.Sleep(time.Minute)
		cancel()
	}()

	config := config.NewConfig()
	wg := new(sync.WaitGroup)

	ch := make(chan Data)
	defer close(ch)

	wg.Add(2)
	go Client(ctx, wg, ch, config.SenderInterval, config.Message, &isClosed)
	go Server(ctx, wg, ch, &isClosed)

	wg.Wait()
}

func Server(ctx context.Context, wg *sync.WaitGroup, ch chan Data, isClosed *bool) {
	defer wg.Done()

	for {
		select {
		case data := <-ch:
			fmt.Printf("Timestamp: %s \nMessage: %s \n-----------------------\n", data.Timestamp.Format("2006-01-02 15:04:05Z07:00"), data.Message)
		case <-ctx.Done():
			*isClosed = true
			return
		}

		time.Sleep(5 * time.Second)
	}
}

func Client(ctx context.Context, wg *sync.WaitGroup, ch chan Data, interval uint, message string, isClosed *bool) {
	defer wg.Done()
	counter := 0

	for {
		counter++

		select {
		case <-time.After(time.Second * time.Duration(interval)):
			if !*isClosed {
				ch <- Data{
					Timestamp: time.Now(),
					Message:   fmt.Sprintf("%s #%d", message, counter),
				}
			}
		case <-ctx.Done():
			*isClosed = true
			return
		}
	}
}
