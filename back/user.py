# Built-in
import os
import time

# PIP
import aioredis
from sanic import Sanic
from sanic.response import json
from sanic_session import Session, AIORedisSessionInterface

# Custom
from samaria import Samaria
from utils import importConfig

if __name__ == "__main__":
    HOST, USER, PASS, DATABASE, REDIS_HOST, REDIS_PORT = importConfig()

    app             = Sanic(__name__)
    samaria         = Samaria(host=HOST, user=USER, password=PASS, database=DATABASE)
    redisSession    = Session()

    @app.listener('before_server_start')
    async def redis_init(app, loop):
        app.redis   = await aioredis.create_redis_pool((REDIS_HOST, REDIS_PORT))
        redisSession.init_app(app, interface=AIORedisSessionInterface(app.redis, expiry=43200))

    def check_auth(request):
        if request.ctx.session.get('devQueueId') and request.ctx.session.get('devQueueLastLogin'):
            return True
        return False

    @app.route('/')
    async def all_user(request):
        await samaria.execute('SELECT DISTINCT User.id, (SELECT COUNT(*) FROM Record WHERE Record.user_id = User.idx) AS game_count FROM User ORDER BY game_count DESC, User.id ASC;')
        ret = await samaria.fetchall()
        return json(ret)

    @app.route('/login', methods=['POST'])
    async def login(request):
        if check_auth(request):
            return json({"result": True})

        loginId = request.json['loginId']
        loginPw = request.json['loginPw']

        await samaria.execute("SELECT IF(SHA2(CONCAT(creation_date, '{}'), 512) = password, 1, 0) FROM User WHERE id = '{}'".format(loginPw, loginId))
        
        ret = await samaria.fetchall()

        # Login log collection logic to be added here

        if ret[0][0]:
            request.ctx.session['devQueueId'] = loginId
            request.ctx.session['devQueueLastLogin'] = time.strftime("%Y/%m/%d %H:%M:%S")
            return json({"result": True})
        return json({"result": False})

    @app.route('/logout', methods=['GET'])
    async def logout(request):
        request.ctx.session['devQueueId'] = None
        request.ctx.session['devQueueLastLogin'] = None
        return json({'result': True})

    app.run(host='0.0.0.0', port=8080, debug=True)
