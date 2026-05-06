package contracts

import "time"

type Queue interface {
	Push(name string, payload any) error
	Later(name string, payload any, delay time.Duration) error
	Pop(name string) (string, error)
}
