# Rate Limiter in Go

This repository demonstrates three different approaches to implementing rate limiting in Go:

1. **Per-Client Rate Limiting (IP-Based)**
2. **Token Bucket Algorithm**
3. **Tollbooth Middleware**

Each implementation is organized into separate folders for clarity and modularity.

---

## 1. Per-Client Rate Limiting (IP-Based)

This method limits requests per unique IP address. It maintains a map of clients and applies rate limiting per client using the `golang.org/x/time/rate` package.

### Key Features:
- Tracks clients based on IP
- Uses `rate.Limiter` for rate limiting
- Cleans up inactive clients after 3 minutes

### How to Run:
```sh
cd perclientlimit
go run main.go
```

---

## 2. Token Bucket Algorithm

A general rate-limiting technique that allows a burst of requests up to a limit and refills tokens at a fixed rate.

### Key Features:
- Implements a token bucket system
- Uses `rate.Limiter(2, 4)`, allowing 2 requests per second with a burst capacity of 4

### How to Run:
```sh
cd tokenbucket
go run main.go
```

---

## 3. Tollbooth Middleware

This method uses the `tollbooth` package to implement rate limiting as middleware, making it easy to integrate with existing applications.

### Key Features:
- Simple and effective middleware
- Customizable error messages and response types
- Limits requests using `tollbooth.NewLimiter(1, nil)`, allowing 1 request per second

### How to Run:
```sh
cd tollbooth
go run main.go
```

---

## API Endpoints

| Endpoint        | Description                        | Rate Limit |
|---------------|--------------------------------|------------|
| `/pingp`      | IP-based rate limiter         | 2 req/sec, burst 4 |
| `/ping`       | Token bucket rate limiter     | 2 req/sec, burst 4 |
| `/pingtoll`   | Tollbooth middleware limiter  | 1 req/sec |

---

## Installation & Dependencies

Ensure you have Go installed and run:
```sh
go mod tidy
```

Install Tollbooth if needed:
```sh
go get -u github.com/didip/tollbooth/v7
```

---

## Author
[Likhith K.P](https://github.com/likhithkp)

## License
MIT License
