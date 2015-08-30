go-slackbox-example
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
