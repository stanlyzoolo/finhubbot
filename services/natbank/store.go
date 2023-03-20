package natbank

import (
	"context"
	"database/sql"

	"github.com/stanlyzoolo/smartLaFamiliaBot/storage/dbo"
)

func (c *service) storeRates(ctx context.Context, rates []Rate) error {
	for _, rate := range rates {
		if err := c.storage.Create(ctx, dbo.NatBankRate{
			ID: rate.ID,
			Abbreviation: sql.NullString{
				String: rate.Abbreviation,
				Valid:  len(rate.Abbreviation) != 0,
			},
			Name: sql.NullString{
				String: rate.Name,
				Valid:  len(rate.Name) != 0,
			},
			Scale:        rate.Scale,
			OfficialRate: rate.OfficialRate,
		}); err != nil {
			return err
		}
	}

	return nil
}
