package users

import (
	"github.com/creasty/defaults"
	"github.com/rs/zerolog/log"
	"strings"
	"xmasx/internal/data"
	"xmasx/internal/xlist"
)

type Users []User

func (p Users) name() {

}

func (p *Users) Load(ds data.ListStore) {
	data := ds.Load()

	var list xlist.XList
	for _, row := range data {
		p := User{}
		if err := defaults.Set(&p); err != nil {
			log.Error().Err(err).Str("row", strings.Join(row, " ")).Msg("Failed to set defaults")
		}
		p.Name = row[0]
		p.Group = row[1]

		list.AddGifter(xlist.Gifter{User: p})
	}
}
