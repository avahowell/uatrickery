# uatrickery

uatrickery is a package that implements a HTTP server that serves valid image data for some user agents, and arbitrary HTML/JS for others. This can be used as an attack on some social networks that have image thumbnailing bots with predictable user agents.

# usage

uatrickery takes a few critical flags. `-targets` should be a path to a newline separtaed file which defines user agents to serve the valid image for. `-image` should be the path to the image to serve. `-payload` should be the path to the HTML/JS/CSS to serve non-matching useragents. You can change the bind address of the server using `-bind`.

# LICENSE

The MIT License (MIT)
