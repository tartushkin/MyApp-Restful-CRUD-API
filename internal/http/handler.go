package http

import (
	"myApp/internal/app"
	"myApp/internal/logic"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// PersonHandlers обработчик запросов связанных с моделью Person
type PersonHandlers struct {
	logic *logic.Logic
}

// NewPersonHandlers создает новый экземпляр обработчиков запросов для Person
func NewPersonHandlers(logic *logic.Logic) *PersonHandlers {
	return &PersonHandlers{logic: logic}
}

// GetAllPerson обрабатывает GET запрос для получения списка записей из базы
func (ph *PersonHandlers) GetAllPerson(c echo.Context) error {
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	offset, _ := strconv.Atoi(c.QueryParam("offset"))
	search := c.QueryParam("search")

	persons, err := ph.logic.GetAllPerson(limit, offset, search)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Ошибка при получении persons")
	}
	return c.JSON(http.StatusOK, persons)
}

// GetPerson обрабатывает GET запрос для получения одной записи из базы
func (ph *PersonHandlers) GetPerson(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Неверный ID")
	}
	person, err := ph.logic.GetPerson(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Ошbбка при получении Person")
	}
	return c.JSON(http.StatusOK, person)
}

// CreatePerson обрабатывает POST запрос для создания новой записи модели Person
func (ph *PersonHandlers) CreatePerson(c echo.Context) error {
	var person app.Person
	if err := c.Bind(&person); err != nil {
		return c.JSON(http.StatusBadRequest, "Неверный формат запроса")
	}
	if err := ph.logic.CreatePerson(&person); err != nil {
		return c.JSON(http.StatusInternalServerError, "Ошибка при создании Person")
	}
	return c.JSON(http.StatusCreated, "Person успешно создан")
}

// UpdatePerson выполняет Put запрос дя обновления сущ. записи в базе
func (ph *PersonHandlers) UpdatePerson(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Неверный ID")
	}
	var person app.Person
	if err := c.Bind(&person); err != nil {
		return c.JSON(http.StatusBadRequest, "Неверный форат запроса")
	}
	if err := ph.logic.UpdatePerson(id, &person); err != nil {
		return c.JSON(http.StatusInternalServerError, " Ошибка при обновлении Person")
	}
	return c.JSON(http.StatusOK, "Person успешно обновлен")
}

// DeletePerson выполнят DELETE запрос для удаления сущ. записи в базе
func (ph *PersonHandlers) DeletePerson(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Неверный Id")
	}
	if err := ph.logic.DeletePerson(id); err != nil {
		return c.JSON(http.StatusInternalServerError, "Ошибка при удалении Person")
	}
	return c.JSON(http.StatusOK, "PERSON успешно удален")
}
