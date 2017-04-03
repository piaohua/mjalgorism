package algo

import (
	"basic/utils"
	"sort"
)

type IntSlice []int

//排序接口实现
func (b IntSlice) Len() int {
	return len(b)
}

func (b IntSlice) Less(i, j int) bool {
	return b[i] < b[j]
}

func (b IntSlice) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

//洗牌
func Shuffle() []int {
	utils.Seed()
	d := make([]int, TOTAL, TOTAL)
	copy(d, CARDS)
	// 测试暂时去掉洗牌
	for i := range d {
		j := utils.RandIntN(i + 1)
		d[i], d[j] = d[j], d[i]
	}
	return d
}

//排序
func Sort(cards []int) {
	if !sort.IsSorted(IntSlice(cards)) {
		sort.Sort(IntSlice(cards))
	}
}

//胡牌检测
func HasHu(cards []int) bool {
	Sort(cards)
	return agari_index(cards)
}

//杠牌检测
func HasKong(cards []int) bool {
	Sort(cards)
	var n []int = analyse(cards)
	isKong := func(val int) bool {
		if val == 4 {
			return true
		}
		return false
	}
	out, err := utils.Any(isKong, n)
	if err != nil {
		panic(err)
	}
	return out
}

//碰牌检测
func HasPong(cards []int) bool {
	Sort(cards)
	var n []int = analyse(cards)
	isPong := func(val int) bool {
		if val == 3 {
			return true
		}
		return false
	}
	out, err := utils.Any(isPong, n)
	if err != nil {
		panic(err)
	}
	return out
}

//牌检测
func HaveCard(card int, cards []int) bool {
	for _, c := range cards {
		if c == card {
			return true
		}
	}
	return false
}

//牌检测,返回下标,没找到返回长度
func IndexCard(card int, cards []int) int {
	for i, c := range cards {
		if c == card {
			return i
		}
	}
	return len(cards)
}

//移除n个指定牌
func RemoveCard(n, card int, cards []int) []int {
	for i := 0; i < n; i++ {
		for j, c := range cards {
			if c == card {
				cards = append(cards[:j], cards[j+1:]...)
				break
			}
		}
	}
	return cards
}

//添加牌
func AddCards(cards []int, ends []int) []int {
	return append(cards, ends...)
}

//添加牌
func AddCard(cards []int, card int) []int {
	return append(cards, card)
}

//杠牌检测
func HaveKong(card int, cards []int) bool {
	var num int = 0
	for _, c := range cards {
		if c == card {
			num++
			if num == 3 {
				return true
			}
		}
	}
	return false
}

//碰牌检测
func HavePong(card int, cards []int) bool {
	var num int = 0
	for _, c := range cards {
		if c == card {
			num++
			if num == 2 {
				return true
			}
		}
	}
	return false
}

//计算下一家位置
func TurnNext(seat int) int {
	if seat == 4 {
		return 1
	}
	return seat + 1
}

//摸牌时检测(扛胡)
func DrawKong(card int, pongs, hands []int) uint32 {
	var length int = len(hands)
	var cards []int = make([]int, length, length)
	copy(cards, hands)
	var ret uint32 = 0
	if HasKong(cards) {
		ret |= AN_KONG
	}
	if HaveKong(card, cards) {
		ret |= AN_KONG
	}
	if HaveKong(card, pongs) {
		ret |= BU_KONG
	}
	cards = append(cards, card)
	if HasHu(cards) {
		ret |= HU
	}
	return ret
}

//打牌时检测(碰扛胡),state=true表示可以胡
func TurnPong(state bool, card int, hands []int) uint32 {
	var length int = len(hands)
	var cards []int = make([]int, length, length)
	copy(cards, hands)
	var ret uint32 = 0
	if HasKong(cards) {
		ret |= AN_KONG
	}
	if HaveKong(card, cards) {
		ret |= MI_KONG
	}
	if HavePong(card, cards) { //优化,碰杠可一起检测
		ret |= PONG
	}
	if state {
		cards = append(cards, card)
		if HasHu(cards) {
			ret |= HU
		}
	}
	return ret
}

//打牌(选一张最优出牌)
func Discard(hands []int) int {
	var length int = len(hands)
	var cards []int = make([]int, length, length)
	copy(cards, hands)
	Sort(cards)
	for i, l := 0, len(cards); i < l-3; {
		if cards[i] == cards[i+1] &&
			cards[i] == cards[i+2] {
			i += 3
		} else if cards[i] == cards[i+1] {
			i += 2
		} else if cards[i] == cards[i+1]-1 &&
			cards[i] == cards[i+2]-2 {
			i += 3
		} else {
			return cards[i]
		}
	}
	return 0
}
