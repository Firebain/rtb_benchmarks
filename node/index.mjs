import Fastify from "fastify";

const fastify = Fastify();

const timeout = (ms) => new Promise((res) => setTimeout(res, ms));

const arrays = [];

const handler = async (request, reply) => {
  let i = 0;

  for (let b = 0; b < 50000; b++) {
    const array = arrays[b % arrays.length];

    if (array.indexOf("qw2") !== -1) {
      i++;
    }
  }

  await timeout(10);

  for (let b = 0; b < 50000; b++) {
    const array = arrays[b % arrays.length];

    if (array.indexOf("qw5") !== -1) {
      i++;
    }
  }

  await timeout(10);

  for (let b = 0; b < 50000; b++) {
    const array = arrays[b % arrays.length];

    if (array.indexOf("qw8") !== -1) {
      i++;
    }
  }

  return { result: i };
};

fastify.get("/", handler);

for (let i = 0; i < 20; i++) {
  if (i % 2 === 0) {
    arrays.push(["qw1", "qw2", "qw3", "qw4", "qw5"]);
  } else {
    arrays.push(["qw5", "qw4", "qw3", "qw2", "qw1"]);
  }
}

await fastify.listen({ port: 8080, host: "0.0.0.0" });
