package v1

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type Score struct {
	uid   int
	mid   int
	score int
}

type ScoreArray []Score

func (s ScoreArray) Len() int {
	return len(s)
}

func (s ScoreArray) Less(i, j int) bool {
	return s[i].mid < s[j].mid
}

func (s ScoreArray) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type FidScore struct {
	fid   int
	score float64
}

type FidScoreArray []FidScore

func (s FidScoreArray) Len() int {
	return len(s)
}

func (s FidScoreArray) Less(i, j int) bool {
	return s[i].score > s[j].score
}

func (s FidScoreArray) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type MidScore struct {
	mid   int
	score float64
}

type MidScoreArray []MidScore

func (s MidScoreArray) Len() int {
	return len(s)
}

func (s MidScoreArray) Less(i, j int) bool {
	return s[i].score > s[j].score
}

func (s MidScoreArray) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	movies := getMoviesList("/home/go/src/recommend_code/ml-100k/u.item")
	recommendList := recommendByUserCF("/home/go/src/recommend_code/ml-100k/u.data", 456, 80)
	//开始计算推荐(推荐5条数据)
	for _, moviesId := range recommendList[:5] {
		fmt.Println(movies[moviesId.mid])
	}
}

func recommendByUserCF(path string, uid int, k int) []MidScore {
	data := FileReadLine(path)
	userScore, itemUser := getScoreList(data)
	//寻找有相同兴趣的并排序(从高到低)
	info := calcNearestNeighbor(uid, k, userScore, itemUser)
	var reconnebdMap = make(map[int]float64)
	for _, v := range info {
		for _, v1 := range userScore[v.fid] { //查询好友对那些电影评分
			if _, ok := reconnebdMap[v1.mid]; ok {
				reconnebdMap[v1.mid] = v.score
			} else {
				reconnebdMap[v1.mid] += v.score
			}
		}
	}
	var dataArray MidScoreArray
	for k, v := range reconnebdMap {
		dataArray = append(dataArray, MidScore{
			mid:   k,
			score: v,
		})
	}
	sort.Sort(dataArray)
	return dataArray
}

//寻找与你有相同爱好的
func calcNearestNeighbor(uid, k int, userScore map[int][]Score, itemUser map[int][]int) FidScoreArray {
	var neighbors []int
	var endFor = false
	for _, vUserScore := range userScore[uid] { //该uid对那些电影打过分
		for _, vItemUser := range itemUser[vUserScore.mid] { //查询该电影那些人评分过
			if vItemUser != uid && !containsDuplicate(neighbors, vItemUser) {
				neighbors = append(neighbors, vItemUser)
				if len(neighbors) >= k { //只去与你有相同爱好的K个人
					endFor = true
					break
				}
			}
		}
		if endFor {
			break
		}
	}
	//neighbors 那些与你关注的相同的电影
	myUids := getUserScoreString(userScore, uid)
	var endData = FidScoreArray{}
	for _, val := range neighbors {
		neighborsVal := getUserScoreString(userScore, val)
		var (
			Intersection int
			Union        int
			king         float64
		)
		//计算交集
		for _, vUid := range myUids {
			for _, vNeighbors := range neighborsVal {
				if vUid == vNeighbors {
					Intersection++
				}
			}
		}
		//计算并集
		Union = len(myUids) + len(neighborsVal) - Intersection
		//得出相似度
		king = Decimal(float64(Intersection) / float64(Union))
		endData = append(endData, FidScore{
			fid:   val,
			score: king,
		})
	}
	sort.Sort(endData)
	return endData
}

func getUserScoreString(data map[int][]Score, uid int) []string {
	val := ScoreArray{}
	val = data[uid]
	sort.Sort(val)
	var dataString []string
	var Average int
	for _, v := range val {
		Average += v.score
	}
	Average = Average / len(val)
	for _, v := range val {
		//计算用户大于该平均分，认为用户喜欢该推荐
		if v.score >= Average {
			dataString = append(dataString, fmt.Sprintf("%d", v.mid))
		}
	}
	return dataString
}

//用户id,电影id,打分,时间
func getScoreList(data []string) (map[int][]Score, map[int][]int) {
	var info = make([]Score, len(data))
	for index, line := range data {
		rate := strings.Split(line, "\t")
		intUid, err := strconv.Atoi(rate[0])
		CheckErr(err)
		intMid, err := strconv.Atoi(rate[1])
		CheckErr(err)
		intScore, err := strconv.Atoi(rate[2])
		CheckErr(err)
		info[index] = Score{
			uid:   intUid,
			mid:   intMid,
			score: intScore,
		}
	}
	////用户id,电影id,打分,时间
	var userScore = make(map[int][]Score) //用户对那些电影评分了
	var itemUser = make(map[int][]int)    //那些人对这个电影评分了
	for _, v := range info {
		if _, ok := userScore[v.uid]; ok {
			userScore[v.uid] = append(userScore[v.uid], v)
		} else {
			userScore[v.uid] = []Score{v}
		}
		if _, ok := itemUser[v.mid]; ok {
			itemUser[v.mid] = append(itemUser[v.mid], v.uid)
		} else {
			itemUser[v.mid] = []int{v.uid}
		}
	}
	return userScore, itemUser
}

//获取电影列表
func getMoviesList(path string) map[int]string {
	info := FileReadLine(path)
	var dataMap = make(map[int]string, len(info))
	for _, v := range info {
		val := strings.Split(string(v), "|")
		intIndex, err := strconv.Atoi(val[0])
		CheckErr(err)
		dataMap[intIndex] = v[1:]
	}
	return dataMap
}

func FileReadLine(path string) []string {
	data, err := ioutil.ReadFile(path)
	CheckErr(err)
	var info []string
	reader := bufio.NewReader(bytes.NewReader(data))
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		info = append(info, line)
	}
	return info
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func containsDuplicate(nums []int, index int) bool {
	for _, val := range nums {
		if val == index {
			return true
		}
	}
	return false
}

func Decimal(value float64) float64 {
	value, err := strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	CheckErr(err)
	return value
}
