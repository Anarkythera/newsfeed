# Newsfeed API

## How To Run

### Run newsfeed tests

To run the newsfeed tests

```sh
docker-compose -f docker-compose-testing.yaml build
docker-compose -f docker-compose-testing.yaml up
```

To destroy the newsfeed testing environment

```sh
docker-compose -f docker-compose-testing.yaml down
```

### Run newfeed application

```sh
docker-compose -f docker-compose-prod.yaml build
docker-compose -f docker-compose-prod.yaml up
```

To destroy the newsfeed testing environment

```sh
docker-compose -f docker-compose-production.yaml down
```

example request for endpoint /GetNews

```sh
curl -H "Content-type: application/json" "http://localhost:12345/v1/GetNews?page=0&pageSize=10"
curl -X GET -H "Content-type: application/json" -d '{"page":0,"pageSize":10}' "http://localhost:12345/v1/GetNews"
```

example request for endpoint /GetFilteredNews, to get the first 10 results from BBC News UK

```sh
curl -X GET -H "Content-type: application/json" -d '{"page":0,"pageSize":10,"provider":["BBC News"],"category":["UK"]}' "http://localhost:12345/v1/GetFilteredNews"
```

## API documentation

Currently the newsfeed application has 2 endpoints:
/v1/GetNews that _requires_ the parameters page and pageSize.
/v1/GetFilteredNews that _requires_ the parameters page, pageSize, category and provider.

### GetNews endpoint

The /GetNews endpoint receives as parameters two ints: page and pageSize.
Page is the current page of news we're currently in, while pageSize is the number of elements we want to receive.
For instance if we want to load the first 10 news, we would the a request to the /GetNews endpoint with the query params ?page=0&pageSize=10. For the second set of 10 news we would send the params ?page=10&pageSize=10 and so on.
If there aren't any news to display on the range of page and pageSize sent, the server will return an empty array along with a 200 response code.

### GetFilteredNews endpoint

The /GetFilteredNews endpoint receives as parameters two ints: page and pageSize and two string arrays: category and provider.
The page and pageSize work the same way as in the /GetNews endpoint, while the provider and category parameters are string arrays that will store the providers and categories we want news from.
For instance if we want to load the first 10 news from BBC News UK we would send these parameters: ?page=0&pageSize=10?provider=[BBC News]&category=[UK].
If there aren't any news to display on the range of page and pageSize sent, the server will return an empty array along with a 200 response code.

## Future Improvements

A cache implementation either an in memory cache to start, or use a database like (redis) to store the cache.
Currently the thumbnail returned is based on a best effort basis, in the future we could crawl each article and create a thumbnail from the first found image, if the RSS has no thumbnails.
