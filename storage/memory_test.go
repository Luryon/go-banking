package storage

import (
	"testing"

	"github.com/luryon/go-banking/model"
)

func BenchmarkCreate(b *testing.B) {
	table := []struct {
		Name      string
		account   *model.Account
		wantError error
	}{
		{"empty person", nil, ErrAccountCantBeNil},
		{"Jorge", &model.Account{Name: "Jorge"}, nil},
		{"Luca", &model.Account{Name: "Luca"}, nil},
	}

	m := NewMemory()
	for _, v := range table {
		b.Run(v.Name, func(b *testing.B) {
			gotErr := m.Create(v.account)
			if gotErr != v.wantError {
				b.Errorf("Se esperaba %v, se obtuvo %v", v.wantError, gotErr)
			}
		})
	}

	wantCount := len(table) - 1
	if m.CurrentID != wantCount {
		b.Errorf("Se esperaba un elemento %d, se obtuvo %d", wantCount, m.CurrentID)
	}
}
