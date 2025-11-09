# server-python/server.py
import zmq
import json
import os
import time
from pathlib import Path

DATA_DIR = Path("data")
USERS_FILE = DATA_DIR / "users.json"
CHANNELS_FILE = DATA_DIR / "channels.json"
LOGINS_FILE = DATA_DIR / "logins.jsonl"

def ensure_data():
    DATA_DIR.mkdir(exist_ok=True)
    if not USERS_FILE.exists():
        USERS_FILE.write_text(json.dumps([]))
    if not CHANNELS_FILE.exists():
        CHANNELS_FILE.write_text(json.dumps([]))
    if not LOGINS_FILE.exists():
        LOGINS_FILE.write_text("")

def load_users():
    return json.loads(USERS_FILE.read_text())

def save_users(users):
    USERS_FILE.write_text(json.dumps(users, indent=2))

def load_channels():
    return json.loads(CHANNELS_FILE.read_text())

def save_channels(channels):
    CHANNELS_FILE.write_text(json.dumps(channels, indent=2))

def append_login(entry):
    with LOGINS_FILE.open("a") as f:
        f.write(json.dumps(entry) + "\n")

def make_reply(service, data):
    return json.dumps({"service": service, "data": data})

def handle_login(data):
    user = data.get("user")
    ts = data.get("timestamp", int(time.time()))
    if not user:
        return {"status":"erro", "timestamp": int(time.time()), "description":"campo 'user' obrigatório"}
    users = load_users()
    if user not in users:
        users.append(user)
        save_users(users)
    # persist login
    append_login({"user": user, "timestamp": ts, "server_ts": int(time.time())})
    return {"status":"sucesso", "timestamp": int(time.time())}

def handle_users(data):
    users = load_users()
    return {"timestamp": int(time.time()), "users": users}

def handle_channel(data):
    channel = data.get("channel")
    ts = data.get("timestamp", int(time.time()))
    if not channel:
        return {"status":"erro", "timestamp": int(time.time()), "description":"campo 'channel' obrigatório"}
    channels = load_channels()
    if channel in channels:
        return {"status":"erro", "timestamp": int(time.time()), "description":"canal já existe"}
    channels.append(channel)
    save_channels(channels)
    return {"status":"sucesso", "timestamp": int(time.time())}

def handle_channels(data):
    channels = load_channels()
    return {"timestamp": int(time.time()), "channels": channels}

def main():
    ensure_data()
    context = zmq.Context()
    socket = context.socket(zmq.REP)
    bind_addr = "tcp://0.0.0.0:5555"
    socket.bind(bind_addr)
    print("Server REP listening on", bind_addr)
    while True:
        try:
            raw = socket.recv()
            # incoming assumed JSON at this stage
            try:
                msg = json.loads(raw.decode())
            except Exception as e:
                reply = make_reply("error", {"status":"erro","timestamp":int(time.time()), "description":"invalid json"})
                socket.send_string(reply)
                continue

            service = msg.get("service")
            data = msg.get("data", {})

            if service == "login":
                res = handle_login(data)
                socket.send_string(make_reply("login", res))
            elif service == "users":
                res = handle_users(data)
                socket.send_string(make_reply("users", res))
            elif service == "channel":
                res = handle_channel(data)
                socket.send_string(make_reply("channel", res))
            elif service == "channels":
                res = handle_channels(data)
                socket.send_string(make_reply("channels", res))
            else:
                socket.send_string(make_reply("error", {"status":"erro","timestamp":int(time.time()), "description":"serviço desconhecido"}))
        except KeyboardInterrupt:
            break
        except Exception as e:
            print("Erro no servidor:", e)

if __name__ == "__main__":
    main()
