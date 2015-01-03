package main

import "time"
import "fmt"
import "os"
import "io/ioutil"
import "math/rand"
import "github.com/antonholmquist/jason"

var LOADED *jason.Object

func LoadRollTables() (err error) {
	fi, err := os.Open("./rolltables.json")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	r, err := ioutil.ReadAll(fi)
	if err != nil {
		panic(err)
	}

	LOADED, err = jason.NewObjectFromBytes(r)
	if err != nil {
		panic(err)
	}

	return err
}

func RollAllTables() (err error) {
	rtables, err := LOADED.GetObjectArray("rolltables")
	for _, value := range rtables {
		err = RollOneTable(value)
		if err != nil {
			fmt.Println(err)
		}
	}
	return
}

func RollOneTable(rt *jason.Object) (err error) {

	var i int64
	var rolls []int64
	var total int64
	var dnum int64

	dnum, err = rt.GetInt64("Dicenum")
	if err != nil {
		panic(err)
	}

	for i = 0; i < dnum; i++ {
		var dsize int64

		dsize, err = rt.GetInt64("Dicesize")
		if err != nil {
			panic(err)
		}

		roll := rand.Int63n(dsize) + 1 // + 1 makes it 1-100 instead of 0-99
		rolls = append(rolls, roll)
		total += roll
	}

	var dmod int64
	dmod, err = rt.GetInt64("Dicemod")
	if err != nil {
		panic(err)
	}

	total += dmod
	rollsarray, err := rt.GetObjectArray("Rolls")
	if err != nil {
		panic(err)
	}

	for _, individRolls := range rollsarray {

		var themin int64
		var themax int64

		themin, err = individRolls.GetInt64("Min")
		if err != nil {
			panic(err)
		}

		themax, err = individRolls.GetInt64("Max")
		if err != nil {
			panic(err)
		}

		if total >= themin && total <= themax {

			var result string
			result, err = individRolls.GetString("Result")
			if err != nil {
				panic(err)
			}

			var name string
			name, err = rt.GetString("Name")
			if err != nil {
				panic(err)
			}

			fmt.Printf("%s: %s\n", name, result)

			var subrolls *jason.Object

			subrolls, err = individRolls.GetObject("rolltable")
			if err != nil {
			} else {
				err = nil
				err = RollOneTable(subrolls)
				if err != nil {
					fmt.Println(err)
					return
				}
			}

		}
	}
	err = nil

	return
}

func main() {

	rand.Seed(time.Now().UnixNano()) // sets a unique seed for the random number generator

	LoadRollTables() // loads the json into a *jason.Object global
	RollAllTables()  // rolls it all
}
