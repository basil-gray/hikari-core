package activity

type Repository interface {
    Save(h Activity) (Activity, error)
    FindByID(id int) (Activity, error)
    Update(h Activity) (Activity, error)
    Delete(id int) error
}

type InMemoryRepository struct {
    activity map[int]Activity
}

func NewRepository() Repository {
    return &InMemoryRepository{activity: make(map[int]Activity)}
}

func (r *InMemoryRepository) Save(h Activity) (Activity, error) {
    // Implement save logic
}

// Other repository methods (FindByID, Update, Delete)