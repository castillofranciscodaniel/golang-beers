package domain

import (
	"database/sql"
	"github.com/castillofranciscodaniel/golang-beers/infrastructure/err"
)

type Beer struct {
	id       int64
	name     string
	brewery  string
	country  string
	price    float64
	currency string
}

func NewBeer(id int64, name string, brewery string, country string, price float64, currency string) (Beer, error) {
	if id < 1 {
		return Beer{}, err.IdCanNotBeMinorThanOneError
	}

	if price < 0 {
		return Beer{}, err.PriceCanNotBeMinorThanZeroError
	}

	return Beer{
		id:       id,
		name:     name,
		brewery:  brewery,
		country:  country,
		price:    price,
		currency: currency,
	}, nil
}

func (b *Beer) GetId() int64 {
	return b.id
}

func (b *Beer) GetName() string {
	return b.name
}

func (b *Beer) GetCountry() string {
	return b.country
}

func (b *Beer) GetPrice() float64 {
	return b.price
}

func (b *Beer) GetCurrency() string {
	return b.currency
}

func (b *Beer) GetBrewery() string {
	return b.brewery
}

type BeerSql struct {
	Id       sql.NullInt64
	Name     sql.NullString
	Brewery  sql.NullString
	Country  sql.NullString
	Price    sql.NullFloat64
	Currency sql.NullString
}

func (b *BeerSql) MapToDomain() (Beer, error) {
	var id int64
	var name string
	var brewery string
	var country string
	var price float64
	var currency string

	if b.Id.Valid {
		id = b.Id.Int64
	}
	if b.Name.Valid {
		name = b.Name.String
	}
	if b.Brewery.Valid {
		brewery = b.Brewery.String
	}
	if b.Country.Valid {
		country = b.Country.String
	}
	if b.Price.Valid {
		price = b.Price.Float64
	}
	if b.Currency.Valid {
		currency = b.Currency.String
	}

	return NewBeer(id, name, brewery, country, price, currency)
}

//go:generate mockgen -destination=./mockBeerRepository.go -package=domain github.com/castillofranciscodaniel/golang-beers/domain BeerRepository
type BeerRepository interface {
	Get() ([]BeerSql, error)
	Post(beer Beer) error
	GetById(id int64) (*BeerSql, error)
}
