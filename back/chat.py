# PIP
import aioredis
import asyncio
from sanic import Sanic
from sanic.response import json
from sanic_session import Session, AIORedisSessionInterface

# Custom
from samaria import Samaria
from utils import importConfig, timeNow
from chatUtils import recvChat, sendChat

if __name__ == "__main__":
    HOST, USER, PASS, DATABASE, REDIS_HOST, REDIS_PORT = importConfig()

    app             = Sanic(__name__)
    samaria         = Samaria(host=HOST, user=USER, password=PASS, database=DATABASE)
    redisSession    = Session()

    @app.listener('before_server_start')
    async def redis_init(app, loop):
        app.redis   = await aioredis.create_redis_pool((REDIS_HOST, REDIS_PORT))
        redisSession.init_app(app, interface=AIORedisSessionInterface(app.redis, expiry=43200))

    @app.listener('after_server_stop')
    async def redis_end(app, loop):
        await app.redis.close()

    @app.websocket('/')
    async def lobby(request, ws):
        #rooms get refreshed when front wants them to
        rooms = await app.redis.get('rooms')
        return json({'rooms': rooms})

    @app.route('/room/create/<roomNo>')
    async def createRoom(request, roomNo):
        raise NotImplementedError

    @app.websocket('/room/subscribe/<roomNo>')
    async def subscribe(request, ws, roomNo):
        room = await app.redis.subscribe('room:{}'.format(roomNo))
        
        send = asyncio.create_task(sendChat(ws, app, room))
        recv = asyncio.create_task(recvChat(ws, app, room))

        done, pending = await asyncio.wait([send, recv], return_when=asyncio.FIRST_COMPLETED)
        
        for pend in pending:
            pend.cancel()

    app.run('0.0.0.0', 8080)