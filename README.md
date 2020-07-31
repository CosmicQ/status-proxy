# status-proxy
A status "proxy" for load balancers that require a HTTP 200 response, but the application has other indicators that don't reply with http 200 response.

I ran into a situation where I wasn't able to use non HTTP 200 as a status for the service.  I needed to have a local check that if I received a 405 response, then reply to the load balancer with a 200 response.

The goal of this project is to create a service that can run just about any local check (Nagios check style) and return a response to a load balancer.  For instance, If a URL check AND database query return, then reply with a 200 status code to the load balancer.  

# How this works
This is a simple service written in Go that receives an HTTP request.  The service will execute a file named in the path located in /usr/local/status.  For example:

`http://mysystem.company.tld:8181/foo` will execute a script called `foo` if it is located in /usr/local/status.

These scripts simply use the same method as Nagios checks, so anything that exits with an exit code of `0` will result in the status returning a 200 http response to whatever is checking it.  Any other exit code will respond with a 503 error code.