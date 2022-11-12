package storage

import (
	"github.com/luryon/go-banking/model"
)

// Memory es la estructura donde se cargan los datos para tranajr
type Memory struct {
	CurrentID int
	Accounts  map[int]model.Account
}

// NewMemoery devuelve una instancia de Memory
func NewMemory() Memory {
	return Memory{
		CurrentID: 0,
		Accounts:  make(map[int]model.Account),
	}
}

func (m *Memory) Create(acc *model.Account) error {
	if acc == nil {
		return ErrAccountCantBeNil
	}

	m.CurrentID += 1
	acc.Id = m.CurrentID
	m.Accounts[m.CurrentID] = *acc

	return nil
}

func (m *Memory) Update(id int, acc *model.Account) (*model.Account, error) {
	if acc == nil {
		return nil, ErrAccountCantBeNil
	}

	acc.Id = id
	m.Accounts[id] = *acc

	return acc, nil
}

func (m *Memory) GetAll() (*model.Accounts, error) {
	accts_registered := model.Accounts{}

	for _, v := range m.Accounts {
		data := v
		accts_registered = append(accts_registered, &data)
	}

	return &accts_registered, nil
}

func (m *Memory) GetById(id int) (*model.Account, error) {
	acc := m.Accounts[id]

	return &acc, nil
}

func (m *Memory) Delete(id int) error {
	delete(m.Accounts, id)
	return nil
}

func (m *Memory) Send(tx model.Operation) error {
	acc1 := m.Accounts[tx.Origin_acc]
	acc2 := m.Accounts[tx.Dest_acc]

	if tx.Amount > acc1.Amount_on_acc {
		return ErrAmountExceedAccountAmount
	}

	acc1.Amount_on_acc -= tx.Amount
	acc2.Amount_on_acc += tx.Amount

	m.Accounts[tx.Origin_acc] = acc1
	m.Accounts[tx.Dest_acc] = acc2

	return nil
}

func Modify(n int) int {

	return n
}
