package main

import (
	"bufio"
	"flag"
	"log"
	"net/http"
	"os"
	"path"
)

// getUATargets reads the array of user agent targets from a newline separated
// file specified by `source`.
func getUATargets(source string) ([]string, error) {
	var targets []string
	f, err := os.Open(source)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		targets = append(targets, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return targets, nil
}

func main() {
	targetsPath := flag.String("targets", "targets.txt", "path to the user agent targets file")
	imagePath := flag.String("image", "image.jpg", "path to the image file")
	payloadPath := flag.String("payload", "payload.html", "path to the html payload")
	bindAddr := flag.String("bind", ":80", "bind address")
	flag.Parse()

	uaTargets, err := getUATargets(*targetsPath)
	if err != nil {
		log.Fatal("error reading user agent targets:", err)
	}
	th, err := newTrickyHandler(uaTargets, *imagePath, *payloadPath)
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/"+path.Base(*imagePath), th)
	log.Fatal(http.ListenAndServe(*bindAddr, nil))
}
