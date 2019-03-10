package problem0947

func removeStones(stones [][]int) int {
	size := len(stones)

	u := newUnion(size)
	for i := 0; i < size; i++ {
		ix, iy := stones[i][0], stones[i][1]
		for j := i + 1; j < size; j++ {
			jx, jy := stones[j][0], stones[j][1]
			if ix == jx || iy == jy {
				u.union(i, j)
			}
		}
	}

	return u.res
}

// Robert Sedgewick 算法（第4版） 1.5.2.7
// union-find (加权 quick-union)，还作了路径压缩优化

// union is ...
type union struct {
	id   []int // 父链接数组(由触点索引)
	size []int // (由触点索引的) 各个根节点所对应的分量的大小
	res  int   // NOTICE: 题目需要的结果
}

func newUnion(N int) *union {
	id := make([]int, N)
	for i := range id {
		id[i] = i
	}
	sz := make([]int, N)
	for i := range sz {
		sz[i] = 1
	}
	return &union{
		id:   id,
		size: sz,
		res:  0,
	}
}

func (u *union) find(p int) int {
	// 跟随连接找到根节点
	for p != u.id[p] {
		p = u.id[p]
	}
	return p
}

func (u *union) union(p, q int) {
	i, j := u.find(p), u.find(q)
	if i == j {
		return
	}
	if u.size[i] > u.size[j] {
		i, j = j, i
	}
	// 将小树的根节点连接到大树的根节点
	u.id[i] = j
	u.size[j] += u.size[i]
	u.res++
	return
}
