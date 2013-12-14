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

Tests will fail if all ran at once due to the default rate limiting. Please run one at a time before filing an issue

I also have a library I wrote prior that went through the RTMP(s) protocol, but I haven't open sourced it because it is essentially obsolete now. If anyone is interested in that source as well, let me know. If there is enough interest I can put what I currently have up.
