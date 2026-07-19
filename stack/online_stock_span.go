package stack

type Stock struct {
	price int
	span  int
}

type StockSpanner struct {
	stack []Stock
}

func StockConstructor() StockSpanner {
	return StockSpanner{stack: []Stock{}}
}

func (this *StockSpanner) Next(price int) int {

	newStock := Stock{price: price, span: 1}

	if len(this.stack) != 0 {
		for len(this.stack) != 0 && this.stack[len(this.stack)-1].price <= newStock.price {
			stock := this.stack[len(this.stack)-1]
			newStock.span += stock.span
			this.stack = this.stack[:len(this.stack)-1]
		}
	}

	this.stack = append(this.stack, newStock)
	return newStock.span
}
