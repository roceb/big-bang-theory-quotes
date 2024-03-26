package queries

import (
	"github.com/jeyren95/big-bang-theory-quotes/db"
	"github.com/jeyren95/big-bang-theory-quotes/models"
)

func InsertQuote(quote models.Quote) error {
	statement, err := db.DB.Prepare("INSERT INTO quote(quote, character, season, episode, title) VALUES($1, $2, $3, $4, $5)")

	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(quote.Quote, quote.Character, quote.Season, quote.Episode, quote.Title)

	if err != nil {
		return err
	}

	return nil
}

func GetAllQuotes() ([]*models.Quote, error) {
	statement, err := db.DB.Prepare("SELECT * FROM quote")

	if err != nil {
		return nil, err
	}
	defer statement.Close()

	rows, err := statement.Query()

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quotes []*models.Quote

	for rows.Next() {
		quote := new(models.Quote)
		err = rows.Scan(&quote.Id, &quote.Quote, &quote.Character, &quote.Season, &quote.Episode, &quote.Title)

		if err != nil {
			return nil, err
		}

		quotes = append(quotes, quote)
	}

	return quotes, nil
}
