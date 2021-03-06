package main

import (
	"fmt"
	"math/rand"
)

var nouns = []string{
	"joker",
	"magneto",
	"durden",
	"monkey",
	"megatron",
	"ultron"}

var adjectives = []string{
	"emboldened",
	"impatient",
	"bloodthirsty",
	"mischievious",
	"ambivolent",
	"wildeyed",
	"shorn",
	"unpleasant",
	"brutish"}

func GenerateName() string {
	return fmt.Sprintf("%s_%s", adjectives[rand.Intn(len(adjectives))], nouns[rand.Intn(len(nouns))])
}
