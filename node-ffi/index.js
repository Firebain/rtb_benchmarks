const fastify = require("fastify")();
const ffi = require("./ffi");

const handler = async (request, reply) => {
  let i = await ffi.exists();

  return { result: i };
};

fastify.get("/", handler);

fastify.listen({ port: 8080, host: "0.0.0.0" });
