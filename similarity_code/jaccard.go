package similarity_code

//两个集合A和B交集元素的个数在A、B并集中所占的比例，称为这两个集合的杰卡德系数(0~1之间)
//k-Shingles
//一篇文档可以看成是一个字符串，文档的k-shingle为在该文档中长度为k的所有子串。任意一篇文档都可以表示为k-shingles的集合，比如“A document is a string of characters”这句话的所有3-shingles为{ ”A d”, ”do”, ”doc”, ”ocu”, ”cum”, ”ume”, ”men”, ”ent”, . . . , ”ers” }。如果k非常小，那么k个字符的序列会出现在大多数的文档中，如k=1，许多文档都有相同的字符，几乎所有的文档都有很高的相似性。如果k应该足够大，那么对于给定的shingle出现在不同的文档中的概率是非常低的。比如这两个单词“ document”和“monumen
type JacCard struct {
	A    string
	B    string
	K    int
	a    map[string]string
	b    map[string]string
	aLen int
	bLen int
}

func NewJacCard(a, b string, k int) *JacCard {
	return &JacCard{
		A:    a,
		B:    b,
		K:    k,
		aLen: len(a),
		bLen: len(b),
	}
}

func (j *JacCard) kShingles() {
	j.a = make(map[string]string, j.aLen/j.K+1)
	j.b = make(map[string]string, j.bLen/j.K+1)
	for i := 0; i < j.aLen; i = i + j.K {
		end := i + j.K
		if end > j.aLen {
			end = j.aLen
		}
		j.a[j.A[i:end]] = ""
	}
	for i := 0; i < j.bLen; i = i + j.K {
		end := i + j.K
		if end > j.bLen {
			end = j.bLen
		}
		j.b[j.B[i:end]] = ""
	}
}

func (j *JacCard) King() float64 {
	var (
		Intersection int
		Union        int
		king         float64
	)
	//计算交集
	for k, _ := range j.a {
		if _, ok := j.b[k]; ok {
			Intersection++
		}
	}
	//计算并集
	Union = len(j.a) + len(j.b) - Intersection
	//得出相似度
	king = float64(Intersection) / float64(Union)
	return king
}
