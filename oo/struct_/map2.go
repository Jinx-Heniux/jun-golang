package struct_

import (
	"fmt"
	"sort"
)

func MapExample2() {
	map1 := make(map[int]string, 5)
	map1[1] = "www.topgoer.com"
	map1[2] = "rpc.topgoer.com"
	map1[5] = "ceshi"
	map1[3] = "xiaohong"
	map1[4] = "xiaohuang"
	fmt.Printf("%v,%p\n", map1, &map1)
	fmt.Println()

	for k, v := range map1 {
		fmt.Printf("%d -> %s, %p\n", k, v, &v)
	}
	fmt.Println()

	sli := []int{}
	for k, _ := range map1 {
		sli = append(sli, k)
	}
	sort.Ints(sli)
	fmt.Println(sli)
	fmt.Println()

	for i := 0; i < len(map1); i++ {
		fmt.Printf("%s,%p\n", map1[sli[i]], &map1)
	}
}

/*
map[1:www.topgoer.com 2:rpc.topgoer.com 3:xiaohong 4:xiaohuang 5:ceshi],0xc00000e028

1 -> www.topgoer.com, 0xc0000102a0
2 -> rpc.topgoer.com, 0xc0000102a0
5 -> ceshi, 0xc0000102a0
3 -> xiaohong, 0xc0000102a0
4 -> xiaohuang, 0xc0000102a0

[1 2 3 4 5]

www.topgoer.com,0xc00000e028
rpc.topgoer.com,0xc00000e028
xiaohong,0xc00000e028
xiaohuang,0xc00000e028
ceshi,0xc00000e028
*/
