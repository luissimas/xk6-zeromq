import zeromq from "k6/x/zeromq";

const socket = zeromq.newSocket("tcp://127.0.0.1:6969");

export const options = {
  duration: "5s",
  vus: 10,
};

export default function () {
  zeromq.send(socket, "foo");
}

export function tearDown() {
  zeromq.close(socket);
}
