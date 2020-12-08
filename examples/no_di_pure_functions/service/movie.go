package service

func GetActorMovieCategories(name string) ([]string, error) {
	movies, err := getByActorName(name)
	if err != nil {
		return nil, err
	}

	var categories []string
	for _, m := range movies {
		categories = append(categories, m.Category)

		//votes := third_party.ActorVotesByCategory(name, m.Category)
		//categoryVotes := fmt.Sprintf("%s-%v", m.Category, votes)
		//categories = append(categories, categoryVotes)
	}

	return categories, nil
}
