package health

// https://codewithyury.com/how-to-create-health-check-for-restful-microservice-in-golang/

type HealthStatus struct {
	Status string `json:"status"`
}

type Service interface {
	Health() HealthStatus
}

type service struct {
	Status string
}

func New() Service {
	return &service{}
}

func (s *service) Health() HealthStatus {
	status := "ok"

	return HealthStatus{
		Status: status,
	}
}
