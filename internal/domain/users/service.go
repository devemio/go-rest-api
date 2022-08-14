package users

type Service struct{}

func (s *Service) GetTemporaryUrl(url string) string {
	return url
}
