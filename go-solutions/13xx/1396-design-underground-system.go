package _3xx

/**
https://leetcode.com/problems/design-underground-system/submissions/
1396. Design Underground System
Medium
*/
type Average struct {
	Average float64
	Nums    float64
}
type UserOnStation struct {
	Station string
	Time    int
}
type UndergroundSystem struct {
	stations map[string]Average
	users    map[int]UserOnStation
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		stations: make(map[string]Average),
		users:    make(map[int]UserOnStation),
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.users[id] = UserOnStation{Station: stationName, Time: t}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	diff := t - this.users[id].Time
	key := this.users[id].Station + "-" + stationName
	delete(this.users, id)

	if _, ok := this.stations[key]; !ok {
		this.stations[key] = Average{
			Average: float64(diff),
			Nums:    1,
		}
		return
	}
	average := this.stations[key]
	sum := average.Average * average.Nums
	sum += float64(diff)
	average.Nums += 1
	average.Average = sum / average.Nums
	this.stations[key] = average
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	return this.stations[startStation+"-"+endStation].Average
}

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
