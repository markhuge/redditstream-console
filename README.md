I wrote this thing really fast to get UFC 203 updates in my terminal. It's really bad and I'll fix it after the fights.

## Abstract

The just spits out a stream of the latest comments on a reddit post to the console (basically a console version of redditstream.com)

If you are reading this, you probably shouldn't use this for anything. I've had to write a lot of ruby @ work for the last 10 months and forgot how to golang.
I'm writing all my random bs in go to get back into practice. This could be curl + a few lines of bash with the same result. 

## TODO

### In process caching

This isn't for performance, just de-duping records.

- Check ID. If in cache: next, else: write to cache and pass to print.
