package service

type Service struct {
	User
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) SetUser(u User) {
	s.User = u
}
