package main

import "fmt"

type User struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

func NewUser1(id int, name string, gender string, age int, phone string, email string) User {
	return User{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

func NewUser2(name string, gender string, age int, phone string, email string) User {
	return User{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

func (this User) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t", this.Id, this.Name, this.Gender, this.Age, this.Phone, this.Email)
	return info
}

// 结构体完成增删改查的操作
type UserService struct {
	users   []User
	UserNum int
}

func NewUserService() *UserService {
	userService := &UserService{}
	userService.UserNum = 1
	user1 := NewUser1(1, "李子康", "男", 18, "13610850940", "916990143@qq.com")
	userService.users = append(userService.users, user1)
	return userService
}

// 查询用户（切片方法）
func (this *UserService) List() []User {
	return this.users
}

// 增加用户（切片方法）
func (this *UserService) Add(user User) bool {
	this.UserNum++
	user.Id = this.UserNum

	this.users = append(this.users, user)
	return true
}

// 删除用户（切片方法）
func (this *UserService) Delete(id int) bool {
	index := this.FindById(id)
	if index == -1 {
		return false
	} else {
		this.users = append(this.users[:index], this.users[index+1:]...)
		return true
	}
}

// 根据id查找客户在切片中对应的下标，如果没有客户，那么返回-1
func (this *UserService) FindById(id int) int {
	index := -1

	for i := 0; i < len(this.users); i++ {
		if this.users[i].Id == id {
			index = i
		}
	}
	return index
}

type userView struct {
	key  string //接收用户输入
	loop bool   //循环显示主菜单

	userService *UserService
}

func (this *userView) list() {
	users := this.userService.List()

	fmt.Println("---------------------------客户列表---------------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")

	for i := 0; i < len(users); i++ {
		fmt.Println(users[i].GetInfo())
	}

	fmt.Printf("\n-------------------------客户列表完成-------------------------\n\n")
}

func (this *userView) add() {
	fmt.Println("---------------------添加客户---------------------")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱：")
	email := ""
	fmt.Scanln(&email)

	user := NewUser2(name, gender, age, phone, email)
	if this.userService.Add(user) {
		fmt.Println("---------------------添加完成---------------------")
	} else {
		fmt.Println("---------------------添加失败---------------------")
	}
}

func (this *userView) delete() {
	fmt.Println("---------------------删除客户---------------------")
	fmt.Print("请选择待删除客户编号(-1退出)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return //放弃删除操作
	}
	fmt.Println("确认是否删除(Y/N)：")
	//这里同学们可以加入一个循环判断，直到用户输入 y 或者 n,才退出..
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y" {
		//调用customerService 的 Delete方法
		if this.userService.Delete(id) {
			fmt.Println("---------------------删除完成---------------------")
		} else {
			fmt.Println("---------------------删除失败，输入的id号不存在----")
		}
	}
}

func (this *userView) exit() {
	fmt.Println("是否退出(Y/N)")
	for {
		fmt.Scanln(&this.key)
		if this.key == "Y" || this.key == "y" || this.key == "N" || this.key == "n" {
			break
		}
		fmt.Println("你的输入有误，请重新输入")
	}

	if this.key == "Y" || this.key == "y" {
		this.loop = false
	}
}

func (this *userView) mainMenu() {

	for {
		fmt.Println("-----------------客户信息管理软件-----------------")
		fmt.Println("                 1 添 加 客 户")
		fmt.Println("                 2 修 改 客 户")
		fmt.Println("                 3 删 除 客 户")
		fmt.Println("                 4 客 户 列 表")
		fmt.Println("                 5 退       出")
		fmt.Print("请选择(1-5)：")

		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.add()
		case "2":
			fmt.Println("修 改 客 户")
		case "3":
			this.delete()
		case "4":
			this.list()
		case "5":
			this.exit()
		default:
			fmt.Println("你的输入有误，请重新输入...")
		}

		if !this.loop {
			break
		}

	}
	fmt.Println("已退出了客户关系管理系统...")
}

func main() {
	userView := userView{
		key:         "",
		loop:        true,
		userService: nil,
	}
	userView.userService = NewUserService()
	userView.mainMenu()
}
