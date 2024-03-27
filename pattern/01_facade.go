// за фасадом скрываются 2 подсистемы, когда мы выполняем Operation,
// на самом деле выполняются так же операции для подсистему
// примером может служить покупка товара в магазине по карте с запросом баланса в банке

package pattern

import "fmt"

// SubsystemA ...
type SubsystemA struct{}

func (s *SubsystemA) OperationA() {
	fmt.Println("SubsystemA: OperationA")
}

// SubsystemB ...
type SubsystemB struct{}

func (s *SubsystemB) OperationB() {
	fmt.Println("SubsystemB: OperationB")
}

// Facade ...
type Facade struct {
	subsystemA *SubsystemA
	subsystemB *SubsystemB
}

func NewFacade() *Facade {
	return &Facade{
		subsystemA: &SubsystemA{},
		subsystemB: &SubsystemB{},
	}
}

func (f *Facade) Operation() {
	fmt.Println("Facade: initializing subsystems")
	f.subsystemA.OperationA()
	f.subsystemB.OperationB()
}
