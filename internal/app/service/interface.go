package service

import tg "github.com/sanches1984/tg-client"

type UserService interface {
	RegisterUser(msg *tg.IncomingMessage) error
}
