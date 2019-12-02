# status-proxy
A status "proxy" for load balancers that require a HTTP 200 response, but the application has other indicators that don't reply with http 200 response.

I ran into a situation where I wasn't able to use non HTTP 200 as a status for the service.  I needed to have a local check that if I received a 405 response, then reply to the load balancer with a 200 response.

The goal of this project is to create a service that can run just about any local check (Nagios check style) and return a response to a load balancer.  For instance, If a URL check AND database query return, then reply with a 200 status code to the load balancer.  
