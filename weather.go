package rockweather

import "github.com/xfort/gohttp"

const ChinaWeather_CityData = "http://i.tq121.com.cn/j/wap2016/news/city_search_data.js"

type RockWeather struct {
	RWHttp *gohttp.GoHttp
}

//
func (rock *RockWeather) LoadCityWeather(cityName string, cityWeather *CityWeather) (*CityWeather, error) {

	return nil, nil
}

func (rock *RockWeather) LoadCityCode() () {

}
