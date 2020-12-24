import http from "k6/http";
import { Counter } from "k6/metrics";
import { check } from "k6";

let headerFile = open("../example/header.html", "b");
let pageFile = open("../example/page.html", "b");

let failCounter = new Counter("failed requests");

export let options = {
  stages: [{ duration: "1m", target: __ENV.MAX_VUS }],
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
    headerTemplate: http.file(headerFile, "header.html"),
    html: http.file(pageFile, "page.html"),
    format: "A4",
  };
  let res = http.post(__ENV.BASE_URL + "/convert", data);
  check(res, {
    "is status 200": (r) => r.status === 200,
    "is not status 504": (r) => r.status !== 504,
    "is not status 500": (r) => r.status !== 500,
    "is not status 429": (r) => r.status !== 429,
  });
  if (res.status !== 200) {
    failCounter.add(1);
  }
}
