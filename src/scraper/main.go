package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	totalProxies = []string{}
	urls         = []string{}
)

func main() {
	input := flag.String("input", "urls.txt", "File to input urls")
	output := flag.String("output", "proxies.txt", "File to output scraped proxies")
	flag.Parse()

	file, err := os.Open(*input)
	if err != nil {
		fmt.Println("Error reading " + *input)
		return
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}
	file.Close()

	fmt.Println(urls)

	var wg sync.WaitGroup

	for i, url := range urls {
		wg.Add(1)
		go worker(i, &wg, url)
	}

	wg.Wait()

	fmt.Println("[Main] Scraped " + strconv.Itoa(len(totalProxies)) + " proxies")
	if len(totalProxies) == 0 {
		return
	}
	file, err = os.Create(*output)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = file.WriteString(strings.Join(totalProxies, "\n"))
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("[Main] Wrote proxies to " + *output)
}

func worker(id int, wg *sync.WaitGroup, url string) {
	fmt.Println("[" + strconv.Itoa(id+1) + "] Scraping @ " + url)
	proxies, err := scrape(url)
	if err != nil {
		fmt.Println("[" + strconv.Itoa(id+1) + "] " + err.Error() + " @ " + url)
		wg.Done()
		return
	}
	fmt.Println("[" + strconv.Itoa(id+1) + "] Scraped " + strconv.Itoa(len(proxies)) + " proxies @ " + url)
	for _, proxy := range proxies {
		totalProxies = append(totalProxies, proxy)
	}
	wg.Done()
}
