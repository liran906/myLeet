// https://leetcode.cn/problems/insert-delete-getrandom-o1/description/?envType=study-plan-v2&envId=top-interview-150

package hot150

import "math/rand"

type RandomizedSet struct {
	data map[int]int // key : index
	keys []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		data: make(map[int]int),
		keys: []int{},
	}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, exists := rs.data[val]; exists {
		return false
	}
	rs.keys = append(rs.keys, val)
	rs.data[val] = len(rs.keys) - 1
	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	if _, exists := rs.data[val]; !exists {
		return false
	}
	pos := rs.data[val]
	last := len(rs.keys) - 1
	rs.data[rs.keys[last]] = pos                              // update new position
	rs.keys[pos], rs.keys[last] = rs.keys[last], rs.keys[pos] // swap
	rs.keys = rs.keys[:last]

	delete(rs.data, val)
	return true
}

func (rs *RandomizedSet) GetRandom() int {
	n := len(rs.keys)
	rndNum := rand.Intn(n)
	return rs.keys[rndNum]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
