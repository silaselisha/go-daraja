package handler

// Deprecated aliases kept so existing callers keep compiling while they migrate
// to the Pokea names. They will be removed in the next major release.

// Deprecated: use Pokea.
type Daraja = Pokea

// Deprecated: use PokeaClient.
type DarajaClient = PokeaClient

// Deprecated: use PokeaResParams.
type DarajaResParams = PokeaResParams

// Deprecated: use PokeaErrorParams.
type DarajaErrorParams = PokeaErrorParams

// Deprecated: use PokeaAuth.
type DarajaAuth = PokeaAuth

// NewDarajaClient constructs a client from an optional .env path.
//
// Deprecated: use NewPokeaClient or NewClient.
func NewDarajaClient(path string) (Daraja, error) {
    return NewPokeaClient(path)
}
