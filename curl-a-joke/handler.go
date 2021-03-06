package function

import (
	"math/rand"
	"time"
)

var jokes = [...]string{
	"What goes up and down but does not move? Stairs!",
	"Where should a 500 pound alien go? On a diet!",
	"What did one toilet say to the other? You look a bit flushed!",
	"Why did the picture go to jail? Because it was framed!",
	"What did one wall say to the other wall? I'll meet you at the corner!",
	"What did the paper say to the pencil? Write on!",
	"What gets wetter the more it dries? A towel!",
	"Why do bicycles fall over? Because they are two-tired!",
	"What did Cinderella say when her photos did not show up? Someday my prints will come!",
	"Why was the broom late? It over swept!",
	"What part of the car is the laziest? The wheels, because they are always tired!",
	"What's the difference between a guitar and a fish? You can't tuna fish!",
	"What do you call a pile of kittens? A meowtain!",
	"Did you hear about the hungry clock? It went back four seconds!",
	"What do you get from a pampered cow? Spoiled milk!",
	"Did you hear about that new broom? It's sweeping the nation!",
	"What do lawyers wear to court? Lawsuits!",
	"What is brown and has a head and a tail but no legs? A penny!",
	"What do you get when you cross a snowman with a vampire? Frostbite!",
	"What do you call a bee that lives in America? USB!"}

func Handle(req []byte) string {
	rand.Seed(time.Now().UTC().UnixNano())
	return jokes[rand.Intn(len(jokes))]
}
