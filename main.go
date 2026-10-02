package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

const (
	clearScreen = "\x1b[2J"
	cursorHome  = "\x1b[H"
	hideCursor  = "\x1b[?25l"
	showCursor  = "\x1b[?25h"
	reset       = "\x1b[0m"
	red         = "\x1b[31m"
	green       = "\x1b[32m"
)

/*var frames = []string{
	red + "  __\n /  \\\n|    |\n \\__/\n /  \\\n^^  ^^" + reset,
	green + "\n  __\n /  \\\n|    |\n \\__/\n /  \\\n^^  ^^" + reset,
}*/

func handler(w http.ResponseWriter, r *http.Request) {
	if !strings.Contains(r.UserAgent(), "curl") &&
		!strings.Contains(r.UserAgent(), "wget") {
		fmt.Fprint(w, "Try: curl localhost:3000")
		return
	}
	//above part for handle when it access through methods other than curl

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming unsupported", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")

	fmt.Fprint(w, hideCursor+clearScreen+cursorHome)
	flusher.Flush()

	for i := 0; ; i++ {
		fmt.Fprint(w, clearScreen+cursorHome+frames[i%len(frames)])
		flusher.Flush()
		time.Sleep(75 * time.Millisecond)
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
