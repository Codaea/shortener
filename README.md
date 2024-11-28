# Link Shortener

That is incredibly overdesigned.

core api written in go, 

using postgresql, and redis as a caching layer.


built to handle up to __ calls a second

# Performance

## Just postgres, no caching. 

with just postgres and no caching, on a virtual machine running with 

4vCpu
4gb ram

it would top out on the current jmeter test at 613ms to create a link, and viewing at 282ms. with 1200 new links a second and 1370 new links a second

