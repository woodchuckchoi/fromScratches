import asyncio
import json

async def pub(pool, channel, ws, app):
    while True:
        try:
            msg = await ws.recv()
            await pool.publish(channel, msg)
        except Exception as e:
            await pool.publish(channel, 'Error : {}'.format(e))
            # should find a way to gracefully exit
            # app.connections.remove(ws) # when disconnected, websocket object is designed to remove itself automatically

async def sub(pool, channel, ws, app):
    sub = (await pool.subscribe(channel))[0]
    while True:
        try:
            while await sub.wait_message():
                var = await sub.get()
                for conn in app.connections:
                    await conn.send(json.dumps({'msg': var.decode('utf-8')}))
        except:
            # for conn in app.connections:
            #     await conn.send(json.dumps({'msg': 'Unexpected connection error occured...'}))
            # await asyncio.sleep(5)
            # again, all the ws objects in app.connections delete themselves in case of Exception
            pass
        finally:
            await pool.unsubscribe(channel)
