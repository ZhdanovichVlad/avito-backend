package avito_backend

type UserService struct {
}

func NewUserService() *UserService {
	return new(UserService)
}

func (s *UserService) GetById(id int) (any, error) {
	return nil, nil
}

func main() {
	us := NewUserService()
	u, err := us.GetById(123)
	u, err = UserService.GetById(*us, 123)
}
