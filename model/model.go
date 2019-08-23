package model

// RawData
type RawData struct {

	// Розіграш
	Game string

	//Дата проведення
	Date string

	//Лототрон
	Lototron string

	//Комплект кульок
	BallSet string

	//Кулька номер 1
	Ball1 string

	//Кулька номер 2
	Ball2 string

	//Кулька номер 3
	Ball3 string

	//Кулька номер 4
	Ball4 string

	//Кулька номер 5
	Ball5 string

	//Кулька номер 6
	Ball6 string

	//переможців які вгадали 2 кульки
	Ball2Winners string

	//приз за 2 кульки
	Ball2Price string

	//переможців які вгадали 3 кульки
	Ball3Winners string

	//приз за 3 кульки
	Ball3Price string

	//переможців які вгадали 4 кульки
	Ball4Winners string

	//приз за 4 кульки
	Ball4Price string

	//переможців які вгадали 5 кульок
	Ball5Winners string

	//приз за 5 кульок
	Ball5Price string

	//переможців які вгадали 6 кульок
	Ball6Winners string

	//приз за 6 кульок
	Ball6Price string
}

// PreparedData
type PreparedData struct {
	Game         int
	Date         string
	Lototron     string
	BallSet      int
	Ball1        int
	Ball2        int
	Ball3        int
	Ball4        int
	Ball5        int
	Ball6        int
	Ball2Winners int
	Ball2Price   int
	Ball3Winners int
	Ball3Price   int
	Ball4Winners int
	Ball4Price   int
	Ball5Winners int
	Ball5Price   int
	Ball6Winners int
	Ball6Price   int
}
