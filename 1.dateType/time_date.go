package main

import (
	"fmt"
	"time"
)

var week time.Duration

func main() {
	//值为win10 zh 环境下
	t := time.Now() //当前本地时间
	//t = time.Now().UTC()//UTC时间
	fmt.Println(t)                                              // 2019-11-21 09:19:25.7051102 +0800 CST m=+0.002978101 CST中国标准时间
	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year()) // 20.11.2019

	fmt.Println(t.Nanosecond()) //纳秒 415148800
	fmt.Println(t.UnixNano())   //unix时间撮+纳秒 1574299364415148800
	fmt.Println(t.Unix())       //unix时间撮 1574299364

	fmt.Printf("%v\n", t.String()) //2019-11-21 09:28:57.8300647 +0800 CST m=+0.003999301 和fmt.Println(t)一样
	zone, offset := t.Zone()
	fmt.Printf("zone:%v;offset:%d\n", zone, offset) //zone:CST;offset:28800
	fmt.Printf("%d\n", t.Weekday())                 // Sunday=0 Monday~Saturday=1~6

	t = time.Now().UTC()
	fmt.Println(t.Local()) // 会转化为当地时间 2019-11-21 09:33:23.1444786 +0800 CST
	fmt.Println(t)         //2019-11-21 01:33:23.1444786 +0000 UTC
	//t = t.Local() //改为当地时间
	fmt.Println(time.Now()) // 2019-11-20 16:37:52.4243357 +0800 CST m=+0.020002401
	// calculating times:
	week = 60 * 60 * 24 * 7 * 1e9 // must be in nanosec 毫秒*1000 微妙*1000 纳秒*1000 即 1e9 = 1 000 000 000
	weekFromNow := t.Add(week)
	fmt.Println(weekFromNow) // 2019-11-28 01:33:23.1444786 +0000 UTC 加了7天 计算时是已UTC时间计算的 打印出来也是UTC
	// formatting times:
	fmt.Println(t.Format(time.RFC822))       // 20 Nov 19 08:37 UTC
	fmt.Println(t.Format(time.ANSIC))        // Wed Nov 20 08:37:52 2019
	fmt.Println(t.Format(time.RFC3339))      // 2011 Dec 201111 08:5220
	s := t.Format("2006-01-02 15:04:05.999") //月 日 时 分 秒年为123456 微妙等为9 类似Ymd只是这里是数字代替
	fmt.Println(t, "=>", s)                  //2019-11-21 01:38:22.4232006 +0000 UTC => 2019-11-21 01:38:22.423
}
