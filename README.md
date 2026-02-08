# American Football Score Combinations — GraphQL API

A GraphQL API that calculates how many possible scoring combinations produce a given American football score. Built with Go, gqlgen, and a dynamic programming algorithm with cross-request caching.

> Originally built as a backend technical assessment for Studio Sol (May 2024).

## Problem

Given a score like `"3x15"`, determine how many unique combinations of plays (field goal, touchdown, touchdown+1, touchdown+2) can produce each team's score, then multiply the results.

**Possible plays:**
- Touchdown: 6 points
- Touchdown + extra point: 7 points
- Touchdown + 2-point conversion: 8 points
- Field goal: 3 points

## Tech Stack

| Technology | Purpose |
|------------|---------|
| Go 1.20 | Language |
| [gqlgen](https://github.com/99designs/gqlgen) | GraphQL code generation |
| [Ginkgo](https://github.com/onsi/ginkgo) + [Gomega](https://github.com/onsi/gomega) | BDD testing framework |
| Docker | Multi-stage build (scratch base) |
| GitHub Actions | CI pipeline |

## Architecture

```
cmd/server/              → Entry point
graph/                   → GraphQL schema, resolvers, models
src/football/
├── team_points_service  → Dynamic programming algorithm + cache
└── game_score_service   → Score parsing + coordination
```

- **Service layer** separates business logic from GraphQL resolvers
- **Dependency injection** — `GameScoreService` is injected into resolvers
- **Singleton cache** — `map[int]int` persisted across requests for O(1) amortized lookups

## Algorithm

I broke the solution into 3 steps:

![solution steps](assets/img/image.png)

### 1. Parse the score string

Split on `'x'` and convert each side to `int`.

![parse score](assets/img/image-1.png)

### 2. Calculate combinations per team (Dynamic Programming)

For each total score, compute how many unique play combinations produce it. The algorithm builds up from the smallest scores, reusing intermediate results:

```
function count_combinations(total_points)
  if total_points < 0
    return 0

  combinations = array[total_points + 1]
  combinations[0] = 1  // zero points = one way (no plays)

  possible_plays = [3, 6, 7, 8]

  for play in possible_plays
    for i from play to length(combinations)
      combinations[i] += combinations[i - play]

  return combinations[total_points]
```

Impossible scores (like 2 or 4) naturally accumulate 0 combinations.

### 3. Multiply both teams

`total = combinations(team1) × combinations(team2)`

If either team has an impossible score (0 combinations), the result is 0.

### Optimization: Cross-Request Caching

- **Worst case:** O(n) where n = score value
- **Amortized:** O(1) for cached scores
- A request for `"30x6"` computes all values up to 30 — subsequent requests for any score ≤ 30 return instantly from cache
- Implemented as a singleton `map[int]int` shared across requests

## Running

### Locally

```bash
go run cmd/server/server.go
# GraphQL Playground: http://localhost:8080
# GraphQL endpoint:   http://localhost:8080/graphql
```

Set `PORT` env var to change the default port.

### Docker

```bash
docker build -t football-scores .
docker run -d -p 8080:8080 football-scores
```

### Example Query

```graphql
mutation {
  verify(score: "3x15") {
    combinations
  }
}
# → { "data": { "verify": { "combinations": 4 } } }
```

## Tests

Three-layer BDD testing with Ginkgo/Gomega:

```bash
# Install Ginkgo CLI
go install github.com/onsi/ginkgo/v2/ginkgo@v2.6.1

# Run all tests
ginkgo ./...
```

- **Unit tests** — `TeamPointsService`: valid scores (0→1, 3→1, 320→6044) and impossible scores (2, 4 → 0)
- **Service tests** — `GameScoreService`: score parsing + combination multiplication
- **Integration tests** — full GraphQL mutation pipeline via gqlgen test client

## Development Process

1. Initialized project with gqlgen
2. Created multi-stage Dockerfile (scratch base for minimal image)
3. Set up Ginkgo/Gomega + GitHub Actions CI
4. **Wrote integration tests first** (TDD) — when they pass, the full pipeline works
5. Implemented the algorithm incrementally with unit tests
6. Added cross-request caching as optimization
