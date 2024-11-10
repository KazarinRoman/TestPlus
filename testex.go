package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	a := bufio.NewScanner(os.Stdin)
	_ = a.Scan()
	num := a.Text()

	str := strings.Split(num, " ")

	var w []string
	var t []string

	for k := 0; k < len(str); k++ {
		if str[k] == "-" || str[k] == "+" || str[k] == "/" || str[k] == "*" {
			w = append(w, strings.Join(t, " "))
			w = append(w, str[k])
			t = t[:0]
		} else {
			t = append(t, str[k])
			str[k] = ""
		}

		if k == len(str)-1 {
			w = append(w, strings.Join(t, " "))
			w = append(w, str[k])
			t = t[:0]
		}
	}

	w = w[:3]

	wq := strings.Split(w[0], "")
	if string(wq[0]) == "\"" && string(wq[len(wq)-1]) == "\"" {
	} else {
		panic("Error: wrong variable 1")
	}

	if w[1] == "-" || w[1] == "+" || w[1] == "/" || w[1] == "*" {
	} else {
		panic("Error: wrong operation")
	}

	if len(w[0]) > 12 || len(w[2]) > 12 {
		panic("Error: incorrect input of values")
	}

	var f []string
	f = append(f, "\"")
	switch w[1] {
	case "*":
		ww := strings.Split(w[2], "")
		if string(ww[0]) == "\"" || string(ww[len(ww)-1]) == "\"" {
			panic("Error: wrong variable 2")
		}

		for k := 0; k < len(w); k++ {
			w[k] = strings.Trim(w[k], "\"")
		}

		q, _ := strconv.Atoi(w[2])

		if q > 10 || q < 1 {
			panic("Error: wrong variable 2")
		}

		for i := 1; i <= q; i++ {
			f = append(f, w[0])
		}
	case "+":
		ww := strings.Split(w[2], "")
		if string(ww[0]) == "\"" && string(ww[len(ww)-1]) == "\"" {
		} else {
			panic("Error: wrong variable 2")
		}

		for k := 0; k < len(w); k++ {
			w[k] = strings.Trim(w[k], "\"")
		}

		f = append(f, w[0])
		f = append(f, w[2])
	case "-":
		ww := strings.Split(w[2], "")
		if string(ww[0]) == "\"" && string(ww[len(ww)-1]) == "\"" {
		} else {
			panic("Error: wrong variable 2")
		}

		for k := 0; k < len(w); k++ {
			w[k] = strings.Trim(w[k], "\"")
		}

		f = append(f, strings.Replace(w[0], w[2], "", 1))
	case "/":
		ww := strings.Split(w[2], "")
		if string(ww[0]) == "\"" || string(ww[len(ww)-1]) == "\"" {
			panic("Error: wrong variable 2")
		}

		for k := 0; k < len(w); k++ {
			w[k] = strings.Trim(w[k], "\"")
		}

		v := w[0]
		l := len(w[0])
		q, _ := strconv.Atoi(w[2])

		if q > 10 || q < 1 {
			panic("Error: wrong variable 2")
		}

		l = l / q
		for i := 0; i < l; i++ {
			f = append(f, string(v[i]))
		}
	default:
	}

	f = append(f, "\"")
	fj := strings.Join(f, "")

	for k := 0; k < len(fj); k++ {
		if k <= 40 {
			fmt.Print(string(fj[k]))
		} else {
			if k == 41 && len(fj) == 42 {
				fmt.Print("\"")
				break
			} else {
				fmt.Print("...\"")
				break
			}
		}
	}

}
