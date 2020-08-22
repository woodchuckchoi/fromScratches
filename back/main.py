import asyncio
import aioredis

import json

from sanic import Sanic

from chat.redis_pub_sub import pub, sub

app = Sanic(__name__)

@app.listener('before_server_start')
async def init_redis(app, loop):
    app.pub = await aioredis.create_redis_pool(('0.0.0.0', 6379))
    app.sub = await aioredis.create_redis_pool(('0.0.0.0', 6379))

    app.lobby   = await aioredis.create_redis_pool(('0.0.0.0', 6379))

    app.connections = set()

@app.listener('after_server_stop')
async def init_redis(app, loop):
    app.pub.close()
    app.sub.close()
    app.lobby.close()

@app.websocket('/lobby')
async def lobby(request, ws):
    while True:
        try:
            rooms = await app.lobby.pubsub_channels('channel:*')
            rooms = list(map(lambda x: x.decode('UTF-8'), rooms))
            await ws.send(json.dumps({'rooms':rooms})) # this part probably needs editing
            await asyncio.sleep(5)
        except Exception as e:
            print(e)
            # debugging
    

@app.websocket('/chat/<roomName>')
async def chat(request, ws, roomName):
    app.connections.add(ws)
    roomName = 'channel:'.format(roomName)
    pub_task = asyncio.create_task(pub(app.pub, roomName, ws, app))
    sub_task = asyncio.create_task(sub(app.sub, roomName, ws, app))

    # when one dies, the other should discontinue serving
    done, pending = await asyncio.wait([pub_task, sub_task], return_when=asyncio.FIRST_COMPLETED)
    for pend in pending:
        pend.cancel()

app.run('0.0.0.0', 8000)
