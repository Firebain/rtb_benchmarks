module.exports = [
  {
    name: "service",
    script: "index.mjs",
    instances: "max",
    exec_mode: "cluster",
    autorestart: true,
    kill_timeout: 5000,
    watch: false,
  },
];
