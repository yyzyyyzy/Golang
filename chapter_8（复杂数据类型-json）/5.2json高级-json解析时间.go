package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Date struct {
	time.Time
}

func (d *Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, d.Format("2006-01-02"))), nil
}

func (d *Date) UnmarshalJSON(b []byte) error {
	var err error
	// 指定时区
	d.Time, err = time.ParseInLocation(`"2006-01-02"`, string(b), time.Local)
	if err != nil {
		return err
	}
	return nil
}

type Datetime struct {
	time.Time
}

func (d *Datetime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%s", d.Format("2006-01-02 15:04:05"))), nil
}

func (d *Datetime) UnmarshalJSON(b []byte) error {
	var err error
	d.Time, err = time.Parse(`"2006-01-02 15:04:05"`, string(b))
	if err != nil {
		return err
	}
	return nil
}

type Amazing struct {
	D Date `json:"data"`
}

func main() {
	demo := &Amazing{}
	_ = json.Unmarshal([]byte(`{"data":"2022-02-14"}`), demo) //json反序列化为时间标准格式
	fmt.Println(demo)                                         // &{2021-03-30 00:00:00 +0800 CST}

	demo2 := &Amazing{D: Date{time.Now()}}
	v, _ := json.Marshal(demo2) //时间标准格式序列化为json
	fmt.Println(string(v))      // {"data":"2021-03-30"}
}
