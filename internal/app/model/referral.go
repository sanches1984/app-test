package model

import "time"

type Referral struct {
	UserID  int64     `firestore:"userId"`
	ChatID  int64     `firestore:"chatId"`
	Company string    `firestore:"company"`
	URL     string    `firestore:"url"`
	Comment string    `firestore:"comment"`
	Created time.Time `firestore:"created"`
}
