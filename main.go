package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// configure the songs directory name and port
	const songsDir = "video"
	//const songsDir = "songs/crab-rave"

	const port = 8080

	// add a handler for the song files
	http.Handle("/", addHeaders(http.FileServer(http.Dir(songsDir))))
	fmt.Printf("Starting server on %v\n", port)
	log.Printf("Serving %s on HTTP port: %v\n", songsDir, port)

	// serve and log errors
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

// addHeaders will act as middleware to give us CORS support
func addHeaders(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}

//mp4 TO m3u8------   ffmpeg -i FileName.mp4 -b:v 1M -g 60
//-hls_time 2 -hls_list_size 0 -hls_segment_size 500000 output.m3u8
//mp3 TO m3u8------   ffmpeg -i FileName.mp3 -c:a libmp3lame -b:a 128k -map 0:0 -f segment -segment_time 10 -segment_list outputlist.m3u8 -segment_format mpegts output%03d.ts
