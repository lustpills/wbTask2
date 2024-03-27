package pattern

import (
	"testing"

	"github.com/lustpill/wbTask2/pattern"
	//"github.com/lustpill/wbTask2/pattern"
)

func TestBuilderPattern(t *testing.T) {

	t.Log("Строитель\n")
	builder1 := pattern.GetBuilder("Company1")
	builder2 := pattern.GetBuilder("Company2")
	factory := pattern.NewFactory(builder1)

	product := factory.Construct()
	//product.Show()

	t.Log(product)
	t.Log("hello")

	factory.SetBuilder(builder2)

	product = factory.Construct()

	t.Log(product)

	// expectedPart1 := "Part 1"
	// if product.part1 != expectedPart1 {
	//  t.Errorf("Expected part1 to be %s, but got %s", expectedPart1, product.part1)
	// }

	// expectedPart2 := "Part 2"
	// if product.part2 != expectedPart2 {
	//  t.Errorf("Expected part2 to be %s, but got %s", expectedPart2, product.part2)
	// }
}
