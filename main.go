package main

import (
	"fmt"
	"github.com/jackSpanrrows/go-share-library/src/downloader"
	"log"
	"runtime"
	"time"
)

const (
	numDowner      = 10   //下载进程
	chanBufferSize = 1000 //通道缓存
	version        = 1.1
)

func init() {
	log.Printf("当前版本:%v", version)
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var URL string
	fmt.Print("Enter url:")
	fmt.Scanf("%s", &URL)
	log.Println("Working, pls wait ...")

	downChan := make(chan downloader.Item, chanBufferSize)
	okChan := make(chan int, chanBufferSize)
	failChan := make(chan int, chanBufferSize)

	if a, b, err := downloader.ParseURL(URL); err == nil {
		var result []downloader.Item
		switch b {
		case "Album":
			var in downloader.Album
			result = in.Action(a)
		case "Song":
			var in downloader.Song
			result = in.Action(a)
		case "Artist":
			var in downloader.Artist
			result = in.Action(a)
		case "ArtistAlbum":
			var in downloader.ArtistAlbum
			result = in.Action(a)
		case "Playlist":
			var in downloader.Playlist
			result = in.Action(a)
		case "Program":
			var in downloader.Program
			result = in.Action(a)
		case "DJradio":
			var in downloader.DJradio
			result = in.Action(a)
		}

		for i := 0; i < len(result); i++ {
			downChan <- result[i]
			okChan <- i
		}
	} else {
		log.Println(err.Error())
	}

	for i := 0; i < numDowner; i++ {
		go func() {
			for {
				downloader.DownloadFile(<-downChan, downChan, okChan, failChan)
			}
		}()
	}

//WAITING:
	if len(okChan) != 0 || len(downChan) != 0 || (len(downChan)-len(okChan)) != 0{
		fmt.Printf("=====================================\nDownloading: %v\t\tQueuing: %v\tFail:%v\n=====================================\n", len(downChan), len(okChan), len(okChan))
		time.Sleep(3 * time.Second)
		log.Println("\nAll tasks finished1.")
		//goto WAITING
	}
	fmt.Printf("=====================================\nDownloading: %v\t\tQueuing: %v\tFail:%v\n=====================================\n", len(downChan), len(okChan), len(okChan))
	log.Println("\nAll tasks finished2.")
}
