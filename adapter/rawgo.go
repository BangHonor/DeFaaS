package adapter

type RawGoAdapterData struct {
	// TODO
}

type RawGoAdapter struct {
	// TODO
}

func (adapter *RawGoAdapter) Alloc() AdapterData {

	data := &RawGoAdapterData{}

	return data
}

func (adapter *RawGoAdapter) Assign(data AdapterData, func(data AdapterData) ) {

	data, ok := data.(RawGoAdapterData)

}

func (adapter *RawGoAdapter) Encode(data AdapterData) ([]byte, error) {}

func (adapter *RawGoAdapter) Decode(encoded []byte, data AdapterData) {}

func (adapter *RawGoAdapter) Deploy(data AdapterData) error {}
