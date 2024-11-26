package water

import "runtime"

func (ifce *Interface) SetTunNetwork(network string) error {
	if runtime.GOOS == "windows" {
		rwfile := ifce.ReadWriteCloser.(*wfile)
		return setTUN(rwfile.fd, network)

	}
	return nil
}
