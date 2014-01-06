goriot
======

Go bindings for official Riot/League of legends REST API

Add any requests to issues.

For help using the libary look at the goriot_test.go file

TODO: Improve Readme, improve error handling, improve date handling, rate handling, improved testing (all will be filed as issues)

To install just use 
```
go get github.com/TrevorSStone/goriot
```

To run tests successfully change the apikey value at the top to your key. Key's can be obtained through https://developer.riotgames.com/

Tests will take upwards of 130s if ran together due to the default rate limiting settings. This is because it has to block for 10s after each 10 requests. If you have different rate limits, change the tests to have them run faster.

The way rate limiting works is semi-experimental. I believe it works under all conditions, but if you get a 429 error please report it.

I also have a library I wrote prior that went through the RTMP(s) protocol, but I haven't open sourced it because it is essentially obsolete now. If anyone is interested in that source as well, let me know. If there is enough interest I can put what I currently have up.

http://godoc.org/github.com/TrevorSStone/goriot
