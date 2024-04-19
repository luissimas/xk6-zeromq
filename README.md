# xk6-zeromq

A [k6](https://k6.io/) extension for the [ZeroMQ](https://zeromq.org/) networking library.

## Build

To build a `k6` binary with this extension, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- Git

Then:

1. Install `xk6`:

```shell
go install github.com/k6io/xk6/cmd/xk6@latest
```

2. Build the binary:

``` sh
xk6 build --with github.com/luissimas/xk6-zeromq@latest
```

## Example

Given a `example.js` file with the following contents:

``` javascript
import zeromq from "k6/x/zeromq";

const socket = zeromq.newSocket("tcp://127.0.0.1:6969");

export const options = {
  duration: "5s",
  vus: 10,
};

export default function () {
  const resp = zeromq.send(socket, "foo");
}

export function tearDown() {
  zeromq.close(socket);
}
```

The following output is produced:

```
$ ./k6 run example.js

          /\      |‾‾| /‾‾/   /‾‾/
     /\  /  \     |  |/  /   /  /
    /  \/    \    |     (   /   ‾‾\
   /          \   |  |\  \ |  (‾)  |
  / __________ \  |__| \__\ \_____/ .io

     execution: local
        script: example.js
        output: -

     scenarios: (100.00%) 1 scenario, 10 max VUs, 35s max duration (incl. graceful stop):
              * default: 10 looping VUs for 5s (gracefulStop: 30s)


     data_received.........: 0 B   0 B/s
     data_sent.............: 0 B   0 B/s
     iteration_duration....: avg=574.88µs min=148.18µs med=538µs    max=4.48ms p(90)=675.49µs p(95)=739.47µs
     iterations............: 85946 17187.806038/s
     vus...................: 10    min=10         max=10
     vus_max...............: 10    min=10         max=10
     zeromq_req_count......: 85946 17187.806038/s
     zeromq_req_duration...: avg=533.98µs min=5.06µs   med=497.57µs max=4.42ms p(90)=634.62µs p(95)=694.12µs
     zeromq_req_failed.....: 0.00% ✓ 0            ✗ 85946


running (05.0s), 00/10 VUs, 85946 complete and 0 interrupted iterations
default ✓ [======================================] 10 VUs  5s
```

## Development

For local development, use the `Makefile` provided in the repository. The default target will setup `xk6`, format the code and build the extension from the local repository. 

```sh
make
```

The binary will be available at `bin/k6`. You can then run tests using the locally built extension:

```sh
./bin/k6 run example.js
```

## Acknowledgments

This project was only possible with the contributions from these projects:

- https://github.com/go-zeromq/zmq4
- https://github.com/dgzlopes/xk6-zmq
- https://github.com/NAlexandrov/xk6-tcp
- https://github.com/MATRIXXSoftware/xk6-diameter
