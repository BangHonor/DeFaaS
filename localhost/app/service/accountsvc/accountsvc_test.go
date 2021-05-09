package accountsvc

import "testing"

func TestLoad(t *testing.T) {
	Service()
}

func TestStore(t *testing.T) {
	svc := Service()
	svc.store()
}
