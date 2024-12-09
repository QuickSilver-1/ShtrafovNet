package entity

import (
	derr "auction/internal/domain/errors"
	"encoding/json"
)

type Lot struct {
    id     int     // идентификатор лота
    name   string  // название лота
    owner  *User   // владелец лота
    params string  // параметры лота в формате JSON
}

// LotParams - параметры лота
type LotParams struct {
    Description string  `json:"desc"`  // описание лота
    MinPrice    float64 `json:"min"`   // минимальная цена лота
}

func CreateLot(id int, name, desc string, owner *User, minPrice float64) (*Lot, error) {
    // Проверки валидности
    if name == "" {
        return nil, &derr.NonValidData{
            Err:  "Invalid name",
            Code: 400,
        }
    }

    if owner.Id() == 0 {
        return nil, &derr.NonValidData{
            Err:  "Invalid owner value",
            Code: 400,
        }
    }

    if minPrice == 0 {
        return nil, &derr.NonValidData{
            Err:  "Invalid minimal price value",
            Code: 400,
        }
    }

    // Формирование параметров лота
    params := &LotParams{
        Description: desc,
        MinPrice:    minPrice,
    }

    jsonParams, err := json.Marshal(params)
    if err != nil {
        return nil, &derr.JsonCodeError{
            Err:  "Json code error",
            Code: 500,
        }
    }

    return &Lot{
        id:     id,
        name:   name,
        owner:  owner,
        params: string(jsonParams),
    }, nil
}

// Геттеры для полей класса
func (lot *Lot) Id() int { return lot.id }

func (lot *Lot) Name() string { return lot.name }

func (lot *Lot) Owner() *User { return lot.owner }

func (lot *Lot) Params() string { return lot.params }

func (lot *Lot) MinPrice() float64 {
    var params *LotParams

    json.Unmarshal([]byte(lot.params), &params)
    return params.MinPrice
}