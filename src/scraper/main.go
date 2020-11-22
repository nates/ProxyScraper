package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	totalProxies []string
	urls         []string
	timeout      int    = 5
	input        string = "urls.txt"
	output       string = "proxies.txt"
)

func main() {
	flag.StringVar(&input, "input", input, "File to input urls")
	flag.StringVar(&output, "output", output, "File to output scraped proxies")
	flag.IntVar(&timeout, "timeout", timeout, "Timeout to proxy list websites.")
	flag.Parse()

	file, err := os.Open(input)
	if err != nil {
		log.Println("Error reading " + input)
		return
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}
	file.Close()

	var wg sync.WaitGroup

	for i, url := range urls {
		wg.Add(1)
		go worker(i, &wg, url)
	}

	wg.Wait()

	totalProxies = uniqueArray(totalProxies)
	log.Println("[Main] Scraped " + strconv.Itoa(len(totalProxies)) + " unique proxies")
	if len(totalProxies) == 0 {
		return
	}
	file, err = os.Create(output)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = file.WriteString(strings.Join(totalProxies, "\n"))
	if err != nil {
		log.Println(err)
		file.Close()
		return
	}
	err = file.Close()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("[Main] Wrote proxies to " + output)
}

func worker(id int, wg *sync.WaitGroup, url string) {
	log.Println("[" + strconv.Itoa(id+1) + "] Scraping @ " + url)
	proxies, err := scrape(url, timeout)
	if err != nil {
		log.Println("[" + strconv.Itoa(id+1) + "] " + err.Error() + " @ " + url)
		wg.Done()
		return
	}
	log.Println("[" + strconv.Itoa(id+1) + "] Scraped " + strconv.Itoa(len(proxies)) + " proxies @ " + url)
	for _, proxy := range proxies {
		totalProxies = append(totalProxies, proxy)
	}
	wg.Done()
}

func uniqueArray(array []string) []string {
	m := make(map[string]bool)
	for _, item := range array {
		_, ok := m[item]
		if ok == false {
			m[item] = true
		}
	}
	var unique []string
	for item := range m {
		unique = append(unique, item)
	}
	return unique
}
