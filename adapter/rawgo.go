package adapter

import "defaas/core/data"

type RawGoAdapter struct {
	GoFilePath string
}

func NewRawGoAdapter() *RawGoAdapter {
	return nil
}

func (adapter *RawGoAdapter) Encode() ([]byte, error) {
	return nil, nil
}

func (adapter *RawGoAdapter) Decode(encoded []byte) error {
	return nil
}

func (adapter *RawGoAdapter) Deploy(item *data.DeploymentItem) error {
	return nil
}
