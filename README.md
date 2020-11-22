# ProxyScraper
A multi-threaded proxy scraper made in Go.
## Requirements
* Go (latest)

## Installation
```
go build ./src/scraper
```

## Setup
Put your proxy list websites into urls.txt
```
https://www.proxy-list.download/api/v1/get?type=https
https://api.proxyscrape.com/?request=getproxies&proxytype=http&timeout=10000&country=all&ssl=all&anonymity=all
```

## Usage
```
./scraper.exe -input=urls.txt -output=proxies.txt -timeout=5
```
