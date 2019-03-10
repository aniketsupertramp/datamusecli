package process

import (
	"strconv"
	"time"
	"github.com/franela/goreq"
	"strings"
	"fmt"
	"os"
)

const (
	// DatamuseAPIURL datamuse word api url
	DatamuseAPIURL = "https://api.datamuse.com/words"
)

var (
	// FLAGS list of allowed flags
	FLAGS = []string{"similar", "sound", "left", "right", "spell", "rhyme", "adj", "topic", "noun", "syn", "ant", "par", "hom", "cns"}
)

// FlagRequest struct
type FlagRequest struct {
	Similar        *string
	Sound          *string
	LeftOf         *string
	RighOf         *string
	Spelled        *string
	Rhyme          *string
	Adjective      *string
	Topic          *string
	Noun           *string
	Synonym        *string
	Antonym        *string
	PartOf         *string
	Homophones     *string
	ConsonantMatch *string
}

// Response struct
type Response struct {
	Word         string   `json:"word"`
	Score        int      `json:"score"`
	NumSyllables int      `json:"numSyllables"`
	Tags         []string `json:"tags"`
}

// Process processes the command and args passed
func Process(flagReq FlagRequest, max *string) string {
	maxWords, parsErr := strconv.Atoi(*max)
	if parsErr != nil {
		fmt.Println(*max)
		return "ERROR: Invalid value for flag max"
	}
	query := buildQuery(flagReq, maxWords)
	response, err := fetchResponse(query)
	if err != nil {
		return err.Error()
	}
	words, parseErr := parseData(response)
	if parseErr != nil {
		return parseErr.Error()
	}
	return "[ " + strings.Join(words, ",") + " ]"
}

func buildQuery(flagReq FlagRequest, maxWords int) string {
	var query string
	for _, flag := range FLAGS {
		switch strings.ToLower(flag) {
		case "similar":
			buildQueryUtil(&query, "ml", *flagReq.Similar)
		case "sound":
			buildQueryUtil(&query, "sl", *flagReq.Sound)
		case "spell":
			buildQueryUtil(&query, "sp", *flagReq.Spelled)
		case "adj":
			buildQueryUtil(&query, "rel_jjb", *flagReq.Adjective)
		case "rhyme":
			buildQueryUtil(&query, "rel_rhy", *flagReq.Rhyme)
		case "syn":
			buildQueryUtil(&query, "rel_syn", *flagReq.Synonym)
		case "par":
			buildQueryUtil(&query, "rel_par", *flagReq.PartOf)
		case "hom":
			buildQueryUtil(&query, "rel_hom", *flagReq.Homophones)
		case "ant":
			buildQueryUtil(&query, "rel_ant", *flagReq.Antonym)
		case "cns":
			buildQueryUtil(&query, "rel_cns", *flagReq.ConsonantMatch)
		case "topic":
			buildQueryUtil(&query, "topics", *flagReq.Topic)
		case "noun":
			buildQueryUtil(&query, "rel_jja", *flagReq.Noun)
		case "left":
			buildQueryUtil(&query, "lc", *flagReq.LeftOf)
		case "right":
			buildQueryUtil(&query, "rc", *flagReq.RighOf)
		default:
			fmt.Println("ERROR: INAVLID flag encountered")
			os.Exit(1)

		}
	}
	return query + "max=" + strconv.Itoa(maxWords)
}

func fetchResponse(query string) (*goreq.Response, error) {
	url := DatamuseAPIURL + "?" + query
	resp, err := goreq.Request{
		Method:      "GET",
		Uri:         url,
		ContentType: "application/json",
		Accept:      "application/json",
		Timeout:     20 * time.Second,
	}.Do()
	return resp, err
}

func buildQueryUtil(query *string, param, val string) {
	if val != "" {
		*query += param + "=" + formatStringInput(val) + "&&"
	}
}

func parseData(resp *goreq.Response) ([]string, error) {
	var responseDataList []Response
	err := resp.Body.FromJsonTo(&responseDataList)
	if err != nil {
		return nil, err
	}
	words := getStringDataFromResponse(responseDataList)
	return words, nil
}

func getStringDataFromResponse(responseList []Response) ([]string) {
	var words []string
	for _, response := range responseList {
		words = append(words, response.Word)
	}
	return words
}

func formatStringInput(input string) string {
	return strings.Join(strings.Split(input, " "), "+")
}
