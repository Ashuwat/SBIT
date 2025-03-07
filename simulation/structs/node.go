package main

type Node struct {
	// data (n x 1) matrix
    // credibility (n x 1) matrix
    // location (3 x 1) matrix - computed when looked at a frame
    // velocity (3 x 1) matrix
    // trust {node_address: weight}
    // status - is deterministic of trust (eq) - further work on 
    // address - some 32 bit number - deterministic not random
    // release for information - latest piece - represented with time
    // release factor - a number/threshold the release has to reach
}
