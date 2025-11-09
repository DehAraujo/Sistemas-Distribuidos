const zmq = require("zeromq");

// Função para gerar UUID manualmente
function uuidv4() {
    return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function (c) {
        const r = Math.random() * 16 | 0;
        const v = c === 'x' ? r : (r & 0x3 | 0x8);
        return v.toString(16);
    });
}

async function main() {
    const sock = new zmq.Push();
    await sock.connect("tcp://server:5556");

    console.log("Bot conectado ao servidor...");

    setInterval(async () => {
        let msg = {
            id: uuidv4(),
            timestamp: Date.now(),
            type: "log",
            message: "Olá do bot!"
        };
        await sock.send(JSON.stringify(msg));
        console.log("Mensagem enviada:", msg);
    }, 3000);
}

main();
