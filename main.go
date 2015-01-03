package main

import "time"
import "fmt"
import "os"
import "io/ioutil"
import "math/rand"
import "github.com/antonholmquist/jason"

var LOADED *jason.Object // global var to hold the entire json object from file

// loads from file into the jason.Object pointer.
func LoadRollTables() (err error) {
	fi, err := os.Open("./rolltables.json") // hardcoded input file
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	r, err := ioutil.ReadAll(fi) // read entire file at once, may need to change if files get large
	if err != nil {
		panic(err)
	}

	LOADED, err = jason.NewObjectFromBytes(r) // take the bytes object and turn into the jason object
	if err != nil {
		panic(err)
	}

	return err
}

// rolls all the tables in the jason.Object pointer LOADED.
func RollAllTables() (err error) {
	rtables, err := LOADED.GetObjectArray("rolltables")
	for _, value := range rtables {
		err = RollOneTable(value)
		if err != nil {
			fmt.Println(err) // just want to print, not panic so it keeps going through other rolls
		}
	}
	return
}

func RollOneTable(rt *jason.Object) (err error) {

	// initialize vars
	var i int64       // counter for the loop of rolls
	var rolls []int64 // holds a list of all rolls if needed later for debugging
	var total int64   // result amount
	var dnum int64    // how many dice to roll
	var dmod int64    // this is to add a single fixed modifier amount to the roll if you desire

	dnum, err = rt.GetInt64("Dicenum")
	if err != nil {
		panic(err)
	}

	dmod, err = rt.GetInt64("Dicemod")
	if err != nil {
		panic(err)
	}

	// generates a roll based off of dnum and dsize
	for i = 0; i < dnum; i++ {
		var dsize int64 // number of sides of dice, 1 being the lowest always

		dsize, err = rt.GetInt64("Dicesize")
		if err != nil {
			panic(err)
		}

		roll := rand.Int63n(dsize) + 1 // + 1 makes it 1-100 instead of 0-99
		rolls = append(rolls, roll)
		total += roll
	}

	// adds dmod to total from the rolls
	total += dmod

	rollsarray, err := rt.GetObjectArray("Rolls")
	if err != nil {
		panic(err)
	}

	// this portion of the function loads the inidividual rolls and then checks if the generated roll matches
	for _, individRolls := range rollsarray {

		var themin int64 // "Min" in json
		var themax int64 // "Max" in json

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

			var subrolls *jason.Object // place to store any subrolls

			subrolls, err = individRolls.GetObject("rolltable")
			if err != nil { // idea is if there's an error, there's no subrolls so just pass
			} else { // but, if err is nil, we call ourselves recursively
				err = RollOneTable(subrolls)
				if err != nil {
					fmt.Println(err) // don't panic as their might be a non-nil error with further sub rolls
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
