// Package weather 天气系统.
package weather

var (
	// CurrentCondition 当前状态.
	CurrentCondition string
	// CurrentLocation 当前位置.
	CurrentLocation string
)

// Forecast 开始预测天气.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
