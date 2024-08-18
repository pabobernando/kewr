package types

const (
	// ModuleName defines the module name
	ModuleName = "kewr"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_kewr"
)

var (
	ParamsKey = []byte("p_kewr")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
