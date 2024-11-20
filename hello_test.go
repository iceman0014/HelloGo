package hellogo_test

import (
	"fmt"
	"testing"

	hellogo "github.com/iceman0014/HelloGo"
)

func TestHello(t *testing.T) {
	data := "jack"
	expected := fmt.Sprintf("hello %s!", data)
	result := hellogo.Hello(data)

	if result != expected {
		t.Fatalf("expected result %s, but got %s", expected, result)
	}
}
