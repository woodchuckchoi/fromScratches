import asyncio
import aioredis
import json

from sanic import Sanic

from app.redis_pub_sub import pub, sub

app = Sanic(__name__)

@app.listener('before_server_start')
async def init_redis(app, loop):
    app.pub = await aioredis.create_redis_pool(('0.0.0.0', 6379))
    app.sub = await aioredis.create_redis_pool(('0.0.0.0', 6379))

    app.connections = set()

@app.listener('after_server_stop')
async def init_redis(app, loop):
    app.pub.close()
    app.sub.close()

@app.websocket('/')
async def chat(request, ws):
    app.connections.add(ws)
    chan = 'test:1'
    pub_task = asyncio.create_task(pub(app.pub, chan, ws, app))
    sub_task = asyncio.create_task(sub(app.sub, chan, ws, app))
    done, pending = await asyncio.wait([pub_task, sub_task], return_when=asyncio.FIRST_COMPLETED)
    for pend in pending:
        pend.cancel()

app.run('0.0.0.0', 8000)