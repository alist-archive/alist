package data

import (
	"context"

	"codeberg.org/alist/alist/v3/cmd/flags"
	"codeberg.org/alist/alist/v3/internal/db"
	"codeberg.org/alist/alist/v3/internal/message"
	"codeberg.org/alist/alist/v3/internal/model"
	"codeberg.org/alist/alist/v3/internal/op"
	log "github.com/sirupsen/logrus"
)

func initDevData() {
	_, err := op.CreateStorage(context.Background(), model.Storage{
		MountPath: "/",
		Order:     0,
		Driver:    "Local",
		Status:    "",
		Addition:  `{"root_folder_path":"."}`,
	})
	if err != nil {
		log.Fatalf("failed to create storage: %+v", err)
	}
	err = db.CreateUser(&model.User{
		Username:   "Noah",
		Password:   "hsu",
		BasePath:   "/data",
		Role:       0,
		Permission: 512,
	})
	if err != nil {
		log.Fatalf("failed to create user: %+v", err)
	}
}

func initDevDo() {
	if flags.Dev {
		go func() {
			err := message.GetMessenger().WaitSend(message.Message{
				Type:    "string",
				Content: "dev mode",
			}, 10)
			if err != nil {
				log.Debugf("%+v", err)
			}
			m, err := message.GetMessenger().WaitReceive(10)
			if err != nil {
				log.Debugf("%+v", err)
			} else {
				log.Debugf("received: %+v", m)
			}
		}()
	}
}
