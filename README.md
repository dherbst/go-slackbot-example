go-slackbot-example
===================

An Example of a slack integration.

Testing
=======

Use the appengine testing framework

    goapp test

Coverage
--------

To get coverage with goapp you have to convert the path to the files.

    goapp test -coverprofile=coverage.out
    sed -i -e "s#.*/\(.*\.go\)#\./\\1#" coverage.out
    goapp tool cover -html coverage.out

Slides
------

Can be viewed live here:

    http://go-talks.appspot.com/github.com/dherbst/go-slackbot-example/present/go-slackbot-example.slide

Or install go present

    go get golang.org/x/tools/cmd/present
    present go-slackbot-example.slide
