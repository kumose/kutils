// Check if epoll exclusive available on the host

//go:build !cgo || !linux
// +build !cgo !linux

package insight

// checkEpollExclusive checks if the host system support epoll exclusive mode
func checkEpollExclusive() bool {
	// If CGO is disabled, always report false
	return false
}
