package validator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func Check(obj interface{}) error {
	dist := map[string]map[string]string{}

	t := reflect.TypeOf(obj)
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("validator")
		if tag != "" {
			//拆分.
			field := t.Field(i).Name

			mp := map[string]string{}
			arr := strings.Split(tag, ";")
			for _, v := range arr {
				//正则表达式需要单独处理一下.
				pattern := "^(expression:)"
				res, _ := regexp.Match(pattern, []byte(v))
				if res {
					expression := strings.Replace(v, "expression:", "", -1)
					mp["expression"] = expression
				} else {
					item := strings.Split(v, ":")
					mp[item[0]] = item[1]
				}
			}

			dist[field] = mp
		}
	}

	return valid(dist, obj)
}

func valid(mp map[string]map[string]string, obj interface{}) error {
	v := reflect.ValueOf(obj)
	t := reflect.TypeOf(obj)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i).Name

		mpData := mp[field]

		// 获取对应的类型.
		types, ok := mpData["type"]
		if !ok {
			continue
		}

		//字符串.
		if types == "string" {
			value := v.Field(i).Interface().(string)

			//检查必填.
			requiredVal, ok := mpData["required"]
			if ok && requiredVal == "true" {
				err := checkStringRequired(field, value)
				if err != nil {
					return err
				}
			}

			//检查长度，是否必填.
			if requiredVal == "true" || value != "" {
				length, ok := mpData["length"]
				if ok {
					err := checkStringLength(length, mpData["name"], value)
					if err != nil {
						return err
					}
				}
			}
		}

		//int.
		if types == "int" {
			value := v.Field(i).Interface().(int)
			minVal, ok := mpData["min"]
			if ok {
				min := parseint(minVal)
				if value < min {
					return errors.New(fmt.Sprintf("%s不能小于%d", mpData["name"], min))
				}
			}

			maxVal, ok := mpData["max"]
			if ok {
				max := parseint(maxVal)
				if value > max {
					return errors.New(fmt.Sprintf("%s不能大于%d", mpData["name"], max))
				}
			}
		}

		//正则验证.
		if types == "regexp" {
			value := v.Field(i).Interface().(string)

			//检查必填.
			requiredVal, ok := mpData["required"]
			if ok && requiredVal == "true" {
				err := checkStringRequired(field, value)
				if err != nil {
					return err
				}
			}

			//检查正则，是否必填.
			if requiredVal == "true" || value != "" {
				pattern, ok := mpData["expression"]
				if !ok {
					return errors.New(fmt.Sprintf(`%s格式不正确`, mpData["name"]))
				}
				if len(pattern) < 1 {
					return errors.New(fmt.Sprintf(`%s格式不正确`, mpData["name"]))
				}
				reg, err := regexp.Compile(pattern)
				if err != nil {
					return errors.New(fmt.Sprintf(`%s格式不正确`, mpData["name"]))
				}
				regVal := reg.FindString(value)
				if regVal != value {
					return errors.New(fmt.Sprintf(`%s格式不正确`, mpData["name"]))
				}
			}

		}

	}

	return nil
}

//checkStringRequired 检查是否必填.
func checkStringRequired(name, value string) error {
	if len(value) < 1 {
		return errors.New(fmt.Sprintf("%s不能为空", name))
	}
	return nil
}

//checkStringLength 检查字符串的长度是否合法.
func checkStringLength(length string, name, val string) error {
	var min, max int
	length = strings.Replace(length, "[", "", -1)
	length = strings.Replace(length, "]", "", -1)
	arr := strings.Split(length, "-")
	if len(arr) == 2 {
		min = parseint(arr[0])
		max = parseint(arr[1])
	}

	if len([]rune(val)) < min {
		return errors.New(fmt.Sprintf("%s的长度不能小于%d位", name, min))
	}
	if len([]rune(val)) > max {
		return errors.New(fmt.Sprintf("%s的长度不能大于位%d", name, max))
	}

	return nil
}

func parseint(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
