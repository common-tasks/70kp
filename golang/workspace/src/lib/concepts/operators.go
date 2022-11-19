package concepts
type Celcius float32
type Fahrenheit float32
func convertTemp(centigrade Celcius) Fahrenheit {
	var ft Fahrenheit = Fahrenheit((centigrade*9)/5 + 32)
	return ft
}