
# datamusecli
a command-line based word-finding query tool that match a given set of constraints (provided through different flags)

This tool uses datamuse API(https://www.datamuse.com/api/) to query as command line. 

Prerequisite: <br />
1.)Go

Setup: <br />
1.) RUN ***go get github.com/aniketsupertramp/datamusecli/src*** <br />
2.) ***cd $GOPATH*** <br />
3.) ***go build -o datamusee datamusecli/src/main.go*** <br />
4.) ***cp datamuse /usr/local/bin*** <br />

***Done !!!***

Run ***datamuse --help*** to find different flags you can use. 


e.g. :  ***datamuse --similar="ringing in the ears"*** will give words with a meaning similar to **ringing in the ears**

Also, with flag --max you can specify max number of results you want (default is 10)

e.g. :  ***datamuse --similar="ringing in the ears" --max=20***




Feel free to suggest any changes/corrections.
