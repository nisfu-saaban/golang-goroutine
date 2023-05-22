package service

import (
	"fmt"
	"testing"
)

func TestFindByID(t *testing.T) {
	resultCh := FindByID("1")
	student := <-resultCh

	fmt.Println(student)
}
