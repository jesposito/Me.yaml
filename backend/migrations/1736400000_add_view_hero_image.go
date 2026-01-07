package migrations

import (
	"github.com/pocketbase/pocketbase/core"
)

func init() {
	Migrations = append(Migrations, Migration{
		Name: "1736400000_add_view_hero_image",
		Up: func(app core.App) error {
			collection, err := app.FindCollectionByNameOrId("views")
			if err != nil {
				return err
			}

			if collection.Fields.GetByName("hero_image") != nil {
				return nil
			}

			collection.Fields.Add(&core.FileField{
				Name:      "hero_image",
				MaxSize:   10485760,
				MimeTypes: []string{"image/jpeg", "image/png", "image/webp"},
			})

			return app.Save(collection)
		},
		Down: func(app core.App) error {
			collection, err := app.FindCollectionByNameOrId("views")
			if err != nil {
				return err
			}

			field := collection.Fields.GetByName("hero_image")
			if field != nil {
				collection.Fields.RemoveById(field.GetId())
				return app.Save(collection)
			}

			return nil
		},
	})
}
