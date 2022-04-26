package storage

import (
	"game/interfaces"
	"game/tools"
	"strconv"
)

type Bag struct {
	BagLayout [5][10]string //4*10 背包 + 1*10 装备栏
	UI        interfaces.UIInterface
}

func New() *Bag {
	n := &Bag{
		BagLayout: [5][10]string{
			//包裹区
			{"dun-2_0,0", "dun-2_0,0", "book_0,2", "MP0", "MP0", "", "", "", "", ""},
			{"dun-2_0,0", "dun-2_0,0", "book_0,2", "box_1,3", "box_1,3", "", "dun-4_1,6", "dun-4_1,6", "", ""},
			{"dun-2_0,0", "dun-2_0,0", "box_1,3", "box_1,3", "", "", "dun-4_1,6", "dun-4_1,6", "", ""},
			{"dun-2_0,0", "dun-2_0,0", "HP0", "HP0", "HP0", "HP0", "dun-4_1,6", "dun-4_1,6", "", ""},
			//装备区域
			{"head-5", "futou-2", "dun-6", "neck", "body-4", "hand", "ring", "blet", "ring", "shose"},
			//头盔526,8  左手武器412,54 右手武器644,54 项链599,36 铠甲526,80 手套413,182 左戒指485,181 腰带527,181 右戒指599,183 靴子644,183
		}}
	return n
}

//插入DB
func (b *Bag) InsertBag(name string) bool {
	w, h := tools.GetItemsCellSize(name)
	//是否1*1 size的时候
	if w == 1 && h == 1 {
		for k, v := range b.BagLayout {
			for kk, vv := range v {
				if vv == "" {
					b.BagLayout[k][kk] = name
					b.UI.AddItemToBagByHand(k, kk, name)
					return true
				}
			}
		}
		return false
	} else {
		//
		sum_err := 0
		for i := 0; i < 4; i++ {
			for j := 0; j < 10; j++ {
				if b.BagLayout[i][j] == "" {
					for a := 0; a <= w-1; a++ {
						for c := 0; c <= h-1; c++ {
							if i+c > 3 || j+a > 9 {
								sum_err++
							} else if b.BagLayout[i+c][j+a] != "" {
								sum_err++
							}
						}
					}
					if sum_err == 0 {
						subName := strconv.Itoa(i) + "," + strconv.Itoa(j)
						b.UI.AddItemToBagByHand(i, j, name)
						for a := 0; a <= w-1; a++ {
							for c := 0; c <= h-1; c++ {
								b.BagLayout[i+c][j+a] = name + "_" + subName
							}
						}
						return true
					}
					sum_err = 0
				}
			}
		}
		return false
	}
}
