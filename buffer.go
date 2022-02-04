package gotester

const (
	defaultBufSize = 4096
	chunk          = 512
)

// chunkMultiple returns a multiple of chunk size closest to but greater than size.
func chunkMultiple(size int) int { return ((size / chunk) + 1) * chunk }

// InitialCapacity returns the multiple of 'chunk' one more than needed to
// accomodate the given capacity.
func InitialCapacity(capacity int) int {
	if capacity <= defaultBufSize {
		return chunkMultiple(capacity)
	}
	return defaultBufSize
}
