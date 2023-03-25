package api

import "github.com/akshaypatil3096/Hex-Architecture-GO/internal/ports"

type Adapter struct {
	db    ports.DbPort
	arith ports.ArithmeticPort
}

func NewAdapter(db ports.DbPort, arith ports.ArithmeticPort) *Adapter {
	return &Adapter{
		db:    db,
		arith: arith,
	}
}

func (apiAdapter Adapter) GetAddition(a, b int) (int, error) {
	ans, err := apiAdapter.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}

	err = apiAdapter.db.AddToHistory(ans, "Addition")
	if err != nil {
		return 0, err
	}
	return ans, nil
}

func (apiAdapter Adapter) GetSubtraction(a, b int) (int, error) {
	ans, err := apiAdapter.arith.Subtraction(a, b)
	if err != nil {
		return 0, err
	}

	err = apiAdapter.db.AddToHistory(ans, "Subtraction")
	if err != nil {
		return 0, err
	}
	return ans, nil
}

func (apiAdapter Adapter) GetMultiplication(a, b int) (int, error) {
	ans, err := apiAdapter.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}

	err = apiAdapter.db.AddToHistory(ans, "Multiplication")
	if err != nil {
		return 0, err
	}
	return ans, nil
}
func (apiAdapter Adapter) GetDivision(a, b int) (int, error) {
	ans, err := apiAdapter.arith.Division(a, b)
	if err != nil {
		return 0, err
	}

	err = apiAdapter.db.AddToHistory(ans, "Division")
	if err != nil {
		return 0, err
	}
	return ans, nil
}
