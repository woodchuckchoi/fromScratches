import asyncio
import aioredis

import json

from sanic import Sanic

from chat.redis_pub_sub import pub, sub
from utils import importConfig

app = Sanic(__name__)
_, _, _, _, REDIS_HOST, REDIS_PORT = importConfig()

@app.listener('before_server_start')
async def init_redis(app, loop):
    app.pub = await aioredis.create_redis_pool((REDIS_HOST, REDIS_PORT))
    app.sub = await aioredis.create_redis_pool((REDIS_HOST, REDIS_PORT))
    app.lobby   = await aioredis.create_redis_pool((REDIS_HOST, REDIS_PORT))

    app.connections = dict()
    app.pendings = []

@app.listener('after_server_stop')
async def init_redis(app, loop):
    app.pub.close()
    app.sub.close()
    app.lobby.close()
    
    for pending in app.pendings:
        pending.cancel()

@app.websocket('/lobby')
async def lobby(request, ws):
    while True:
        try:
            rooms = await app.lobby.pubsub_channels('channel:*')
            rooms = list(map(lambda x: x.decode('UTF-8'), rooms))
            await ws.send(json.dumps({'rooms':rooms})) # this part probably needs editing
            await asyncio.sleep(30)
        except Exception as e:
            print(e)
            # debugging
    

@app.websocket('/chat/<roomName>')
async def chat(request, ws, roomName):
    ws.userName = 'Roman Gherkins'

    if roomName not in app.connections:
        app.connections[roomName] = []
    app.connections[roomName].append(ws)

    pub_task = asyncio.create_task(pub(app.pub, roomName, ws, app))
    tasks = [pub_task]

    if roomName not in app.subs:
        sub_task = asyncio.create_task(sub(app.sub, roomName, app))
        tasks.append(sub_task)

    d, p = await asyncio.wait(tasks)
    if p:
        app.pendings += p

    # when one dies, the other should discontinue serving
    # done, pending = await asyncio.wait([pub_task, sub_task], return_when=asyncio.FIRST_COMPLETED)
    # for pend in pending:
    #     pend.cancel()

app.run('0.0.0.0', 8000)