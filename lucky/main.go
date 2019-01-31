package main

import (
	"fmt"
	"math/rand"
	"time"
)

var userPool map[string]int
var luckySeed map[int]string
var totalSeed int

func init() {
	userPool = make(map[string]int)
	userPool["mi4tin"] = 10
	userPool["siblesos"] = 10
	userPool["ioformat"] = 10
	userPool["920631770"] = 11
	userPool["shaocongcong"] = 10
	userPool["xiam123"] = 10
	userPool["Zerooo0"] = 11 //表示相比正常用户多10%概率
	userPool["q920447939"] = 10
	userPool["gaoy13800"] = 10
	userPool["xiaodi2015"] = 13 //表示相比正常用户多30%概率
	userPool["WeilaScofield"] = 10

	resolveLuckySeed()
}

func main() {
	for i := 0; i < 3; i++ {
		rand.Seed(time.Now().UnixNano())
		seed := rand.Intn(totalSeed)
		luckyUser := luckySeed[seed]
		fmt.Println("抽奖池:", totalSeed, "随机数:", seed, "获奖用户:", luckyUser)
		removeUser(luckyUser)
		fmt.Println("移除用户:", luckyUser)
		time.Sleep(time.Second)
	}
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

func removeUser(user string) {
	delete(userPool, user)
	resolveLuckySeed()
}
