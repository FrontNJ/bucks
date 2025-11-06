package types

const (
	// ModuleName defines the module name
	ModuleName = "bucks"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_bucks"
)

var (
	ParamsKey = []byte("p_bucks")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
