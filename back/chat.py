# PIP
import aioredis
from sanic import Sanic
from sanic.response import json
from sanic_session import Session, AIORedisSessionInterface

# Custom
from samaria import Samaria
from utils import importConfig, timeNow

if __name__ == "__main__":
    HOST, USER, PASS, DATABASE, REDIS_HOST, REDIS_PORT = importConfig()

    app             = Sanic(__name__)
    samaria         = Samaria(host=HOST, user=USER, password=PASS, database=DATABASE)
    redisSession    = Session()

    @app.listener('before_server_start')
    async def redis_init(app, loop):
        app.redis   = await aioredis.create_redis_pool((REDIS_HOST, REDIS_PORT))
        redisSession.init_app(app, interface=AIORedisSessionInterface(app.redis, expiry=43200))

    @app.websocket('/')
    async def lobby(request, ws):
        rooms = await app.redis.get('rooms')
        return json({'rooms': rooms})

    @app.websocket('/room/subscribe/<room>')
    async def subscribe(request, ws):
        raise NotImplementedError