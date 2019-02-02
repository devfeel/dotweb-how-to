package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var userPool map[string]int
var luckySeed map[int]string
var totalSeed int
var userNames string = "bibibi-bug,dotqi,WithLin,youzizzz,zcmiao190,wbly7758,misasky,qizhenshuai,motanzg,Wooshawn,itstudying,vhuyuy,jokerddj,tianguozhifeng,SHIJIEZ,gumochengxing,dfy167,LyricTian,hades2013,vinnking,wweir,violetgo,Imhven,xskingdeekis,wangyun111,memoryxyd,chplain,willxuecn,yibalaodao,xiang2lin,lyw1995,hacfox,xingliuhua,yrx1354,mfcab,hunterfall,solarhell,thesilent,lhkzx007,171869092,haphupan,921034403,yuzhoutong,baidw,Sakurasan,NightFarmer,imxyb,xiam123,qbmiller,Uncho,LinlinDu,shaocongcong,sayhei,lvchengchang,le0l1,1509477745,dcwang1989,swing4,archer-bao,sszgr,83103005,qq976739120,aredcup,sowhaton,angusguo,royromny,wx7217242,thhy,rongxr,MonkeyInWind"

func init() {
	userPool = make(map[string]int)
	users := strings.Split(userNames, ",")
	for _, v := range users {
		userPool[v] = 1
	}

	fmt.Println(`    ____           __                     __`)
	fmt.Println(`   / __ \  ____   / /_ _      __  ___    / /_`)
	fmt.Println(`  / / / / / __ \ / __/| | /| / / / _ \  / __ \`)
	fmt.Println(` / /_/ / / /_/ // /_  | |/ |/ / /  __/ / /_/ /`)
	fmt.Println(`/_____/  \____/ \__/  |__/|__/  \___/ /_.___/`)
	fmt.Println(``)
	fmt.Println(`              → 2019新春抽奖活动 ←              `)
	fmt.Println(`https://github.com/devfeel/dotweb/issues/181`)
	fmt.Println(``)

	fmt.Println("初始化名单，共计", len(userPool), "人")
	fmt.Println("名单公示：", userPool)
	fmt.Println(``)
	resolveLuckySeed()
}

func resolveLuckySeed() {
	beginSeed := 0
	luckySeed = make(map[int]string)
	for k, v := range userPool {
		for i := 1; i <= v; i++ {
			beginSeed += 1
			luckySeed[beginSeed] = k
		}
	}
	totalSeed = beginSeed
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("开始抽奖倒计时！")
	fmt.Println("倒数1秒")
	time.Sleep(time.Second)
	fmt.Println("倒数2秒")
	time.Sleep(time.Second)
	fmt.Println("倒数3秒")
	time.Sleep(time.Second)
	fmt.Println(``)
	for i := 0; i < 3; i++ {
		seed := rand.Intn(totalSeed)
		luckyUser := luckySeed[seed]
		fmt.Println("用户池:", totalSeed, "随机数:", seed, "获奖用户:", luckyUser)
		removeUser(luckyUser)
		fmt.Println("移除用户:", luckyUser)
		time.Sleep(time.Second)
	}

	fmt.Println(``)
	fmt.Println("抽奖结束，恭喜以上用户！")
}

func removeUser(user string) {
	delete(userPool, user)
	resolveLuckySeed()
}
