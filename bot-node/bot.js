// bot-node/bot.js
const zmq = require("zeromq");
const sock = new zmq.Request();
const { v4: uuidv4 } = require("uuid");

async function run() {
  await sock.connect("tcp://server:5555");
  const name = "bot-" + Math.floor(Math.random()*10000);
  const msg = {
    service: "login",
    data: {
      user: name,
      timestamp: Math.floor(Date.now()/1000)
    }
  };
  await sock.send(JSON.stringify(msg));
  const [reply] = await sock.receive();
  console.log("bot login reply:", name, reply.toString());
  // bot stays up; for part1 it does nothing else
  setInterval(()=>{}, 1000);
}

run().catch(err => console.error(err));
