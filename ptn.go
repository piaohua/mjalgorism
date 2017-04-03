package algo

import (
	"basic/utils"
	"fmt"
	"reflect"
)

//排列组合
// a = [[1],[2],[1]]
// [[[1],[2],[1]], [[1],[1],[2]], [[2],[1],[1]]]
func perms(a [][]int) [][][]int {
	var b [][][]int = [][][]int{}
	var c [][]int = [][]int{}
	perms_f(c, a, &b)
	//fmt.Println(b)
	//perms_(a, 0)
	return uniq(b)
}

//排列组合
func perms_(a [][]int, index int) {
	var lenght int = len(a)
	if index == lenght {
		//fmt.Println("a:", a)
	}
	for i := index; i < lenght; i++ {
		a[index], a[i] = a[i], a[index]
		perms_(a, index+1)
		a[index], a[i] = a[i], a[index]
	}
}

// perms_f({}, {1,2,3}, &c)
func perms_f(a [][]int, b [][]int, c *[][][]int) [][]int {
	var lenght int = len(b)
	if lenght == 1 {
		//fmt.Println(append(a, b...))
		return append(a, b...)
	}
	for i := 0; i < lenght; i++ {
		var tmpA [][]int = [][]int{}
		var tmpB [][]int = [][]int{}
		tmpA = append(tmpA, a...)
		tmpA = append(tmpA, b[i]) //添加当前值
		tmpB = append(tmpB, b[:i]...)
		tmpB = append(tmpB, b[i+1:]...) //移除当前值
		r := perms_f(tmpA, tmpB, c)
		if r != nil {
			*c = append(*c, r)
		}
	}
	return nil
}

//去重
func uniq(a [][][]int) [][][]int {
	if len(a) == 0 {
		return a
	}
	var t [][][]int = [][][]int{}
	for k, i := range a {
		if k == 0 {
			t = append(t, i)
		}
		var ok bool = false
		for _, j := range t {
			if reflect.DeepEqual(i, j) {
				ok = true
				break
			}
		}
		if !ok {
			t = append(t, i)
		}
	}
	return t
}

//去重
func uniq_(a []int) []int {
	if len(a) == 0 {
		return a
	}
	var t []int = []int{}
	for k, i := range a {
		if k == 0 {
			t = append(t, i)
		}
		var ok bool = false
		for _, j := range t {
			if i == j {
				ok = true
				break
			}
		}
		if !ok {
			t = append(t, i)
		}
	}
	return t
}

//扁平化处理
func flatten(a [][]int) []int {
	var t []int = []int{}
	for _, i := range a {
		for _, j := range i {
			t = append(t, j)
		}
	}
	return t
}

//数值和
func total(a []int) int {
	var t int = 0
	for _, i := range a {
		t += i
	}
	return t
}

//a : [[1,1,1], [1,1,1], [1,1,1], [1,1,1], [2]]
//a : [[2], [3]] -> [[[2], [3]], [[3], [2]], [[3], [2]], [[2], [3]]]
//a : [[2], [4]] -> [[[2], [2]], [[2], [2]], [[4]]]
func ptn(a [][]int) (ret [][][]int) {
	var size int = len(a)
	if size == 1 {
		return append(ret, a)
	}
	ret = perms(a)
	h1 := make(map[string]bool)
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			var key []int = []int{}
			var zero []int = []int{0}
			key = append(key, a[i]...)
			key = append(key, zero...)
			key = append(key, a[j]...)
			key_s := ""
			for _, k_i := range key {
				key_s += utils.Itoa(k_i)
			}
			//fmt.Println("key -> ", key, "key_s -> ", key_s)
			if _, ok := h1[key_s]; !ok {
				h1[key_s] = true
				h2 := make(map[string]bool)
				//a[i]和a[j]范围
				k_len := len(a[i]) + len(a[j])
				for k := 0; k <= k_len; k++ {
					var t1 []int = []int{}
					for t1_i := 0; t1_i < len(a[j]); t1_i++ {
						t1 = append(t1, 0)
					}
					var t []int = []int{}
					t = append(t, t1...)
					t = append(t, a[i]...)
					t = append(t, t1...)
					//fmt.Println("t -> ", t, k)
					m_len := len(a[j])
					for m := 0; m < m_len; m++ {
						t[k+m] += a[j][m]
					}
					//去除多余的0
					var t0 []int = []int{}
					for _, tv := range t {
						if tv != 0 {
							t0 = append(t0, tv)
						}
					}
					t = t0
					//检查是否有大于4的值
					var t_ok bool = false
					for _, tv := range t {
						if tv > 4 {
							t_ok = true
							break
						}
					}
					if t_ok {
						continue
					}
					//检查长度是否有大于9
					if len(t) > 9 {
						continue
					}
					//重复检测
					key_t := ""
					for _, t_i := range t {
						key_t += utils.Itoa(t_i)
					}
					if _, ok := h2[key_t]; !ok {
						h2[key_t] = true
						//其余
						var t5 [][]int = make([][]int, size, size)
						err := utils.Clone(&t5, &a)
						if err != nil {
							panic(err)
						}
						t4 := append(t5[:i], t5[i+1:]...)
						t2 := append(t4[:j-1], t4[j:]...)
						//递归再次调用
						var t3 [][]int = [][]int{}
						t3 = append(t3, t)
						t3 = append(t3, t2...)
						ret2 := ptn(t3)
						ret = append(ret, ret2...)
					}
				}
			}
		}
	}
	return ret
}

//键值计算
//[[1,1,1],[1,1,1],[1,1,1],[1,1,1],[2]]
func calc_key(a [][]int) int {
	var ret int = 0
	var len int = -1
	for _, b := range a {
		for _, i := range b {
			len += 1
			switch i {
			case 2:
				//ret |= 0b11 << len
				ret |= 0x3 << uint32(len)
				len += 2
			case 3:
				//ret |= 0b1111 << len
				ret |= 0xF << uint32(len)
				len += 4
			case 4:
				//ret |= 0b111111 << len
				ret |= 0x3F << uint32(len)
				len += 6
			}
		}
		//ret |= 0b1 << len
		ret |= 0x1 << uint32(len)
		len += 1
	}
	return ret
}

//值计算,返回16进制表示值
//a : [[1,1,1], [1,1,1], [1,1,1], [1,1,1], [2]]
//牌型对应占位:
//   3bit  0: 刻子の数(0～4)
//   3bit  3: 順子の数(0～4)
//   4bit  6: 頭の位置(1～13)
//   4bit 10: 面子の位置１(0～13)
//   4bit 14: 面子の位置２(0～13)
//   4bit 18: 面子の位置３(0～13)
//   4bit 22: 面子の位置４(0～13)
//   1bit 26: 七対子フラグ
//   1bit 27: 九蓮宝燈フラグ
//   1bit 28: 一気通貫フラグ
//   1bit 29: 二盃口フラグ
//   1bit 30: 一盃口フラグ
func find_hai_pos(a [][]int) []int {
	var ret_array []int = []int{}
	var ret int = 0
	var p_atama int = 0
	var a_size int = len(a)
	for i := 0; i < a_size; i++ {
		var a_i_size int = len(a[i])
		for j := 0; j < a_i_size; j++ {
			//查找头
			if a[i][j] >= 2 {
				//刻子,顺子,优先顺序替换
				for kotsu_shuntus := 0; kotsu_shuntus <= 1; kotsu_shuntus++ {
					var t [][]int = make([][]int, a_size, a_size)
					err := utils.Clone(&t, &a)
					if err != nil {
						panic(err)
					}
					t[i][j] -= 2
					var p int = 0
					var p_kotsu []int = []int{}
					var p_shuntsu []int = []int{}
					var t_size int = len(t)
					for k := 0; k < t_size; k++ {
						var t_k_size int = len(t[k])
						for m := 0; m < t_k_size; m++ {
							if kotsu_shuntus == 0 {
								//刻子优先取出
								//刻子
								if t[k][m] >= 3 {
									t[k][m] -= 3
									p_kotsu = append(p_kotsu, p)
								}
								//顺子
								for {
									if len(t[k])-m >= 3 &&
										t[k][m] >= 1 &&
										t[k][m+1] >= 1 &&
										t[k][m+2] >= 1 {
										t[k][m] -= 1
										t[k][m+1] -= 1
										t[k][m+2] -= 1
										p_shuntsu = append(p_shuntsu, p)
									} else {
										break
									}
								}
							} else {
								//顺子优先取出
								//顺子
								for {
									if len(t[k])-m >= 3 &&
										t[k][m] >= 1 &&
										t[k][m+1] >= 1 &&
										t[k][m+2] >= 1 {
										t[k][m] -= 1
										t[k][m+1] -= 1
										t[k][m+2] -= 1
										p_shuntsu = append(p_shuntsu, p)
									} else {
										break
									}
								}
								//刻子
								if t[k][m] >= 3 {
									t[k][m] -= 3
									p_kotsu = append(p_kotsu, p)
								}
							}
							p += 1
						}
					}
					//其它牌型
					t_f := flatten(t)
					var ok bool = true
					for _, t_v := range t_f {
						if t_v != 0 {
							ok = false
							break
						}
					}
					if ok {
						var p_kotsu_len int = len(p_kotsu)
						var p_shuntsu_len int = len(p_shuntsu)
						//求值
						ret = p_kotsu_len + (p_shuntsu_len << 3) + (p_atama << 6)
						var lenght uint32 = 10
						for _, x := range p_kotsu {
							ret |= x << lenght
							lenght += 4
						}
						for _, x := range p_shuntsu {
							ret |= x << lenght
							lenght += 4
						}
						if a_size == 1 {
							//九莲宝灯
							var a_9 [][][]int = [][][]int{
								{{4, 1, 1, 1, 1, 1, 1, 1, 3}},
								{{3, 2, 1, 1, 1, 1, 1, 1, 3}},
								{{3, 1, 2, 1, 1, 1, 1, 1, 3}},
								{{3, 1, 1, 2, 1, 1, 1, 1, 3}},
								{{3, 1, 1, 1, 2, 1, 1, 1, 3}},
								{{3, 1, 1, 1, 1, 2, 1, 1, 3}},
								{{3, 1, 1, 1, 1, 1, 2, 1, 3}},
								{{3, 1, 1, 1, 1, 1, 1, 2, 3}},
								{{3, 1, 1, 1, 1, 1, 1, 1, 4}},
							}
							for _, a_9_v := range a_9 {
								if reflect.DeepEqual(a, a_9_v) {
									ret |= 1 << 27
									break
								}
							}
						}
						// [1,1,1,1,1,1,1,1,1]
						if a_size <= 3 && p_shuntsu_len >= 3 {
							var p_ikki int = 0
							for _, b := range a {
								if len(b) == 9 {
									var b_ikki1 bool = false
									var b_ikki2 bool = false
									var b_ikki3 bool = false
									for _, x_ikki := range p_shuntsu {
										if x_ikki == p_ikki {
											b_ikki1 = true
										}
										if x_ikki == (p_ikki + 3) {
											b_ikki2 = true
										}
										if x_ikki == (p_ikki + 6) {
											b_ikki3 = true
										}
									}
									if b_ikki1 && b_ikki2 && b_ikki3 {
										ret |= 1 << 28
									}
								}
								p_ikki += len(b)
							}
						}
						// [2,2,2, 2,2,2]
						if p_shuntsu_len == 4 &&
							p_shuntsu[0] == p_shuntsu[1] &&
							p_shuntsu[2] == p_shuntsu[3] {
							ret |= 1 << 29
						} else if p_shuntsu_len >= 2 &&
							(p_kotsu_len+p_shuntsu_len) == 4 {
							// [2,2,2]
							p_shuntsu_uniq := uniq_(p_shuntsu)
							if (p_shuntsu_len - len(p_shuntsu_uniq)) >= 1 {
								ret |= 1 << 30
							}
						}
						ret_array = append(ret_array, ret)
					}
				}
			}
			p_atama += 1
		}
	}
	//返回结果
	var ret_array_len int = len(ret_array)
	if ret_array_len > 0 {
		return ret_array
	}
	//七对子判定 [2,2,2,2,2,2,2]
	var a_f []int = flatten(a)
	isPair := func(val int) bool {
		if val == 2 {
			return true
		}
		return false
	}
	out, err := utils.All(isPair, a_f)
	if err != nil {
		panic(err)
	}
	if total(a_f) == 14 && out && err == nil {
		//return "0x" + (1 << 26)
		ret_pair := 1 << 26
		return []int{ret_pair}
	}
	return ret_array
}

func sevenPair(a [][]int) [][][]int {
	b := ptn(a)
	var c [][][]int = [][][]int{}
	isNotPair := func(val int) bool {
		if val != 2 {
			return true
		}
		return false
	}
	for _, v := range b {
		t := flatten(v)
		out, err := utils.Any(isNotPair, t)
		if err != nil {
			panic(err)
		}
		if out {
			continue
		}
		c = append(c, v)
	}
	return c
}

//打印输出各种和牌牌型组合
func chitoi() {
	/*
		//
		var chi01 [][]int = [][]int{{1,1,1},{1,1,1},{1,1,1},{1,1,1},{2}}
		var chi02 [][]int = [][]int{{1,1,1},{1,1,1},{1,1,1},{3},{2}}
		var chi03 [][]int = [][]int{{1,1,1},{1,1,1},{3},{3},{2}}
		var chi04 [][]int = [][]int{{1,1,1},{3},{3},{3},{2}}
		var chi05 [][]int = [][]int{{3},{3},{3},{3},{2}}
		//
		var chi06 [][]int = [][]int{{1,1,1},{1,1,1},{1,1,1},{2}}
		var chi07 [][]int = [][]int{{1,1,1},{1,1,1},{3},{2}}
		var chi08 [][]int = [][]int{{1,1,1},{3},{3},{2}}
		var chi09 [][]int = [][]int{{3},{3},{3},{2}}
		//
		var chi10 [][]int = [][]int{{1,1,1},{1,1,1},{2}}
		var chi11 [][]int = [][]int{{1,1,1},{3},{2}}
		var chi12 [][]int = [][]int{{3},{3},{2}}
		//
		var chi13 [][]int = [][]int{{1,1,1},{2}}
		var chi14 [][]int = [][]int{{3},{2}}
		//
		var chi15 [][]int = [][]int{{2}}
	*/

	//var chi16 [][]int = [][]int{{2}, {2}, {2}, {2}, {2}, {2}, {2}}
	//chi17 := sevenPair(chi16)
	//genTbl(chi17)

	var chi14 [][]int = [][]int{{3}, {2}}
	ptn_chi14 := ptn(chi14)
	genTbl(ptn_chi14)
}

func genTbl(a [][][]int) {
	b := uniq(a)
	for _, x := range b {
		fmt.Printf("0x%X: {", calc_key(x))
		valSlice := find_hai_pos(x)
		val := uniq_(valSlice)
		last := len(val) - 1
		for k, v := range val {
			if k == last {
				fmt.Printf("0x%X},\n", v)
			} else {
				fmt.Printf("0x%X, ", v)
			}
		}
	}
}
