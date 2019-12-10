package consistent_hash_code

import (
	"errors"
	"fmt"
	"hash/crc32"
	"sort"
)

type Consistent struct {
	nodesReplicas int               //添加虚拟节点数量
	hashSortNodes []uint32          //所有节点的排序数组
	circle        map[uint32]string //所有节点对应的node
	nodes         map[string]bool   //node是否存在
}

func NewConsistent() *Consistent {
	return &Consistent{
		nodesReplicas: 20,
		circle:        make(map[uint32]string),
		nodes:         make(map[string]bool),
	}
}

func (c *Consistent) Add(node string) error {
	if _, ok := c.nodes[node]; ok { //判断新加节点是否存在
		return fmt.Errorf("%s already existed", node)
	}
	c.nodes[node] = true
	for i := 0; i < c.nodesReplicas; i++ { //添加虚拟节点
		replicasKey := getReplicasKey(i, node) //虚拟节点
		c.circle[replicasKey] = node
		c.hashSortNodes = append(c.hashSortNodes, replicasKey) //所有节点ID
	}
	sort.Slice(c.hashSortNodes, func(i, j int) bool { //排序
		return c.hashSortNodes[i] < c.hashSortNodes[j]
	})
	return nil
}

func (c *Consistent) Remove(node string) error {
	if _, ok := c.nodes[node]; !ok { //判断新加节点是否存在
		return fmt.Errorf("%s not existed", node)
	}
	delete(c.nodes, node)
	for i := 0; i < c.nodesReplicas; i++ {
		replicasKey := getReplicasKey(i, node)
		delete(c.circle, replicasKey) //删除虚拟节点
	}
	c.refreshHashSortNodes()
	return nil
}
func (c *Consistent) GetNode() (node []string) {
	for v := range c.nodes {
		node = append(node, v)
	}
	return node
}

func (c *Consistent) Get(key string) (string, error) {
	if len(c.nodes) == 0 {
		return "", errors.New("not add node")
	}
	index := c.searchNearbyIndex(key)
	host := c.circle[c.hashSortNodes[index]]
	return host, nil
}

func (c *Consistent) refreshHashSortNodes() {
	c.hashSortNodes = nil
	for v := range c.circle {
		c.hashSortNodes = append(c.hashSortNodes, v)
	}
	sort.Slice(c.hashSortNodes, func(i, j int) bool { //排序
		return c.hashSortNodes[i] < c.hashSortNodes[j]
	})
}

func (c *Consistent) searchNearbyIndex(key string) int {
	hashKey := hashKey(key)
	index := sort.Search(len(c.hashSortNodes), func(i int) bool { //key算出的节点，距离最近的node节点下标  sort.Search数组需要升序排列好
		return c.hashSortNodes[i] >= hashKey
	})
	if index >= len(c.hashSortNodes) {
		index = 0
	}
	return index
}

func getReplicasKey(i int, node string) uint32 {
	return hashKey(fmt.Sprintf("%s#%d", node, i))
}

func hashKey(host string) uint32 {
	return crc32.ChecksumIEEE([]byte(host))
}
