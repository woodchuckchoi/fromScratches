import asyncio
import json

async def pub(pool, roomName, ws, app):
    channel = 'channel:{}'.format(roomName)
    while True:
        try:
            msg = await ws.recv()
            await pool.publish(channel, f'{ws.userName};{msg}')
        except Exception as e:
            await pool.publish(channel, 'Error;{}'.format(e))
            await pool.unsubscribe(channel)
            if roomName in app.connections:
                if ws in app.connections[roomName]:
                    app.connections[roomName].remove(ws)
            return
            
# should find a way to gracefully exit
# app.connections.remove(ws) # when disconnected, websocket object is designed to remove itself automatically

async def sub(pool, roomName, app):
    channel = 'channel:{}'.format(roomName)
    sub = (await pool.subscribe(channel))[0]
    while True:
        try:
            while await sub.wait_message():
                var = (await sub.get()).decode('UTF-8')
                sender, msg = var[:var.find(';')], var[var.find(';')+1:]
                for conn in app.connections[roomName]:
                    try:
                        await conn.send(json.dumps({'sender': sender, 'msg': msg}))
                    except:
                        if roomName in app.connections:
                            if conn in app.connections[roomName]:
                                app.connections[roomName].remove(conn)
                        else:
                            return

                        if not app.connections[roomName]:
                            app.connections.pop(roomName)
                            return
        except:
            print('Redis channel error')
            # for conn in app.connections:
            #     await conn.send(json.dumps({'msg': 'Unexpected connection error occured...'}))
            # await asyncio.sleep(5)
            # again, all the ws objects in app.connections delete themselves in case of Exception
            pass
        # finally:
        #     await pool.unsubscribe(channel)
