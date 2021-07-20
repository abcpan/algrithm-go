package dayliyTemperature

func dailyTemperatures(temperatures []int) []int {
	var i = 0
	var j = 0
	for i < len(temperatures) && j < len(temperatures) {
		j++
		if j == len(temperatures){
			temperatures[i] = 0
			i++
			j = i
			continue
		}
		if temperatures[j] > temperatures[i] {
			temperatures[i] = j  - i
			i++
			j = i
		}
	}
	return temperatures
}