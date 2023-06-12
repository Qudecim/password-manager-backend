package service

type SendingService struct {
}

func NewSendingService() *SendingService {
	return &SendingService{}
}

func (s *SendingService) SendSecret(uid string, secret string) error {
	return nil
}
