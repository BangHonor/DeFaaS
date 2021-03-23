package provider

import (
	"log"
	"testing"
	"time"
)

func TestNewDeployServer(t *testing.T) {

	s, err := NewDeployServer("", 1)
	if err != nil {
		t.Fatal(err)
	}
	if err := s.Start(); err != nil {
		t.Fatal(err)
	}

	log.Println("haha")

	time.Sleep(1 * time.Minute)

	if err := s.Stop(); err != nil {
		t.Fatal(err)
	}
}
