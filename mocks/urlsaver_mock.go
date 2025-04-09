package mocks

import (
	"github.com/stretchr/testify/mock"
)

type URLSaverMock struct {
	mock.Mock
}

func (m *URLSaverMock) SaveURL(urlToSave string, alias string) (int64, error) {
	args := m.Called(urlToSave, alias)
	return args.Get(0).(int64), args.Error(1)
}
