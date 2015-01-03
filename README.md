# Welcome to dm-roller

`dm-roller` is a script written in Go to roll many different roll tables with one shell command. Perhaps you
have a lot of roll tables a la a modern DMG or other RPG rule book. Or perhaps you have your own customized
very detailed roll tables that you want to put into json form and be able to run a simple script? This is the
script for you. 

## Features

Very simple main program interprets the `rolltables.json` file for all of the rolls. You can use `\n` and `\t`
to format names for output in the JSON as well. It allows dice ranges as well as multiple dice rolled as well
as a simple modifier for 1d20+1 type rolls. 

As of now, the modifier is hard coded into the json file. And, all rolls are performed when the script runs. 

In the future, the script will allow choices upon running or use flags to allow categories; i.e. dungeon
generation or treasure drop choices. 

## Please no copyrighted rolltables added

Please do not add any roll tables that would violate the copyright of any party.

## Installation

After you have a working Go installation and set up your $GOPATH correctly:

    $ go get github.com/jrmiller82/dm-roller
    $ cd $GOPATH/src/github.com/jrmiller82/dm-roller
    $ go build
    $ ./dm-roller

If you have your `$GOPATH` properly set up where $GOPATH/bin is in your $PATH, you go change `go build` above
to `go install` and then call `dm-roller` from anywhere on the command line.

## Example


    {
        "rolltables": [
            {
                "Dicemod": 0,
                "Dicenum": 1,
                "Dicesize": 10,
                "Name": "Form of Government",
                "Rolls": [
                    {
                        "Max": 3,
                        "Min": 1,
                        "Result": "Autocracy"
                    },
                    {
                        "Max": 4,
                        "Min": 4,
                        "Result": "Bureaucracy",
                        "rolltable": {
                            "Dicemod": 0,
                            "Dicenum": 1,
                            "Dicesize": 4,
                            "Name": "\tType of Bureaucracy",
                            "Rolls": [
                                {
                                    "Max": 3,
                                    "Min": 1,
                                    "Result": "HHGTTG"
                                },
                                {
                                    "Max": 4,
                                    "Min": 4,
                                    "Result": "USA"
                                }
                            ]
                        }
                    },
                    {
                        "Max": 10,
                        "Min": 5,
                        "Result": "Confederacy"
                    }
                ]
            },
            {
                "Dicemod": 0,
                "Dicenum": 1,
                "Dicesize": 6,
                "Name": "Un-nested Roll",
                "Rolls": [
                    {
                        "Max": 4,
                        "Min": 1,
                        "Result": "you rolled a 1 through 4"
                    },
                    {
                        "Max": 6,
                        "Min": 5,
                        "Result": "you rolled a 5 or 6"
                    }
                ]
            }
        ]
    }


Given the above `rolltables.json` file, running the command will output:

    $ go run main.go
    Form of Government: Autocracy
    Un-nested Roll: you rolled a 1 through 4

If Bureaucracy is hit on the first roll, the Type of Bureaucracy roll will occur:

    $ go run main.go
    Form of Government: Bureaucracy
        Type of Bureaucracy: HHGTTG
    Un-nested Roll: you rolled a 5 or 6
    $

The results will be different based on the random rolls.

## Adding tables


You only need to edit the file `rolltables.json`. As of now, it is auto loaded by the script. Please don't
change the name of the file either as that will cause the script to panic.

Please note that you can *nest* rolls almost as far as you would like. The function is recursive in that
regard. I may clean up the json format in the future but it should not be a large change; it will likely be
the addition of some sort of category or other filtering tags so that not every single roll is rolled on
script call.

A good example of nesting is found in the *bureacracy* section of `rolltables.json` and shows the proper way
to nest rolls. The nested roll will only be called if bureaucracy was the result of the parent roll.
