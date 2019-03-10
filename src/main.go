package main

import (
	"flag"
	"os"
	"fmt"
	"datamusecli/src/process"
)

func main() {
	flagReq := process.FlagRequest{}
	flagReq.Similar = flag.String("similar", "", "Similar meaning words")
	flagReq.Sound = flag.String("sound", "", "Similar sound words")
	flagReq.LeftOf = flag.String("left", "", "words immediate left of given word")
	flagReq.RighOf = flag.String("right", "", "words immediate rigt of given word")
	flagReq.Spelled = flag.String("spell", "", "Similar Spelled words")
	flagReq.Rhyme = flag.String("rhyme", "", "Rhymed words")
	flagReq.Adjective = flag.String("adj", "", "Adjectives")
	flagReq.Topic = flag.String("topic", "", "Words with given topic")
	flagReq.Noun = flag.String("noun", "", "Nouns for a word")
	flagReq.Synonym = flag.String("syn", "", "Synonyms for a word")
	flagReq.Antonym = flag.String("ant", "", "Antonyms for a word")
	flagReq.PartOf = flag.String("par", "", "Part of (meronyms)")
	flagReq.Homophones = flag.String("hom", "", "sound-alike words (Homophones)")
	flagReq.ConsonantMatch = flag.String("cns", "", "Consonant match")
	max := flag.String("max", "10", "Max number of Results")

	flag.Parse()

	if flag.NFlag() == 0 { // if no flags supplied
		flag.PrintDefaults()
		os.Exit(1)
	}
	fmt.Println(process.Process(flagReq, max))
}
