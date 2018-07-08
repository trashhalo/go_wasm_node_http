# go_wasm_node_http
This is a minimal example of how to control node with golang 1.11 webassembly. It creates an http server and listens to requests on port 3000

# Why
Mostly because I can. I work professionally in both node and golang and was curious about golang webassembly and how easily it would connect into a node codebase.

It's actually very easy. The only thing I got caught up by is the documentation is lacking and webassembly modules exit immediately without extra sauce. (see wait function)

# Build
1. clone the repo
2. build the docker container `docker build -t go-node`

# Running
1. run the server! `docker run -t -p 3000:3000 go-node`
2. Open your browser to http://localhost:3000