package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

//基于用户的推荐(通过共同口味与偏好找相似邻居用户，K-邻居算法，你朋友喜欢，你也可能喜欢)，
func UserCfByCosine(path, movie string, uid int, k int) {
	data := GetUserScore(path)
	userScore, MovieScore := GetUserAndMovieRating(data)
	fidS := GetNeighborByUid(userScore, MovieScore, uid, k)
	list := GetNeighborScoreByUid(userScore, fidS, uid)
	recommendMovies := GetRecommendList(list, userScore)
	movies := GetMoviesList(movie)
	for _, v := range recommendMovies {
		fmt.Println("电影ID :", v.Fid, "得分 :", v.Rating, "电影信息", movies[v.Fid])
	}
}

//获取电影列表
func GetMoviesList(path string) map[int]string {
	info := ReadLine(path)
	var dataMap = make(map[int]string, len(info))
	for _, v := range info {
		val := strings.Split(string(v), "|")
		intIndex, err := strconv.Atoi(val[0])
		CheckError(err)
		dataMap[intIndex] = v[1:]
	}
	return dataMap
}

//把用户电影评分放入集合
func GetUserScore(path string) (data []UserScore) {
	//user id | item id | rating | timestamp
	userScoreList := ReadLine(path)
	for _, v := range userScoreList {
		val := strings.Split(v, "\t")
		add := UserScore{}
		var err error
		add.Uid, err = strconv.Atoi(val[0])
		CheckError(err)
		add.Mid, err = strconv.Atoi(val[1])
		CheckError(err)
		add.Rating, err = strconv.Atoi(val[2])
		CheckError(err)
		data = append(data, add)
	}
	return
}

// UserRatingMap 获取用户对那些电影拼分过
// MovieRatingMap 获取该电影那些人评分过
func GetUserAndMovieRating(data []UserScore) (map[int][]MovieScore, map[int][]UidScore) {
	var (
		UserRatingMap  = make(map[int][]MovieScore)
		MovieRatingMap = make(map[int][]UidScore)
	)
	for _, v := range data {
		if _, ok := UserRatingMap[v.Uid]; ok {
			UserRatingMap[v.Uid] = append(UserRatingMap[v.Uid], MovieScore{
				Mid:    v.Mid,
				Rating: v.Rating,
			})
		} else {
			UserRatingMap[v.Uid] = []MovieScore{{
				Mid:    v.Mid,
				Rating: v.Rating,
			}}
		}
		if _, ok := MovieRatingMap[v.Mid]; ok {
			MovieRatingMap[v.Mid] = append(MovieRatingMap[v.Mid], UidScore{
				Uid:    v.Uid,
				Rating: v.Rating,
			})
		} else {
			MovieRatingMap[v.Uid] = []UidScore{{
				Uid:    v.Uid,
				Rating: v.Rating,
			}}
		}
	}
	return UserRatingMap, MovieRatingMap
}

//获取与我有相同兴趣爱好的人（关注过相同的电影）
//k 查多少个与我有相同爱好的
//我喜欢的电影 我评分的总分/我评分的总数  大于等于平均我喜欢
func GetNeighborByUid(userScore map[int][]MovieScore, movieUser map[int][]UidScore, uid, k int) (neighbor []int) {
	if _, ok := userScore[uid]; !ok {
		panic("not find uid")
	}
	var endFor = false
	rating := GetLikeMovieRating(userScore, uid)
	for _, v := range userScore[uid] { //查看我看了那些电影
		if v.Rating < rating {
			continue
		}
		for _, v1 := range movieUser[v.Mid] {
			if v1.Uid != uid && !CheckInArray(neighbor, v1.Uid) {
				if v1.Rating >= GetLikeMovieRating(userScore, v1.Uid) {
					neighbor = append(neighbor, v1.Uid)
					if len(neighbor) >= k { //只去与你有相同爱好的K个人
						endFor = true
						break
					}
				}
			}
		}
		if endFor {
			break
		}
	}
	return
}

//获取与我有相同兴趣的人的相似度
func GetNeighborScoreByUid(userScore map[int][]MovieScore, neighbor []int, uid int) (data []FidScore) {
	//整理数据 计算相似度
	for _, fid := range neighbor {
		//我与我的好友之间对我们合集电影的打分
		var movieMap = make(map[int]UidAndFirScore)
		for _, v := range userScore[uid] {
			if val, ok := movieMap[v.Mid]; ok {
				val.UidRating += v.Rating
				movieMap[v.Mid] = val
			} else {
				movieMap[v.Mid] = UidAndFirScore{
					UidRating: v.Rating,
					FidRating: 0,
				}
			}
		}
		for _, v := range userScore[fid] {
			if val, ok := movieMap[v.Mid]; ok {
				val.FidRating += v.Rating
				movieMap[v.Mid] = val
			} else {
				movieMap[v.Mid] = UidAndFirScore{
					UidRating: 0,
					FidRating: v.Rating,
				}
			}
		}
		var (
			uidLike []float64
			fidLike []float64
		)
		//计算相似度
		for _, v := range movieMap {
			uidLike = append(uidLike, float64(v.UidRating))
			fidLike = append(fidLike, float64(v.FidRating))
		}
		data = append(data, FidScore{
			Fid:    fid,
			Rating: Cosine(uidLike, fidLike),
		})
	}
	array := FidScoreArray{}
	array = data
	sort.Sort(array)
	return array
}

func GetRecommendList(list []FidScore, userScore map[int][]MovieScore) []FidScore {
	//计算推荐列表(计算电影在每个朋友哪里的评分，然后根据每个电影的评分给用户推荐数据)
	var recommendListMap = make(map[int]float64)
	for _, v := range list {
		for _, v1 := range userScore[v.Fid] {
			if _, ok := recommendListMap[v1.Mid]; ok {
				recommendListMap[v1.Mid] += v.Rating
			} else {
				recommendListMap[v1.Mid] = v.Rating
			}
		}
	}
	//对推荐的电影进行排序
	var recommendList []FidScore
	//这里将FID当成MID使用
	for key, v := range recommendListMap {
		recommendList = append(recommendList, FidScore{
			Fid:    key, //MID
			Rating: v,   //的分
		})
	}
	array := FidScoreArray{}
	array = recommendList
	sort.Sort(array)
	return array
}

//计算评分的平均数
func GetLikeMovieRating(userScore map[int][]MovieScore, uid int) int {
	if _, ok := userScore[uid]; !ok {
		panic("not find uid")
	}
	var sum int
	for _, v := range userScore[uid] {
		sum += v.Rating
	}
	return sum / len(userScore)
}

//向量空间余弦相似度(Cosine Similarity)
func Cosine(a []float64, b []float64) float64 {
	var (
		aLen  = len(a)
		bLen  = len(b)
		sum   = 0.0
		sa    = 0.0
		sb    = 0.0
		count = 0
	)
	if aLen > bLen {
		count = aLen
	} else {
		count = bLen
	}
	for i := 0; i < count; i++ {
		if i >= aLen {
			sb += math.Pow(b[i], 2)
			continue
		}
		if i >= bLen {
			sa += math.Pow(a[i], 2)
			continue
		}
		sum += a[i] * b[i]
		sa += math.Pow(a[i], 2)
		sb += math.Pow(b[i], 2)
	}
	return sum / (math.Sqrt(sa) * math.Sqrt(sb))
}
