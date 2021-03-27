package adapter

// data := adapter.ByName("").Alloc()

// adapter.Select("").Build(data)

// adpater.Select("").Encode(data)

// adapter.Select("").Decode(encoded, data)

// adapter.Select("").Deploy(data)

func ByName(name string) Adapter {
	if name == "raw-go" {
		return NewRawGoAdapter()
	}
}

type AdapterData = interface{}

type Adapter interface {
	Alloc() AdapterData
	Assign(data AdapterData)
	Encode(data AdapterData) ([]byte, error)
	Decode(encoded []byte, data AdapterData)
	Deploy(data AdapterData) error
}
