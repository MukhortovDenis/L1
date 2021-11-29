package main

import "fmt"

type user struct {
}

type computer interface {
	connectToWiFi()
}

type compWithModule struct {
}

type compWithoutModule struct {
}

type compWithoutModuleWithAdapter struct {
	comp *compWithoutModule
}

func (u *user) turnOnWiFi(comp computer) {
	fmt.Println("Вайфай включен")
	comp.connectToWiFi()
}

func (with *compWithModule) connectToWiFi() {
	fmt.Println("Компьютер с модулем подключен к вайфаю")
}

func (without *compWithoutModule) connectToWiFiWithountModule() {
	fmt.Println("Компьютер без модуля подключен к вайфаю")
}
func (withAdapter *compWithoutModuleWithAdapter) connectToWiFi() {
	fmt.Println("Вставляем сетевой адаптер Вайфая")
	withAdapter.comp.connectToWiFiWithountModule()
}

func main() {
	user := &user{}
	compWith := &compWithModule{}
	user.turnOnWiFi(compWith)
	fmt.Println("------------------------------------------------------------------------------")

	compWithout := &compWithoutModule{}

	compWithAdapter := &compWithoutModuleWithAdapter{
		comp:   compWithout,
	}
	user.turnOnWiFi(compWithAdapter)
}
