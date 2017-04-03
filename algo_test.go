package algo

import (
	"sort"
	"testing"
	"time"
)

// go test -v -run='Test_Slice' //执行指定
func Test_Slice(t *testing.T) {
	cards := []int{19, 7, 8, 6, 1, 22, 2, 12, 6, 3, 24, 2, 3}
	hands := make(map[int][]int)
	var s []int = make([]int, 3)
	//hands[1] = cards[2:5]
	copy(s, cards[2:5])
	hands[1] = s
	h := hands[1]
	t.Log("h -> %v", h[0])
	t.Log("hands -> %v", hands[1])
	t.Log("cards -> %v", cards)
	h = append(h, 28)
	hands[1] = h
	cards[3] = 100
	t.Log("hands -> %v", hands[1])
	t.Log("cards -> %v", cards)
	cards = cards[3:]
	t.Log("cards -> %v", cards)

	b := []int{19, 7, 3, 24, 2, 3}
	t.Log("b -> ", b[len(b):], b[len(b)-1], b)
	t.Log("b -> ", b[:5], b[6:])
	t.Log("hands -> ", hands[1])
	p := hands[1]
	p = append(p, 10)
	p = append(p, 11, 12)
	t.Log("p -> ", p)
	t.Log("hands -> ", hands[1])
	p = RemoveCard(1, 12, p)
	t.Log("p -> ", p)
}

// go test -v -run='Test_Perms' //执行指定
func Test_Perms(t *testing.T) {
	var a [][]int = [][]int{{1, 1, 1}, {2}, {3}}
	t.Log(a)
	rs := perms(a)
	t.Log(rs)
}

// go test -v -run='Test_DrawKong' //执行指定
func Test_DrawKong(t *testing.T) {
	var hands []int = []int{1, 1, 1, 1, 5, 6, 7, 8, 9, 22, 23, 21, 24}
	var pongs []int = []int{}
	var value uint32 = DrawKong(5, pongs, hands)
	t.Log(value)
}

// go test -v -run='Test_Discard' //执行指定
func Test_Discard(t *testing.T) {
	var hands []int = []int{1, 1, 1, 1, 5, 6, 7, 8, 9, 22, 23, 21, 24}
	var card int = Discard(hands)
	t.Log(card)
}

// go test -v -run='Test_TurnPong' //执行指定
func Test_TurnPong(t *testing.T) {
	var hands []int = []int{1, 1, 1, 1, 5, 6, 7, 8, 9, 22, 23, 21, 24}
	var value uint32 = TurnPong(false, 5, hands)
	t.Log(value)
}

// go test -v -run='Test_chitoi' >> log //执行指定
func Test_chitoi(t *testing.T) {
	//valSlice := []int{0x21104420, 0x211000E0, 0x211000E0, 0x20CC01A0, 0x20CC01A0}
	//t.Log("valSlice: ", valSlice)
	//val := uniq_(valSlice)
	//t.Log("valSlice: ", valSlice)
	//t.Log("valSlice: ", val)
	chitoi()
	t.Log("chitoi")
}

// go test -v -run='Test_ptn' //执行指定
func Test_ptn(t *testing.T) {
	var a [][]int = [][]int{{1, 1, 1}, {2}}
	t.Log(a)
	rs := ptn(a)
	t.Log(rs)
	//s := 1 << 26
	//in := reflect.ValueOf(s).Type()
	//t.Log("in:", in)
	//t.Log("in:", s)
}

// go test -v -run='Test_find_hai_pos' //执行指定
func Test_find_hai_pos(t *testing.T) {
	//var a1 []int = []int{3, 2}
	//var k1 []int = make([]int, len(a1), len(a1))
	//t.Log("a1 -> ", a1)
	//err := utils.Clone(&k1, &a1)
	//t.Log("k1 -> ", k1)
	//t.Log("err -> ", err)

	var a [][]int = [][]int{{3}, {2}}
	t.Log(a)

	//var k [][]int = make([][]int, len(a), len(a))
	//copy(k, a)
	//t.Log(k)
	//for k1, v1 := range k {
	//	t.Log(k1, "-", v1)
	//	v1[0] += 1
	//}
	//t.Log("k -> ", k)
	//t.Log("a -> ", a)

	rs := ptn(a)
	t.Log(rs)
	for _, v := range rs {
		//t.Log(" val -> ", v)
		//t.Log("key -> ", calc_key(v))
		//t.Log("val -> ", find_hai_pos(v))
		//t.Log(" val -> ", v)
		t.Logf("key -> 0x%X, val -> %X", calc_key(v), find_hai_pos(v))
		//fmt.Printf("key -> 0x%X, val -> %X", calc_key(v), find_hai_pos(v))
	}
}

//
func Test_algo(t *testing.T) {
	cards := []byte{0x01, 0x01, 0x08, 0x08, 0x14, 0x14, 0x17, 0x17, 0x22, 0x22, 0x22, 0x29, 0x29, 0x22}
	t.Log(cards)
}

func Test_card(t *testing.T) {

	t.Log("--> %x,%x,%x,%x\n", 0x21, 0x22, 0x23, 0x24)
	t.Log("--> %d,%d,%d,%d\n", 0x21, 0x22, 0x23, 0x24)

	t.Log("--> %x,%x,%x,%x\n", CARDS[72], CARDS[73], CARDS[74], CARDS[75])
	t.Log("--> %d,%d,%d,%d\n", CARDS[72], CARDS[73], CARDS[74], CARDS[75])
}

type ByteSlice []byte

func (b ByteSlice) Len() int {
	return len(b)
}

func (b ByteSlice) Less(i, j int) bool {
	return b[i] < b[j]
}

func (b ByteSlice) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func Test_sort(t *testing.T) {
	a := ByteSlice{0x01, 0x01, 0x08, 0x08, 0x14, 0x14, 0x17, 0x17, 0x22, 0x22, 0x22, 0x29, 0x29, 0x22}
	if !sort.IsSorted(a) {
		sort.Sort(a)
	}
	b := []float64{0.9, 1.49, 11.2, 0.56, 8.11, 1.1}
	if !sort.Float64sAreSorted(b) {
		sort.Float64s(b)
	}
	c := []int{10, 24, 9, 4, 11, 38, 1, 45}
	if !sort.IntsAreSorted(c) {
		sort.Ints(c)
	}
	//当slice a 升序排序时func使用>=
	//当slice a 降序排序时func使用<=
	//返回找到的第一个元素下标，没找到返回长度
	s := sort.Search(len(a), func(i int) bool { return a[i] >= 0x22 })
	t.Log("a -> ", a)
	t.Log("b -> ", b)
	t.Log("c -> ", c)
	t.Log("s -> ", s)
}

// func GuessingGame() {
// 	var s string
// 	fmt.Printf("Pick an integer from 0 to 100.\n")
// 	answer := sort.Search(100, func(i int) bool {
// 		fmt.Printf("Is your number <= %d?", i)
// 		fmt.Scanf("%s", &s)
// 		return s != "" && s[0] == 'y'
// 	})
// 	fmt.Printf("Your number is %d.\n", answer)
// }

// var hai []int = []int{
// 	MAN1, MAN1, MAN1,
// 	MAN2, MAN3, MAN4,
// 	MAN6, MAN7, MAN8,
// 	TON, TON, TON,
// 	SHA, SHA,
// }

//var hai []int = []int{
//	MAN1, MAN1, MAN1,
//	MAN2, MAN3, MAN4,
//	MAN6, MAN7, MAN8,
//	SOU1, SOU1, SOU1,
//	PIN3, PIN3,
//}

var hai []int = []int{
	MAN1, MAN1, MAN1, MAN1,
	MAN2, MAN3, MAN4, MAN4,
	MAN6, MAN7, MAN8, MAN8,
	MAN8, MAN8,
}

// go test -v -run='*.go' //执行所有
// go test -v -run='Test_agari' //执行指定
func Test_agari(t *testing.T) {
	start := time.Now()
	n := analyse(hai)
	//n := []int{3, 1, 1, 1, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 2, 0, 0, 0, 0}
	ret := agari(n)
	end := time.Now()
	t.Log("time consume -> ", end.Sub(start))
	t.Log("雀頭=", ret[0])
	for _, kotsu := range ret[1] {
		t.Log("刻子=", kotsu)
	}
	for _, shuntsu := range ret[2] {
		t.Log("順子=", shuntsu)
	}
}

// go test -run=文件名字 -bench='bench名字' -cpuprofile=生成的cpuprofile文件名字 文件夹
// go test -v bench='.*' //执行所有
// go test -test.bench='.*' //执行所有
// go test -test.bench='BenchmarkAgari' //执行指定
// go test -test.bench='BenchmarkAgari' -count=100000 //执行指定,指定执行次数
func BenchmarkAgari(b *testing.B) {
	n := analyse(hai)
	ret := agari(n)
	b.Logf("雀頭=%v", ret[0])
	for _, kotsu := range ret[1] {
		b.Logf("刻子=%v", kotsu)
	}
	for _, shuntsu := range ret[2] {
		b.Logf("順子=%v", shuntsu)
	}
}

// go test -v -run='Test_calc_key' //执行指定
func Test_calc_key(t *testing.T) {
	//var a [][]int = [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {2}}
	//var a [][]int = [][]int{{3}, {3}, {3}, {2}}
	var a [][]int = [][]int{{1, 1, 1}, {3}, {3}, {2}}
	key := calc_key(a)
	t.Log("key -> ", key)
}

// go test -v -run='Test_Agari_Tbl' //执行指定
func Test_Agari_Tbl(t *testing.T) {
	var n []int
	var pos []int = []int{13: 0} //14个0
	var ret []uint32

	n = analyse(hai)
	//n := []int{3, 1, 1, 1, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 2, 0, 0, 0, 0}
	t.Log("n -> ", n)
	key := calc_key_tbl(n, pos)
	t.Log("key -> ", key)
	ret = agari_tbl(key)
	t.Log("ret -> ", ret)

	for _, r := range ret {
		t.Log("雀頭=", pos[(r>>6)&0xF])
		var num_kotsu uint32 = r & 0x7
		var num_shuntsu uint32 = (r >> 3) & 0x7
		var i uint32
		for i = 0; i < num_kotsu; i++ {
			t.Log("刻子=", pos[(r>>(10+i*4))&0xF])
		}
		for i = 0; i < num_shuntsu; i++ {
			t.Log("順子=", pos[(r>>(10+num_kotsu*4+i*4))&0xF])
		}
	}
}

// go test -test.bench='BenchmarkTblAgari' //执行指定
func BenchmarkTblAgari(b *testing.B) {
	var n []int
	var pos []int = []int{13: 0} //14个0
	var ret []uint32

	n = analyse(hai)
	//n := []int{3, 1, 1, 1, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 2, 0, 0, 0, 0}
	b.Logf("n -> %v", n)
	key := calc_key_tbl(n, pos)
	b.Logf("key -> %v", key)
	ret = agari_tbl(key)
	b.Logf("ret -> %v", ret)

	for _, r := range ret {
		b.Logf("雀頭=%v", pos[(r>>6)&0xF])
		var num_kotsu uint32 = r & 0x7
		var num_shuntsu uint32 = (r >> 3) & 0x7
		var i uint32
		for i = 0; i < num_kotsu; i++ {
			b.Logf("刻子=%v", pos[(r>>(10+i*4))&0xF])
		}
		for i = 0; i < num_shuntsu; i++ {
			b.Logf("順子=%v", pos[(r>>(10+num_kotsu*4+i*4))&0xF])
		}
	}
}

// go test -v -run='Test_Calc_Key_Tbl' //执行指定
func Test_Calc_Key_Tbl(t *testing.T) {
	var n []int = []int{3, 1, 1, 1, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 2, 0, 0, 0, 0}
	t.Log("n -> ", n)
	var pos []int = []int{13: 0} //14个0
	key := calc_key_tbl(n, pos)
	t.Log("key -> ", key)
}

// go test -v -run='TestCalcKeyTbl' //执行指定
func TestCalcKeyTbl(t *testing.T) {
	a := [][]int{{3}, {1, 1, 1}, {1, 1, 1}, {3}, {2}}
	key1 := calc_key(a)
	t.Log("key1 -> ", key1)
	n := analyse(hai)
	t.Log("n -> ", n)
	var pos []int = []int{13: 0} //14个0
	key2 := calc_key_tbl(n, pos)
	t.Log("key2 -> ", key2)
}

// go test -v -run='Test_he' //执行指定
// 测试胡牌
func Test_he(t *testing.T) {
	timeStart := time.Now()
	//cards := []byte{0x01,0x01}
	//cards := []byte{0x01,0x01,0x01,0x02,0x03}
	//cards := []byte{0x03,0x03,0x03,0x04,0x04,0x04,0x09,0x09}
	//cards := []byte{0x01,0x01,0x01,0x01,0x02,0x03,0x04,0x05,0x06,0x07,0x07}
	//
	//cards := []byte{0x01,0x01,0x01,0x01,0x02,0x03,0x04,0x04,0x06,0x07,0x08,0x08,0x08,0x08}
	//cards := []byte{0x01,0x01,0x02,0x02,0x03,0x03,0x11,0x11,0x17,0x17,0x21,0x22,0x08,0x08}
	//ok := false
	//for i:=0; i<2000000;i++{
	//	ok = he(cards)
	//	//ok = he2(cards)
	//}
	//
	cards := []byte{0x01, 0x01, 0x01, 0x01, 0x02, 0x03, 0x04, 0x04, 0x06, 0x07, 0x08, 0x08, 0x08, 0x08}
	ok := he(cards)
	//
	//cards := []byte{0x01,0x01,0x01,0x01,0x02,0x03,0x04,0x04,0x06,0x07,0x08,0x08,0x08,0x08}
	//ok := he2(cards)
	timeEnd := time.Now()
	t.Log("time consume:", timeEnd.Sub(timeStart))
	t.Log(cards, " -> ", ok)
}

func BenchmarkHe(b *testing.B) {
	cards := []byte{0x01, 0x01, 0x01, 0x01, 0x02, 0x03, 0x04, 0x04, 0x06, 0x07, 0x08, 0x08, 0x08, 0x08}
	he(cards)
}

func he(cs []byte) bool {
	ok := false
	le := len(cs)
	if (le-2)%3 != 0 {
		return ok
	}
	//Sort(cs, 0, le-1)
	sort.Sort(ByteSlice(cs))
	for n := 0; n < le-1; n++ {
		if cs[n] == cs[n+1] { //取2
			ok = true
			list := make([]byte, le)
			copy(list, cs)
			list[n] = 0
			list[n+1] = 0
			for i := 0; i < le-2; i++ {
				if list[i] > 0 {
					for j := i + 1; j < le-1; j++ {
						if list[j] > 0 && list[i] > 0 {
							for k := j + 1; k < le; k++ {
								if list[k] > 0 && list[i] > 0 && list[j] > 0 {
									if list[i]+1 == list[j] && list[j]+1 == list[k] {
										//fmt.Println("i, j, k -> ", i, j, k)
										//fmt.Printf("list i, j, k -> %x, %x, %x\n", list[i], list[j], list[k])
										list[i] = 0
										list[j] = 0
										list[k] = 0
										break
									} else if list[i] == list[j] && list[j] == list[k] {
										//fmt.Println("i, j, k -> ", i, j, k)
										//fmt.Printf("list i, j, k -> %x, %x, %x\n", list[i], list[j], list[k])
										list[i] = 0
										list[j] = 0
										list[k] = 0
										break
									}
								}
							}
						}
					}
				}
			}
			//fmt.Println("list -> ", list)
			for i := 0; i < le; i++ {
				if list[i] > 0 {
					ok = false
					break
				}
			}
			if ok {
				break
			}
		}
	}
	return ok
}

func he2(cs []byte) bool {
	le := len(cs)
	if (le-2)%3 != 0 {
		return false
	}
	sort.Sort(ByteSlice(cs))
	//Sort(cs, 0, le-1)
	ok := false
	for n := 0; n < le-1; n++ {
		if cs[n] != cs[n+1] { //取2
			continue
		}
		ok = true
		list := make([]byte, le)
		copy(list, cs)
		list[n] = 0
		list[n+1] = 0
		for i := 0; i < le-2; i++ {
			//if list[i] == 0 {
			if !(list[i] > 0) {
				continue
			}
			for j := i + 1; j < le-1; j++ {
				//if list[j] == 0 || list[i] == 0 {
				if !(list[j] > 0 && list[i] > 0) {
					continue
				}
				for k := j + 1; k < le; k++ {
					//if list[k] == 0 || list[i] == 0 || list[j] == 0 {
					if !(list[k] > 0 && list[i] > 0 && list[j] > 0) {
						continue
					}
					if list[i]+1 == list[j] && list[j]+1 == list[k] {
						//fmt.Println("i, j, k -> ", i, j, k)
						//fmt.Printf("list i, j, k -> %x, %x, %x\n", list[i], list[j], list[k])
						list[i] = 0
						list[j] = 0
						list[k] = 0
						break
					} else if list[i] == list[j] && list[j] == list[k] {
						//fmt.Println("i, j, k -> ", i, j, k)
						//fmt.Printf("list i, j, k -> %x, %x, %x\n", list[i], list[j], list[k])
						list[i] = 0
						list[j] = 0
						list[k] = 0
						break
					}
				}
			}
		}
		//fmt.Println("list -> ", list)
		for i := 0; i < le; i++ {
			if list[i] > 0 {
				ok = false
				break
			}
		}
		if ok {
			break
		}
	}
	return ok
}
