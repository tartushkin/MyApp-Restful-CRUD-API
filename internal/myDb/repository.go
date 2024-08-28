package myDb

import (
	"myApp/internal/app"
	"os"

	"github.com/gocraft/dbr/v2"
	"github.com/sirupsen/logrus"
	_ "modernc.org/sqlite"
)

// инициализация логгера
var logger = logrus.New()

func init() {
	//инициализация форматтера
	formatter := &logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}
	//форматер для логгера
	logger.SetFormatter(formatter)
	logger.SetLevel(logrus.DebugLevel)

	//выходной поток для логгера (stdout)
	logger.SetOutput(os.Stdout)
}

// для взаимодействия с базой данных
type Repository struct {
	//logger  *logrus.Logger //обьект логера
	session *dbr.Session //обьект сессии для запросов
}

func NewDb() (*dbr.Session, error) {
	db, err := dbr.Open("sqlite", "test.db", nil)
	if err != nil {
		logger.Fatalf("Ошибка подключения к БД: %v", err)
		return nil, err
	}
	logger.Info("Сессия БД на месте")
	return db.NewSession(nil), nil
}

// NewRep создает новый экзкмпляр репозитория с базой данных
func NewRepository(db *dbr.Session) *Repository {
	//создаем таблицу
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS person (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT,
		phone TEXT,
		firstName TEXT,
		lastName TEXT
	)`)
	if err != nil {
		logger.Fatalf("Ошибка создания таблицы: %v", err)
	}
	logger.Info("Таблица на месте")
	return &Repository{
		session: db,
	}
}

// GetAllPerson возвращает список всех записей в таблице Person
func (r *Repository) GetAllPerson(limit, offset int, search string) ([]app.Person, error) {
	var persons []app.Person
	query := r.session.Select("*").From("person")
	logger.Infof("Limt: %d, Offset: %d, Search: %s", limit, offset, search)
	if offset > 0 {
		query = query.Offset(uint64(offset))
	}
	if limit > 0 {
		query = query.Limit(uint64(limit))
	}
	if search != "" {
		query = query.Where("firstName LIKE ?", "%"+search+"%")
	}
	logger.Infof("Запрос: %v", query)
	_, err := query.Load(&persons)
	//_, err := r.session.Select("*").
	//	From("person").
	//	Limit(uint64(limit)).
	//	Offset(uint64(offset)).
	//	Where("firstName LIKE ?", "%"+search+"%").
	//	Load(&persons)
	if err != nil {
		logger.Errorf("Ошибка при получении всх записей из таблицы Person: %v", err)
	}
	return persons, err
}

// CreatePerson создает новую запись в базу
func (r *Repository) CreatePerson(person *app.Person) error {
	logger.Infof("Данные для создания записи: %+v", person)
	_, err := r.session.InsertInto("person").
		Columns("email", "phone", "firstName", "lastName").
		Record(person).
		Exec()
	if err != nil {
		logger.Errorf("Ошибка создания записи в базе данных: %v", err)
	} else {
		logger.Info("Запись успешно создана в базе ")
	}
	return err
}

// GetPerson возвращает запись из таблицы Person
func (r *Repository) GetPerson(id int) (*app.Person, error) {
	var person app.Person
	_, err := r.session.Select("*").From("person").Where("id = ?", id).Load(&person)
	if err != nil {
		logger.Errorf("Ошибка при получении записи из таблицы Person %v", err)

	}
	return &person, err

}

// UpdatePerson  обнавляет запись в таблице Person
func (r *Repository) UpdatePerson(id int, person *app.Person) error {
	_, err := r.session.Update("person").
		Set("email", person.Email).
		Set("phone", person.Phone).
		Set("firstName", person.FirstName).
		Set("lastName", person.LastName).
		Where("id = ?", id).
		Exec()
	if err != nil {
		logger.Errorf("Ошибка при обновлении записи в таблице Person: %v", err)
	}
	return err
}

// DeletePerson удаляет запсиь из таблицы Person
func (r *Repository) DeletePerson(id int) error {
	_, err := r.session.DeleteFrom("person").
		Where("id= ?", id).
		Exec()
	if err != nil {
		logger.Errorf("Ошибка при удалении записи в таблице Person: %v", err)
	}
	return err

}
