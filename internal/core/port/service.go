package port

type SystemStatusService interface {
	Status() error
}
