import aioredis
import json

class RedisConnection:
    async def __init__(self, host, port):
        self.__pool = await aioredis.create_redis_pool(host, port) 
        self.__connections = {}

    async def get_redis(self):
        if not self.__pool:
            await self.connect()
        return self.__pool

    async def close(self):
        self.__pool.close()
        await self.__pool.wait_closed()

async def recvChat(ws, app, origin):
    while True:
        try:
            while await origin.wait_message():
                msg = await origin.get()
                await ws.send(msg)
        
        except Exception as e:
            await ws.send(json.dumps({'message': 'Connection unstable...'}))
        
        finally:
            app.redis.unsubscribe('channel:{}'.format(origin))

async def sendChat(ws, app, dest):
    while True:
        try:
            recv = json.loads(await ws.recv())
            await app.redis.publish('channel:{}'.format(dest), recv['message'])
            
        except Exception as e:
            pass

