package interfaces

import "github.com/containernetworking/cni/libcni"

type CNI interface {
	ConfListFromFile(file string) (*libcni.NetworkConfigList, error)
	ConfListFromBytes(bytes []byte) (*libcni.NetworkConfigList, error)
	ConfFromFile(file string) (*libcni.NetworkConfig, error)
	ConfFromBytes(bytes []byte) (*libcni.NetworkConfig, error)
	ConfListFromConf(conf *libcni.NetworkConfig) (*libcni.NetworkConfigList, error)
	ConfListFromConfBytes(confBytes []byte) (*libcni.NetworkConfigList, error)
	ConfFiles(dir string, ext []string) ([]string, error)
}

type Cni struct{}

func (c *Cni) ConfListFromFile(file string) (*libcni.NetworkConfigList, error) {
	return libcni.ConfListFromFile(file)
}

func (c *Cni) ConfListFromBytes(bytes []byte) (*libcni.NetworkConfigList, error) {
	return libcni.ConfListFromBytes(bytes)
}

func (c *Cni) ConfFromFile(file string) (*libcni.NetworkConfig, error) {
	return libcni.ConfFromFile(file)
}

func (c *Cni) ConfFromBytes(bytes []byte) (*libcni.NetworkConfig, error) {
	return libcni.ConfFromBytes(bytes)
}

func (c *Cni) ConfListFromConf(conf *libcni.NetworkConfig) (*libcni.NetworkConfigList, error) {
	return libcni.ConfListFromConf(conf)
}

func (c *Cni) ConfFiles(dir string, ext []string) ([]string, error) {
	return libcni.ConfFiles(dir, ext)
}

func (c *Cni) ConfListFromConfBytes(confBytes []byte) (*libcni.NetworkConfigList, error) {
	conf, err := libcni.ConfFromBytes(confBytes)
	if err != nil {
		return nil, err
	}
	confList, err := libcni.ConfListFromConf(conf)
	if err != nil {
		return nil, err
	}
	return confList, nil
}
