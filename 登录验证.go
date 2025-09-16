package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func isphone(number string) error {
	phone := `^1[3-9]\d{9}$`
	reg := regexp.MustCompile(phone)

	if !reg.MatchString(number) {
		return errors.New("请输入正确的手机号")

	} else {
		return nil

	}
}

func getcode() string {
	q := make([]string, 0)
	for w := 'A'; w <= 'Z'; w++ {
		q = append(q, string(w))
	}
	for i := 0; i <= 9; i++ {
		q = append(q, strconv.Itoa(i))
	}
	y := make([]string, 6)
	for r := 0; r < 6; r++ {
		y[r] = q[rand.Intn(len(q))]
	}
	return strings.Join(y, "")
}

type codesrecord struct {
	code  string
	form  time.Time
	used  bool
	times int
	date  string
}
type phonestore struct {
	data map[string]*codesrecord
}

var store = &phonestore{data: make(map[string]*codesrecord)}

func (s *phonestore) addcode(number string) (string, error) {
	now := time.Now()
	today := now.Format("2006-01-02")
	d, ok := s.data[number]
	if !ok {
		code := getcode()
		s.data[number] = &codesrecord{
			code:  code,
			form:  now,
			used:  false,
			times: 1,
			date:  today,
		}
		return code, nil
	}
	if d.date != today {
		d.times = 0
		d.date = today
	}
	if d.times == 5 {
		return "", errors.New("今日发送次数已达上限")
	}
	if now.Sub(d.form) < 60*time.Second {
		return "", errors.New("获取太频繁，请稍后再试")
	}
	code := getcode()
	d.code = code
	d.form = now
	d.used = false
	d.times++
	return code, nil

}
func (s *phonestore) testcode(number, code string) error {
	now := time.Now()
	d, ok := s.data[number]
	if !ok {
		return errors.New("请先获取验证码")
	}
	if d.code != code || d.used || now.Sub(d.form) > 5*time.Minute {
		return errors.New("验证码错误")
	}
	d.used = true
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n请输入手机号：")

		phone, _ := reader.ReadString('\n')
		phone = strings.TrimSpace(phone)
		err := isphone(phone)
		if err != nil {
			fmt.Println(err)
			continue
		}
		for {

			fmt.Println("请输入操作：1=使用验证码登录  2=获取验证码")
			t, _ := reader.ReadString('\n')
			t = strings.TrimSpace(t)

			switch t {
			case "2":
				code, err := store.addcode(phone)
				if err != nil {
					fmt.Println("获取失败：", err.Error())
					continue
				} else {
					fmt.Println(code)
					continue
				}

			case "1":
				fmt.Println("请输入验证码：")
				g, _ := reader.ReadString('\n')
				g = strings.TrimSpace(g)
				if err := store.testcode(phone, g); err != nil {
					fmt.Println("登陆失败：", err.Error())
					continue
				} else {
					fmt.Println("登录成功！")
				}
			}
			break
		}
		break
	}

}
