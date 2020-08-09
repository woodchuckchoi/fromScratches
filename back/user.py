# Built-in
import os

# PIP
from sanic import Sanic
from sanic.response import json

# Custom
from samaria import Samaria
from utils import importConfig

'''
    "DB_SERVICE":   "0.0.0.0",
    "DB_USER":      "root",
    "DB_PASS":      "worbdj12",
    "DB_DATABASE":  "devqueue"
'''

if __name__ == "__main__":
    HOST, USER, PASS, DATABASE = importConfig()

    app     = Sanic(__name__)
    samaria = Samaria(host=HOST, user=USER, password=PASS, database=DATABASE)

    @app.route('/')
    async def all_user(request):
        await samaria.execute('SELECT DISTINCT User.id, (SELECT COUNT(*) FROM Record WHERE Record.user_id = User.idx) AS game_count FROM User ORDER BY game_count DESC, User.id ASC;')
        ret = await samaria.fetchall()
        return json(ret)

    @app.route('/login', methods=['POST'])
    async def login(request):
        if request.method != 'POST':
            return json(status=404)

        loginId = request.json['loginId']
        loginPw = request.json['loginPw']

        await samaria.execute("SELECT IF(SHA2(CONCAT(creation_date, '{}'), 512) = password, 1, 0) FROM User WHERE id = '{}'".format(loginPw, loginId))
        
        ret = await samaria.fetchall()
        if ret[0][0]:
            return json({"result": True})
        return json({"result": False})

    app.run(host='0.0.0.0', port=8080, debug=True)
