# FAQ

- Why does client construction succeed even if the network is down?
  - To improve local development experience, the auth flow maps network errors to a typed unreachable error rather than failing construction outright.

- Where do I put my `.env` file?
  - Place it in a directory and pass that directory to `WithEnvFile(path)` or `NewDarajaClient(path)`. The file name must be `.env`.
