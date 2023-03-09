package factorymethod

import "fmt"

// In this implementation, we define two types of products:
// ConcreteProductA and ConcreteProductB. Both of these products
// implement the Product interface, which has a single method GetName.

// We also define two types of creators: ConcreteCreatorA and
// ConcreteCreatorB. Both of these creators implement the Creator
// interface, which has a single method FactoryMethod that returns a
// Product.

// In main, we create instances of ConcreteCreatorA and
// ConcreteCreatorB and then call their FactoryMethod functions. These
// functions return products of type ConcreteProductA and
// ConcreteProductB, respectively, which we then print the name of.

// This code demonstrates how the factory method pattern can be used
// to create products of different types without the client code
// needing to know about the concrete classes of those products. The
// choice of which type of product to create is delegated to the
// creator, which can be swapped out as needed.

type Product interface {
	GetName() string
}

// A
type ConcreteProductA struct{}

func (*ConcreteProductA) GetName() string {
	return "ConcreteProductA"
}

// B
type ConcreteProductB struct{}

func (*ConcreteProductB) GetName() string {
	return "ConcreteProductB"
}

// Creator
type Creator interface {
	FactoryMethod() Product
}

// Creator A
type ConcreteCreatorA struct{}

func (c *ConcreteCreatorA) FactoryMethod() Product {
	return &ConcreteProductA{}
}

// Creator B
type ConcreteCreatorB struct{}

func (c *ConcreteCreatorB) FactoryMethod() Product {
	return &ConcreteProductB{}
}

func Run() {
	creatorA := &ConcreteCreatorA{}
	productA := creatorA.FactoryMethod()
	fmt.Println(productA.GetName())

	creatorB := &ConcreteCreatorB{}
	productB := creatorB.FactoryMethod()
	fmt.Println(productB.GetName())
}
