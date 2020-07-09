# Blacklist Selector

The blacklist selector is a go-micro/selector which filters nodes based on which have errored out. 
It operates much like a circuit breaker. If a node returns an error 3 consecutive times it will 
be blacklisted. After a period of 30 seconds it will be put back into the list of nodes.
