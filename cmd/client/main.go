package main

import (
	"context"
	"flag"
	"fmt"
	desc "github.com/romik1505/youtubeThumbnails/pkg/api/thumbnails"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	inputFile := flag.String("f", "./input.txt", "file for urls")
	outputDir := flag.String("o", "./output", "directory for output files")
	async := flag.Bool("async", false, "async requests to grpc server")
	flag.Parse()
	if *async {
		log.Println("Started in async mode")
	}

	data, err := os.ReadFile(*inputFile)
	if err != nil {
		log.Fatalln(err)
	}
	urls := strings.Split(string(data), "\n")

	if err := os.MkdirAll(*outputDir, 0777); err != nil {
		log.Fatalf("Error creating output directory: %v", err)
	}

	ctx := context.Background()
	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure()) // nolint
	if err != nil {
		log.Fatalln("bad connection")
	}

	client := desc.NewThumbnailClient(conn)

	defer func(t time.Time) {
		fmt.Printf("Time since: %s", time.Since(t))
	}(time.Now())

	if *async {
		asyncSaveThumbnails(ctx, client, outputDir, urls)
	} else {
		syncSaveThumbnails(ctx, client, outputDir, urls)
	}
}

func asyncSaveThumbnails(ctx context.Context, client desc.ThumbnailClient, outputDir *string, urls []string) {
	var wg sync.WaitGroup

	wg.Add(len(urls))

	for i, url := range urls {
		go func(i int, url string) {
			defer wg.Done()
			saveThumbnail(ctx, client, url, fmt.Sprintf("%s/%d.jpg", *outputDir, i+1)) //nolint
		}(i, url)
	}
	wg.Wait()
}

func syncSaveThumbnails(ctx context.Context, client desc.ThumbnailClient, outputDir *string, urls []string) {
	for i, url := range urls {
		saveThumbnail(ctx, client, url, fmt.Sprintf("%s/%d.jpg", *outputDir, i+1)) //nolint
	}
}

func saveThumbnail(ctx context.Context, client desc.ThumbnailClient, url string, filename string) error {
	resp, err := client.Get(ctx, &desc.GetRequest{
		Url: url,
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	if err := ioutil.WriteFile(filename, resp.GetImage(), 0777); err != nil {
		log.Println("save error", err.Error())
		return err
	}
	log.Printf("%s file loaded", filename)

	return nil
}
