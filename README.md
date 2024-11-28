# Link Shortener

a basic link shortener written in go, using postgres as a backend

# How to use

make a shortened link

```sh
curl --request POST \
  --url 'https://shorten.codaea.com/api/new?url=example.com'
```
this returns a `slug`. if you visit `shorten.codaea.com/<YOURSLUG>` it will redirect you.

# Performance

## Just postgres, no caching. 

with just postgres and no caching, on a virtual machine running with 

4vCpu
4gb ram

it would top out on the current jmeter test at 613ms to create a link, and viewing at 282ms. with 1200 new links a second and 1370 new links a second

this was mostly database usage though, and if we use in memory go and redis it would go faster.
