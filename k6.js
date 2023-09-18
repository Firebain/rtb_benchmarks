import http from "k6/http";

export const options = {
  stages: [
    { duration: "10s", target: 100 },
    { duration: "1m", target: 100 },
    { duration: "10s", target: 0 },
  ],
};

export default () => {
  http.get("http://localhost:8080");
};
