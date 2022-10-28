package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := make(map[string]int)
	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, ok := units[unit]
	if !ok {
		return false
	}
	_, ok = bill[item]
	switch ok {
	case false:
		bill[item] = units[unit]
	case true:
		bill[item] += units[unit]
	}
	return true
}

// RemoveItem removes an item from customer bill.
// billOk holds the bool result of bill item existance
// unitOk holds the bool value of units unit existance
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, billOk := bill[item]
	_, unitOk := units[unit]
	switch {
	case !billOk || !unitOk:
		return false
	case bill[item]-units[unit] < 0:
		return false
	case bill[item]-units[unit] == 0:
		delete(bill, item)
	default:
		bill[item] -= units[unit]
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
// element holds the numerical value of item
func GetItem(bill map[string]int, item string) (int, bool) {
	element, ok := bill[item]
	if !ok {
		return element, ok
	}
	return element, ok
}
