package repository

import domain "mercadolibre.com/di/examples/with_di"

type RActor struct {
}

func NewRActor() *RActor {
	return &RActor{}
}

func (ra *RActor) FindAllMovies() []domain.Movie {
	// mocked data
	return []domain.Movie{
		{
			Name:     "Terminator",
			Category: "action",
			Cast: map[string]interface{}{
				"arnold":    struct{}{},
				"silvester": struct{}{},
				"clint":     struct{}{},
			},
		},
		{
			Name:     "Dealing with shipping transitions!",
			Category: "horror",
			Cast: map[string]interface{}{
				"john":      struct{}{},
				"silvester": struct{}{},
				"sarah":     struct{}{},
			},
		},
		{
			Name:     "My life at Meli :D",
			Category: "thriller",
			Cast: map[string]interface{}{
				"frank":     struct{}{},
				"johnathan": struct{}{},
				"lisa":      struct{}{},
			},
		},
	}
}
