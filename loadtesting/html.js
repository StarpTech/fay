import http from "k6/http";
import { Counter } from "k6/metrics";
import { check } from "k6";

let headerFile = open("../example/header.html", "b");

let failCounter = new Counter("failed requests");

export let options = {
  stages: [{ duration: "5m", target: __ENV.MAX_VUS }],
  thresholds: {
    "failed requests": [
      {
        threshold: "count<1",
        abortOnFail: true,
      },
    ],
  },
};

export default function () {
  let data = {
    "header.html": http.file(headerFile, "header.html"),
    format: "A4",
    url: "http://localhost:3001/page.html"
  };
  let res = http.post(__ENV.BASE_URL + "/convert", data);
  check(res, {
    "is status 200": (r) => r.status === 200,
    "is not status 504": (r) => r.status !== 504,
    "is not status 500": (r) => r.status !== 500,
  });
  if (res.status !== 200) {
    failCounter.add(1);
  }
}
