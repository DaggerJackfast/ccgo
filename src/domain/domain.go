package domain

type CustomerRepository interface {
	Store(customer Customer)
	FindBy(id int) Customer
}

type ItemRepository interface {
	Store(item Item)
	FindById(id int) Item
}

type OrderRepository interface {
	Store(order Order)
	FindById(id int) Order
}

type Customer struct {
	Id   int
	Name string
}

type Item struct {
	Id        int
	Name      string
	Value     float64
	Available bool
}

type Order struct {
	Id       int
	Customer Customer
	Items    []Item
}

func (order *Order) Add(item Item) error {
	if !item.Available {
		return errors.New("Cannot add unavailable items to order")
	}
	if order.value()+item.Value > 250 {
		return errors.New("An order may not exceed a total value of $250.0")
	}
	order.Items = append(order.Items, item)
	return nil
}

func (order *Order) value() float64 {
	sum := 0.0
	for i:=range order.Items{
		sum = sum + order.Items[i].Value
	}
	return sum
}

