const Consul = require('consul');

const appName = process.env.APP_NAME || 'node-app';
const appPort = parseInt(process.env.APP_PORT || 3000);
const appHost = process.env.APP_HOST || 'localhost';
const consulHost = process.env.CONSUL_HOST || 'localhost';
const consulPort = parseInt(process.env.CONSUL_PORT || 8500);

class ConsulConfig {
    constructor () {
        // 初始化 consul
        this.consul = new Consul({
            host: consulHost,
            port: consulPort,
            promisify: true,
        });
        
        // 服务注册与健康检查配置
        this.consul.agent.service.register({
            name: appName,
            address: appHost,
            port: appPort,
            check: {
                http: `http://${appHost}:${appPort}/health`,
                interval: '10s',
                timeout: '5s',
            }
        }, function(err, result) {
            if (err) {
                console.error(err);
                throw err;
            }

            console.log(appName + ' 注册成功！');
        })
    }
    
    async getConfig(key) {
        const result = await this.consul.kv.get(key);

        if (!result) {
            return Promise.reject(key + '不存在');
        }

        return JSON.parse(result.Value);
    }
    
    // 读取 user 配置简单封装
    async getUserConfig(key) {
        const result = await this.getConfig('develop/user');

        if (!key) {
            return result;
        }

        return result[key];
    }

	// 更新 user 配置简单封装
    async setUserConfig(key, val) {
        const user = await this.getConfig('develop/user');

        user[key] = val;

        return this.consul.kv.set('develop/user', JSON.stringify(user))
    }
}

module.exports = ConsulConfig;
