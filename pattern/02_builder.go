package pattern

import "fmt"

const (
	B1 = "Company1"
	B2 = "Company2"
)

// Product ...
type Product struct {
	part1 string
	part2 string
}

func (p *Product) Show() {
	fmt.Println(p.part1, p.part2)
}

// Builder ...
type Builder interface {
	BuildPart1()
	BuildPart2()
	GetProduct() *Product
}

func GetBuilder(BuilderType string) Builder {
	switch BuilderType {
	default:
		return nil
	case B1:
		return NewCompany1_builder()
	case B2:
		return NewCompany2_builder()
	}
}

// ConcreteBuilder ...
type Company1_builder struct {
	product *Product
}

type Company2_builder struct {
	product *Product
}

func NewCompany1_builder() *Company1_builder {
	return &Company1_builder{product: &Product{}}
}

func (b *Company1_builder) BuildPart1() {
	b.product.part1 = "Part 1 by company1"
	fmt.Println("Building Part 1")
}

func (b *Company1_builder) BuildPart2() {
	b.product.part2 = "Part 2 by company1"
	fmt.Println("Building Part 2")
}

func (b *Company1_builder) GetProduct() *Product {
	return b.product
}

func NewCompany2_builder() *Company2_builder {
	return &Company2_builder{product: &Product{}}
}

func (b *Company2_builder) BuildPart1() {
	b.product.part1 = "Part 1 by company2"
	fmt.Println("Building Part 1")
}

func (b *Company2_builder) BuildPart2() {
	b.product.part2 = "Part 2 by company2"
	fmt.Println("Building Part 2")
}

func (b *Company2_builder) GetProduct() *Product {
	return b.product
}

// Director ...
type Factory struct {
	builder Builder
}

func NewFactory(builder Builder) *Factory {
	return &Factory{builder: builder}
}

func (factory *Factory) SetBuilder(builder Builder) {
	factory.builder = builder
}

func (d *Factory) Construct() *Product {
	d.builder.BuildPart1()
	d.builder.BuildPart2()
	return d.builder.GetProduct()
}
