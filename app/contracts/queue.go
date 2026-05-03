package contracts

type Job func()

type Queue interface {
	Push(job Job)
}
