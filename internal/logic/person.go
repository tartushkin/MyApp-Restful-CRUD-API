package logic

import (
	"myApp/internal/app"
	"myApp/internal/myDb"

	"github.com/sirupsen/logrus"
)

type Logic struct {
	repo   *myDb.Repository
	logger *logrus.Logger
}

func NewLogic(repo *myDb.Repository, logger *logrus.Logger) *Logic {
	return &Logic{
		repo:   repo,
		logger: logger,
	}
}

// GetAllPerson возвращает список вех записей из базы
func (l *Logic) GetAllPerson(limit, offset int, search string) ([]app.Person, error) {
	persons, err := l.repo.GetAllPerson(0, 0, "")
	if err != nil {
		l.logger.Errorf("Ошибка при получении списка записей из базы :%v", err)
	}
	return persons, err
}

// CreatePerson создает новую запись в базе
func (l *Logic) CreatePerson(person *app.Person) error {
	err := l.repo.CreatePerson(person)
	if err != nil {
		l.logger.Errorf("Ошибка при создании записи в базе данных: %v", err)
	}
	return err
}

// GetPerson возвращает запись з базе по id
func (l *Logic) GetPerson(id int) (*app.Person, error) {
	person, err := l.repo.GetPerson(id)
	if err != nil {
		l.logger.Errorf("Ошибка при получении записи из базы: %v", err)
	}
	return person, err
}

// UpdatePerson обналяет запись в базе данных по id
func (l *Logic) UpdatePerson(id int, person *app.Person) error {
	err := l.repo.UpdatePerson(id, person)
	if err != nil {
		l.logger.Errorf("Ошибка при обновлении записи в базы: %v", err)
	}
	return err
}

// DeletePerson удаляет запись в базе
func (l *Logic) DeletePerson(id int) error {
	err := l.repo.DeletePerson(id)
	if err != nil {
		l.logger.Errorf("Ошибка при удалении записи из базы: %v", err)
	}
	return err
}
