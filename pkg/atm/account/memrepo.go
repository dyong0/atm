package account

func NewMemRepository() Repository {
	return &memRepository{
		accounts: make(map[string]*Account),
	}
}

type memRepository struct {
	Repository
	accounts map[string]*Account
}

func (r *memRepository) Account(id string) (*Account, error) {
	acc, ok := r.accounts[id]
	if !ok {
		return nil, ErrAccountNotFound
	}

	return acc, nil
}

func (r *memRepository) Create(acc Account, holderName string, password string) error {
	acc.id = holderName
	acc.pw = password
	r.accounts[acc.ID()] = &acc

	return nil
}
