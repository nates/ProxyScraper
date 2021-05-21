# 🖨️ ProxyScraper
A multi-threaded proxy scraper made in Go. 
## 🧳 Requirements
* Go (latest)

## 🏗️ Building
```
go build ./src/scraper
```

## 🕹️ Usage
Put your proxy list websites into urls.txt separated by each line.
```
./scraper.exe -input=urls.txt -output=proxies.txt -timeout=5
```
