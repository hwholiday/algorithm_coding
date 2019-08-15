package main

import (
	"testing"
)

func TestGetUserScore(t *testing.T) {
	t.Log(GetUserScore("/home/go/src/algorithm_coding/recommend_code/ml-100k/u.data"))
}

func TestGetUserAndMovieRating(t *testing.T) {
	data := GetUserScore("/home/go/src/algorithm_coding/recommend_code/ml-100k/u.data")
	userScore, MovieScore := GetUserAndMovieRating(data)
	t.Log(userScore[196])
	t.Log(MovieScore[242])
}

func TestGetNeighborByUid(t *testing.T) {
	data := GetUserScore("/home/go/src/algorithm_coding/recommend_code/ml-100k/u.data")
	userScore, MovieScore := GetUserAndMovieRating(data)
	t.Log(GetNeighborByUid(userScore, MovieScore, 196, 50))
}

func TestGetNeighborScoreByUid(t *testing.T) {
	data := GetUserScore("/home/go/src/algorithm_coding/recommend_code/ml-100k/u.data")
	userScore, MovieScore := GetUserAndMovieRating(data)
	fidS := GetNeighborByUid(userScore, MovieScore, 196, 50)
	list := GetNeighborScoreByUid(userScore, fidS, 196)
	t.Log("list", list)
}

func TestGetRecommendList(t *testing.T) {
	data := GetUserScore("/home/go/src/algorithm_coding/recommend_code/ml-100k/u.data")
	userScore, MovieScore := GetUserAndMovieRating(data)
	fidS := GetNeighborByUid(userScore, MovieScore, 196, 50)
	list := GetNeighborScoreByUid(userScore, fidS, 196)
	t.Log(GetRecommendList(list, userScore))
}

func TestUserCfByCosine(t *testing.T) {
	UserCfByCosine("/home/go/src/algorithm_coding/recommend_code/ml-100k/u.data",
		"/home/ghost/go/src/algorithm_coding/recommend_code/ml-100k/u.item", 186, 50)
}
