// source -> http://hp.vector.co.jp/authors/VA046927/mjscore/mjalgorism.html
package algo

import "reflect"

//初始化一个长度为34,值为0的列表
//var n_zero []int = []int{33: 0}
var n_zero []int = []int{LEN - 1: 0}

/* 先给手牌(hai)排序,计算连续牌的张数,得到一个特制"牌型"
 * 例如:
 * 四张相同牌=4
 * 三张相同牌=3
 * 两张相同牌=2
 * 三张连续牌=1,1,1
 * 中间不连续0隔开
 * 如：22234577 编码后 311102
 * 牌的个数不同种类的要求,例如:
 * 「１２３」→「１１１」
 * 「５６７」→「１１１」
 * 「１１１」→「３」
 * 「３３３」→「３」
 * 「３３３３」→「４」
 * 「２３４４５６」→「１１２１１」
 */
func analyse(hai []int) []int {
	var n []int = make([]int, LEN, LEN)
	copy(n, n_zero)
	for _, v := range hai {
		n[v]++
	}
	return n
}

//暴力拆解法(比较吃计算)
func agari(n []int) [][]int {
	var ret [][]int = make([][]int, 3, 3)
	var len int = len(n)
	for i := 0; i < LEN; i++ {
		for kotsu_first := 0; kotsu_first < 2; kotsu_first++ {
			var janto []int = make([]int, 0)
			var kotsu []int = make([]int, 0)
			var shuntsu []int = make([]int, 0)
			var t []int = make([]int, len, len)

			copy(t, n)
			if t[i] >= 2 {
				//取出雀頭(两个相同牌)
				t[i] -= 2
				janto = append(janto, i)

				if kotsu_first == 0 {
					//取出刻子(三个相同牌)
					for j := 0; j < LEN; j++ {
						if t[j] >= 3 {
							t[j] -= 3
							kotsu = append(kotsu, j)
						}
					}
					//取出顺子(三个连续的牌)
					for a := 0; a < 3; a++ {
						for b := 0; b < 7; {
							if t[9*a+b] >= 1 &&
								t[9*a+b+1] >= 1 &&
								t[9*a+b+2] >= 1 {
								t[9*a+b]--
								t[9*a+b+1]--
								t[9*a+b+2]--
								shuntsu = append(shuntsu, 9*a+b)
							} else {
								b++
							}
						}
					}
				} else {
					//取出顺子(三个连续的牌)
					for a := 0; a < 3; a++ {
						for b := 0; b < 7; {
							if t[9*a+b] >= 1 &&
								t[9*a+b+1] >= 1 &&
								t[9*a+b+2] >= 1 {
								t[9*a+b]--
								t[9*a+b+1]--
								t[9*a+b+2]--
								shuntsu = append(shuntsu, 9*a+b)
							} else {
								b++
							}
						}
					}
					//取出刻子(三个相同牌)
					for j := 0; j < LEN; j++ {
						if t[j] >= 3 {
							t[j] -= 3
							kotsu = append(kotsu, j)
						}
					}
				}
				//和了
				if reflect.DeepEqual(t, n_zero) {
					ret[0] = janto
					ret[1] = kotsu
					ret[2] = shuntsu
				}
			}
		}
	}
	return ret
}

//查表法(比较吃内存)
func agari_tbl(key uint32) []uint32 {
	return tbl[key]
}

/* 键值计算
得到特制"牌型",再将其二进制化,采用特制规则编码:
1  -> 0
2  -> 110
3  -> 11110
4  -> 1111110
10 -> 10
20 -> 1110
30 -> 111110
40 -> 11111110
如麻将中34种牌(含字牌)至少需要6位表示,手牌14张就要84位
特制编码后每张牌只占用1~2位,最大(14张不连单牌)仅占27位
和牌的排列组合表在特制编码后也就少
var n []int = []int{3, 1, 1, 1, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 2, 0, 0, 0, 0}
var pos []int = []int{13: 0} //14个0
*/
func calc_key_tbl(n, pos []int) uint32 {
	var p int = -1
	var x uint32 = 0
	var pos_p int = 0  //pos的排列索引
	var b bool = false //前一个不是0
	//数牌
	for i := 0; i < 3; i++ {
		for j := 0; j < 9; j++ {
			if n[i*9+j] == 0 {
				if b {
					b = false
					x |= 0x1 << uint32(p)
					p++
				}
			} else {
				p++
				b = true
				pos[pos_p] = i*9 + j
				pos_p++
				switch n[i*9+j] {
				case 2:
					x |= 0x3 << uint32(p)
					p += 2
				case 3:
					x |= 0xF << uint32(p)
					p += 4
				case 4:
					x |= 0x3F << uint32(p)
					p += 6
				}
			}
		}
		if b {
			b = false
			x |= 0x1 << uint32(p)
			p++
		}
	}
	/*//字牌
	for i := TON; i <= CHU; i++ {
		if n[i] > 0 {
			p++
			pos[pos_p] = i
			pos_p++
			switch n[i] {
			case 2:
				x |= 0x3 << uint32(p)
				p += 2
			case 3:
				x |= 0xF << uint32(p)
				p += 4
			case 4:
				x |= 0x3F << uint32(p)
				p += 6
			}
			x |= 0x1 << uint32(p)
			p++
		}
	}
	*/
	return x
}

//查表法
func agari_index(n []int) bool {
	var a []int = analyse(n)
	var pos []int = []int{HAND: 0} //14个0
	var key uint32 = calc_key_tbl(a, pos)
	var ret []uint32 = agari_tbl(key)
	if len(ret) == 0 {
		return false
	}
	return true
	//for _, r := range ret {
	//	log.Debug("雀頭=%v", pos[(r>>6)&0xF])
	//	var num_kotsu uint32 = r & 0x7
	//	var num_shuntsu uint32 = (r >> 3) & 0x7
	//	var i uint32
	//	for i = 0; i < num_kotsu; i++ {
	//		log.Debug("刻子=%v", pos[(r>>(10+i*4))&0xF])
	//	}
	//	for i = 0; i < num_shuntsu; i++ {
	//		log.Debug("順子=%v", pos[(r>>(10+num_kotsu*4+i*4))&0xF])
	//	}
	//}
}
