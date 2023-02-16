package models

type Car struct {
	CarType      string  `json:"car_type"`
	Engine       string  `json:"engine"`
	Fuel         string  `json:"fuel"`
	KmDriven     float64 `json:"km_driven"`
	MaxPower     string  `json:"max_power"`
	Mileage      string  `json:"mileage"`
	Name         string  `json:"name"`
	Owner        string  `json:"owner"`
	Score        float64 `json:"score"`
	SellerType   string  `json:"seller_type"`
	SellingPrice string  `json:"selling_price"`
	Torque       string  `json:"torque"`
	Transmission string  `json:"transmission"`
	Year         float64 `json:"year"`
}
