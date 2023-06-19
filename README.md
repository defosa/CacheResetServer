# CacheResetServer
CacheResetServer is a simple Go application that provides an HTTP server to reset the cache of an Nginx server. 
A simple Go application that provides an HTTP endpoint to reset the cache of an Nginx server. This application listens on a specified port and allows cache reset requests only from a specific IP range.



```bash
$ curl -X GET http://localhost:8999/reset
