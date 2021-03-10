package main

//import "sync"
//
//type ParkingSystem1 struct {
//	m sync.Map
//}
//
//func Constructor1(big int, medium int, small int) ParkingSystem1 {
//	m:=sync.Map{}
//	m.Store(1,big)
//	m.Store(2,medium)
//	m.Store(3,small)
//	return ParkingSystem1{
//	m,
//	}
//}
//
//func (this *ParkingSystem1) AddCar1(carType int) bool {
//	if cnt,ok:=this.m.Load(carType);ok {
//		if cnt.(int)>0 {
//			this.m.Store(carType,cnt.(int)-1)
//			return true
//		}
//	}
//	return false
//}
