# st

[SpaceTraders](https://spacetraders.io/) is a multiplayer programming game where the players use an
[HTTP-API](https://spacetraders.stoplight.io/docs/spacetraders) to manage a fleet of ships for
exploring, trading, and fighting their way across the multiplayer universe.

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

### Versioning

The project implements the games API as described in
[github.com/SpaceTradersAPI/api-docs](https://github.com/SpaceTradersAPI/api-docs).

The repository is tagged with semver-compatible version-tags in the format of
`v<MODVER>+<APIVER>` (e.g. `v0.1.0+2.1.0`), where `<MODVER>` references the
module release and `<API_VER>` references the implemented api version.

### Bugs

If you run into any bugs, please report them [here](https://github.com/nothub/spacetraders/issues). Thanks! 😊
