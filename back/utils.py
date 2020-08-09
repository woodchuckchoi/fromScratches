import json
import os

CONFIG = './testConfig.json'

def importConfig():
    with open(CONFIG, 'r') as f:
        fallbackConfig = json.loads(f.read())

    HOST    = os.environ['DB_SERVICE'] if 'DB_SERVICE' in os.environ else fallbackConfig['DB_SERVICE']
    USER    = os.environ['DB_USER'] if 'DB_USER' in os.environ else fallbackConfig['DB_USER']
    PASS    = os.environ['DB_PASS'] if 'DB_PASS' in os.environ else fallbackConfig['DB_PASS']
    DATABASE= os.environ['DB_DATABASE'] if 'DB_DATABASE' in os.environ else fallbackConfig['DB_DATABASE']

    return HOST, USER, PASS, DATABASE