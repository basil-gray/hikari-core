package activity

type Service struct {
    Repo Repository
}

func NewService() *Service {
    return &Service{Repo: NewRepository()}
}

func (s *Service) CreateActivity(h Activity) (Activity, error) {
    return s.Repo.Save(h)
}

// Other service methods (GetActivity, UpdateActivity, DeleteActivity)