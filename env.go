package env

import (
	"reflect"
	"strings"
	"log"
	"strconv"
	"os"
	"fmt"
	"time"
)

var (
	DEFAULT_SEPARATOR = "_"
)

func upper(v string) string {
	return strings.ToUpper(v)
}

func Fill(v interface{}) error {
	if reflect.ValueOf(v).Kind() != reflect.Ptr {
		return fmt.Errorf("Fill只接受指针类型的值")
	}
	//for _, v := range os.Environ() {
	//	if strings.HasPrefix(v, "CONFIG") {
	//		log.Print(v)
	//	}
	//}
	ind := reflect.Indirect(reflect.ValueOf(v))
	prefix := upper(ind.Type().Name())
	err := fill(prefix, ind)
	if err != nil {
		return err
	}
	return nil
}

func combine(p, n string, v string, ok bool) string{
	if !ok {
		return p + DEFAULT_SEPARATOR + n
	}
	return p + v + n
}

func parseBool(v string) (bool, error) {
	if v == "" {
		return false, nil
	}
	return strconv.ParseBool(v)
}

func fill(pf string, ind reflect.Value) error{
	for i := 0; i < ind.NumField(); i++ {
		f := ind.Type().Field(i)
		name := f.Name
		envName, exist := f.Tag.Lookup("env")
		if exist {
			name = envName
		}
		s, exist := f.Tag.Lookup("sep")
		p := combine(pf, upper(name), s, exist)
		switch ind.Field(i).Kind() {
		case reflect.Struct:
			err := fill(p, ind.Field(i))
			if err != nil {
				return err
			}
		default:
			err := parse(p, ind.Field(i), f)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func parse(prefix string, f reflect.Value, sf reflect.StructField) error {
	log.Print("parse:", prefix, f.String(), f.Type().String(), f.Kind().String())
	df := sf.Tag.Get("default")
	isRequire, err := parseBool(sf.Tag.Get("require"))
	if err != nil {
		return fmt.Errorf("字段:%s require不是合法的布尔值，支持: 1 t T true TRUE True.", prefix)
	}
	ev, exist := os.LookupEnv(prefix)

	if !exist && isRequire {
		return fmt.Errorf("%s 是必须的变量，但没有被设置", prefix)
	}
	if !exist && df != "" {
		ev = df
	}
	log.Print("ev:", ev)
	switch f.Kind() {
	case reflect.String:
		f.SetString(df)
	case reflect.Int:
		iv, err := strconv.ParseInt(ev, 10, 32)
		if err != nil {
			return err
		}
		f.SetInt(iv)
	case reflect.Int64:
		if f.Type().String() == "time.Duration" {
			t, err := time.ParseDuration(ev)
			if err != nil {
				return err
			}
			f.Set(reflect.ValueOf(t))
		} else {
			iv, err := strconv.ParseInt(ev, 10, 64)
			if err != nil {
				return err
			}
			f.SetInt(iv)
		}
	case reflect.Uint:
		uiv, err := strconv.ParseUint(ev, 10, 32)
		if err != nil {
			return err
		}
		f.SetUint(uiv)
	case reflect.Uint64:
		uiv, err := strconv.ParseUint(ev, 10, 64)
		if err != nil {
			return err
		}
		f.SetUint(uiv)
	case reflect.Float32:
		f32, err := strconv.ParseFloat(ev,  32)
		if err != nil {
			return err
		}
		f.SetFloat(f32)
	case reflect.Float64:
		f64, err := strconv.ParseFloat(ev,  64)
		if err != nil {
			return err
		}
		f.SetFloat(f64)
	case reflect.Bool:
		b, err := parseBool(ev)
		if err != nil {
			return err
		}
		f.SetBool(b)
	}
	return nil
}