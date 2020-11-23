package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"strings"
	"time"
)

type GetUserRequest struct {
	Id int64 `form:"id" binding:"required"`
}

type ListUserRequest struct {
	Offset int `form:"offset"`
	Limit  int `form:"limit" binding:"required"`
}

type User struct {
	Id        int64         `json:"id" db:"id"`
	Name      string        `json:"name" db:"name" binding:"required"`
	Birthday  JsonBirthDate `json:"birthday" db:"birthday" binding:"required"`
	CreatedAt *time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time    `json:"updated_at" db:"updated_at"`
}

type JsonBirthDate time.Time

func (j *JsonBirthDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = JsonBirthDate(t)
	return nil
}

func (j JsonBirthDate) MarshalJSON() ([]byte, error) {
	t := time.Time(j)
	return json.Marshal(t)
}

// Maybe a Format function for printing your date
func (j JsonBirthDate) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}

func (j JsonBirthDate) Time() (time.Time, error) {
	if j != JsonBirthDate(time.Time{}) {
		return time.Time(j), nil
	}
	return time.Time{}, errors.New("JsonBirthDate is empty")
}

func (j JsonBirthDate) ConvertValue(v interface{}) (driver.Value, error) {
	if j != JsonBirthDate(time.Time{}) {
		return time.Time(j), nil
	}
	return time.Time{}, errors.New("JsonBirthDate is empty")
}

func (j JsonBirthDate) Value() (driver.Value, error) {
	if j != JsonBirthDate(time.Time{}) {
		return time.Time(j), nil
	}
	return time.Time{}, errors.New("JsonBirthDate is empty")
}
