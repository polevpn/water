//go:build !windows
// +build !windows

package water

func (ifce *Interface) SetTunNetwork(network string) error {
	return nil
}
