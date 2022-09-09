package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomDate()time.Time{
	const date_format="2006-01-02"
	min,_:=time.Parse(date_format,"1999-01-01")
	max,_:=time.Parse(date_format,"2004-12-31")
	delta:=max.Unix()-min.Unix()
	sec:=rand.Int63n(delta)+min.Unix()
	return time.Unix(sec,0)
}