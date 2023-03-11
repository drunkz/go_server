//go:build !debug

package mode

func IsDebugMode() bool {
	return false
}
