package service

import (
	"time"

	"github.com/nisfu-saaban/golang-goroutine/entity"
)

func FindByID(id string) <-chan *entity.Student {
	resChannel := make(chan *entity.Student)

	go func() {
		time.Sleep(2 * time.Second)

		student := &entity.Student{
			Id:   id,
			Name: "Sumarno",
		}

		resChannel <- student
		close(resChannel)
	}()

	return resChannel
}
