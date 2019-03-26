package role

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func init(){
	initMonster()
}

func initMonster(){
	data,err := ioutil.ReadFile("hong.com/role/monster.json")
	if err != nil{
		fmt.Println("打开文件失败",err.Error())
	}
	err = json.Unmarshal(data,&monsters)

	if err != nil{
		fmt.Println("error:",err)
	}
}
var monsters []Monster
type Role struct {
	Name string
	Level int
	Chp,Hp,Attake,Defense float64
	IsAlive bool
}

var MyHero Hero
type living interface {
	Attak()
}

func (r Role) Attak(target *Role){
	if (r.Attake - target.Defense)>1 {
		target.Hp -= r.Attake - target.Defense
	}else {
		target.Hp -= 1
		fmt.Print(target.Name)
	}
}

func(r Monster) Property(){
	fmt.Printf("姓名：%s，等级：%d，当前生命:%g，最大生命:%g，攻击:%g，防御:%g,存活状况:%t\n",r.Name,r.Level,r.Chp,r.Hp,r.Attake,r.Defense,r.IsAlive)
}

//func (r Monster) Attak(target *Monster){
//	damage := 0.0
//	if (r.Attake - target.Defense)>1 {
//		damage = r.Attake - target.Defense
//	}else {
//		damage = 1
//	}
//	target.Chp -= damage
//	fmt.Printf("%s攻击了%s造成了%g点伤害，%s剩余%g血\n",r.Name,target.Name,damage,target.Name,target.Hp)
//}

func (r Monster) Attak(target *Hero) (finish bool){
	damage := 0.0
	if (r.Attake - target.Defense)>1 {
		damage = r.Attake - target.Defense
	}else {
		damage = 1
	}
	target.Chp -= damage

	if target.Chp <= 0{
		target.IsAlive = false
	}

	fmt.Printf("%s攻击了%s造成了%g点伤害，%s剩余%g血\n",r.Name,target.Name,damage,target.Name,target.Chp)
	if target.IsAlive == false {
		fmt.Println("英雄死亡")
		return true
	}
	return false
}

type Monster struct {
	Role
	Id int
}

type Hero struct{
	Role
	exp int
	money int
	killNumber int
}

func (r Hero) Attak(target *Monster) (finish bool){
	damage := 0.0
	if (r.Attake - target.Defense)>1 {
		damage = r.Attake - target.Defense
	}else {
		damage = 1
	}
	target.Chp -= damage

	if target.Chp <= 0{
		target.IsAlive = false
	}

	fmt.Printf("%s攻击了%s造成了%g点伤害，%s剩余%g血\n",r.Name,target.Name,damage,target.Name,target.Chp)
	if !target.IsAlive{
		fmt.Println("战斗结束，恭喜你获得了胜利")
		fmt.Printf("得到了%d经验，%d金币\n",target.Level,target.Level)
		MyHero.exp++;
		MyHero.money++;
		MyHero.killNumber++;
		MyHero.upLevel();
		return true
	}

	return false
}
func NewHero(name string){
	MyHero = Hero{Role{name,1,15.0,15.0,6.0,1.0,true},0,0,0}
}

func (hero Hero) upLevel(){
	canLevel := false
	if hero.exp == hero.Level {
		canLevel = true
	}
	if canLevel{
		hero.exp = 0
		hero.Level++
		hero.Hp += 15
		hero.Chp = hero.Hp
		hero.Attake += 1
		hero.Defense += 1
		fmt.Println("恭喜您升级了，当前等级为：",hero.Level)
	}

}

func Newmonster(id int) Monster{
	var monster Monster
	monster = monsters[id]

	monster.Role.IsAlive = true;
	return monster
}

