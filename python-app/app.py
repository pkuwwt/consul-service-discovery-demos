
import os
from flask import Flask, jsonify
from flask_consulate import Consul

app = Flask(__name__)

port = 5000
host = os.getenv('APP_HOST', 'localhost')
name = os.getenv('APP_NAME', 'python-app')
tags = [i.strip() for i in os.getenv('APP_TAGS', 'webserver').split(',')]
consul_host = os.getenv('CONSUL_HOST', 'localhost')
consul_port = os.getenv('CONSUL_PORT', '8500')
consul_interval = os.getenv('CONSUL_CHECK_INTERVAL', '10s')

consul = Consul(
        app,
        consul_host=consul_host,
        consul_port=consul_port,
        )
consul.apply_remote_config(namespace='mynamespace/')
consul.register_service(
        name=name,
        interval=consul_interval,
        tags=tags,
        port=port,
        httpcheck=f'http://{host}:{port}/health',
        )

@app.route('/api/python/hello')
def hello():
    return 'hello'

@app.route('/health')
def health_check():
    return 'OK!', 200

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=port, debug=True)
