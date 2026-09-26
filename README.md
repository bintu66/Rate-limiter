# Rate Limiter

a simple HTTP rate limiter i built in Go for my web technology lab assignment.

## what does it do?

it limits how many times someone can hit my API in a given time window. if you send too many requests too fast, the server says "slow down" and gives you a `429 Too Many Requests` error.

i used the **Fixed Window Counter** algorithm because its the simplest one to understand — it just counts requests in a time window and resets after the window expires.

## how it works :

1. you send a request to the server
2. the server checks your IP address
3. it looks up how many requests you've already made in the current 60-second window
4. if you're under 5 requests → it lets you through 
5. if you've already made 5 → it blocks you with a 429 error 
6. after 60 seconds, the window resets and you can make requests again

## project structure

```
rate-limiter/
├── cmd/
│   └── server/
│       └── main.go              # starts the server
├── internal/
│   ├── handlers/
│   │   └── handlers.go          # the actual API routes
│   ├── middleware/
│   │   └── ratelimiter.go       # the rate limiting logic
│   └── models/
│       └── models.go            # data structures
├── .gitignore
├── go.mod
└── README.md
```

## how to run

```bash
go run ./cmd/server
```

server starts on `http://localhost:8080`

## how to test

```bash
# hit the ping endpoint
curl http://localhost:8080/ping

# hit it 6 times fast to see the rate limit kick in
for i in {1..6}; do curl -s http://localhost:8080/ping; echo; done
```

first 5 will say `{"message":"pong"}` and the 6th will say rate limit exceeded.


## tech used

- Go (standard library only, no external packages)
- `net/http` for the server
- `sync.Mutex` for thread safety
- `encoding/json` for JSON responses
