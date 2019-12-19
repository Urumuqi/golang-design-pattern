package facade

import (
	"fmt"
	"log"
)

func NewAPI() API {
	return &apiImpl{
		a: NewAModuleAPI2(),
		b: NewBModuleAPI(),
	}
}

//API is facade interface of facade package
type API interface {
	Test() string
}

//facade implement
type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

func (a *apiImpl) Test() string {
	aRet := a.a.TestA()
	bRet := a.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

//NewAModuleAPI return new AModuleAPI
func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

//AModuleAPI ...
type AModuleAPI interface {
	TestA() string
}

type aModuleImpl2 struct{}

func (*aModuleImpl2) TestA() string {
	log.Println("aModuleImpl2 is running")
	return "A module impl2 running"
}

func NewAModuleAPI2() AModuleAPI {
	return &aModuleImpl2{}
}

type aModuleImpl struct{}

func (*aModuleImpl) TestA() string {
	log.Println("TestA is running")
	return "A module running"
}

//NewBModuleAPI return new BModuleAPI
func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}

//BModuleAPI ...
type BModuleAPI interface {
	TestB() string
}

type bModuleImpl struct{}

func (*bModuleImpl) TestB() string {
	log.Println("TestB is running")
	return "B module running"
}
