package main

import (
	"fmt"
	"GoGame/base"
	"time"
	"GoGame/role"
)

var ch chan int
var equip [6]int;
var arr = [10]int{1,2,3,4,5}
var randInt int;

// 自动挂机 开关
var isAutoFight bool;
// 自动挂机 开关
var isAutoWalk bool;

//攻击速度
var attakeSpeed = time.Millisecond*800

//寻怪速度
var fightSpeed = time.Second*2

//闲逛速度
var walkSpeed = time.Second*2

//输入
var input string
func main() {
	fmt.Println("请输入姓名：")
	fmt.Scanf("%s",&input)
	role.NewHero(input)

	for true {

		fmt.Println("请选择行动，f为战斗，w为闲逛,s为暂停")
		fmt.Scanf("%s",&input)

		if GameOver() {
			fmt.Println("您的遗言是：",input)
			role.MyHero.Property()
			break
		}
		switch input {
			case "f":
				go AFK();
			case "w":
				go AWK();
			case "s":
				isAutoWalk = false;
				isAutoFight = false;
		}
	}

	fmt.Println("游戏结束")
}
//开始闲逛
func AWK(){
	//自动闲逛
	isAutoWalk = true;
	isAutoFight = false;

	walk()
}

//开始挂机
func AFK(){
	//自动挂机
	isAutoFight = true;
	isAutoWalk = false;

	fight()
}

//闲逛
func walk(){
	str := ""
	fmt.Println("奇遇中")
	time.Sleep(walkSpeed)
	randInt = Random(5)
	switch randInt{
		case WALK_TYPE_fight:
			str = "糟糕遭遇怪物了!!!"
			fmt.Println(str)
			fight()
		case WALK_TYPE_MEET:
			if walkSpeed > time.Second{
				walkSpeed -= time.Millisecond*100
				str = "偶遇仙人，闲逛速度提升。"
				fmt.Println(str)
			}else if(fightSpeed > time.Second){
				fightSpeed -= time.Millisecond*100
				str = "偶遇魔人，战斗速度提升。"
				fmt.Println(str)
			}else{
				str = "所有速度已到极致，无法提升。"
				fmt.Println(str)
			}
		case WALK_TYPE_BALE:
			str = "遭受不幸，丢失一件随机品质的装备。"
			fmt.Println(str);
			equip[3]++
		case WALK_TYPE_LUCK:
			str = "天降奇物，得到一件珍惜装备。"
			fmt.Println(str)
			equip[4]++
		case WALK_TYPE_NONE:
			str = "你闲逛了半天没有发现什么特殊的东西！"
			fmt.Println(str)
	}
	str = LinkStr(str,"\r\n");
	fmt.Println(str)
	base.CreatAndWrite("log.txt",str)
	if isAutoWalk{
		walk()
	}
}

//战斗
func fight(){
	fmt.Println("寻怪中.....")
	time.Sleep(fightSpeed)
	monster := findMonster()
	fmt.Println("发现怪物",monster.Name,"开始战斗")
	autoAttake(monster)

	if GameOver() {
		isAutoWalk = false
		isAutoFight = false
		return
	}
	randInt = Random(10001)
	equip[0]++
	if randInt == 10000{
		equip[5]++
	}else if randInt >= 9990{
		equip[4]++
	}else if randInt >= 9900{
		equip[3]++
	}else if randInt >= 9000{
		equip[2]++
	}else if randInt >= 5000{
		equip[1]++
	}

	str := "一共战斗：%d次，爆出白装%d件，爆出蓝装%d件，爆出紫装%d件，爆出橙装%d件，爆出金装%d件\r\n"
	str = fmt.Sprintf(str,equip[0],equip[1],equip[2],equip[3],equip[4],equip[5])
	fmt.Println(str)

	base.CreatAndWrite("log.txt",str)

	if isAutoFight{
		fight()
	}
}

func autoAttake(monster role.Monster) {
	for role.MyHero.IsAlive && monster.IsAlive{
		if role.MyHero.IsAlive && monster.IsAlive {
			time.Sleep(attakeSpeed)
			role.MyHero.Attak(&monster)
		}

		if role.MyHero.IsAlive && monster.IsAlive {
			time.Sleep(attakeSpeed)
			monster.Attak(&role.MyHero)
		}
	}
}
func findMonster() (monster role.Monster){
	randInt = Random(role.MyHero.Level*10)

	return role.Newmonster(randInt)
}

func GameOver() (b bool){
	return !role.MyHero.IsAlive
}