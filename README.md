# st

[SpaceTraders](https://spacetraders.io/) is a multiplayer programming game where the players use an
[HTTP-API](https://spacetraders.stoplight.io/docs/spacetraders) to manage a fleet of ships for
exploring, trading, and fighting their way across the multiplayer universe.

This project provides a hand-rolled [Go](https://go.dev/) [module](https://go.dev/ref/mod) for interacting with the game.

---

### Usage

For examples, check [./cmd](./cmd).

### Versioning

The project implements the games API as described in
[github.com/SpaceTradersAPI/api-docs](https://github.com/SpaceTradersAPI/api-docs).

Releases are marked with semver-compatible git tags in the format of
`v<MODVER>+<APIVER>` (e.g. `v0.1.0+2.1.0`), where `<MODVER>` references the
module release and `<API_VER>` references the implemented api version.

### Bugs

If you run into any bugs, please report them [here](https://github.com/nothub/spacetraders/issues). Thanks! 😊
