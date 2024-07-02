# st

[SpaceTraders](https://spacetraders.io/) is a multiplayer programming game
where the players use an
[HTTP-API](https://spacetraders.stoplight.io/docs/spacetraders) to manage a
fleet of ships for exploring, trading, and fighting their way across the
multiplayer universe.

This project provides a [Go](https://go.dev/) [module](https://go.dev/ref/mod) for interacting with the game.

---

### Usage

```go
import "github.com/nothub/spacetraders/pkg/client"
...
st.Token = "eyJh..."
agent, err := st.GetAgent()
...
```

For usage examples, check [cmd](./cmd).
