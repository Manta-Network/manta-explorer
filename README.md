
#Manta Explorer

## QuickStart

### Requirement

* Linux / Mac OSX
* Git
* Golang 1.12.4+
* Redis 3.0.4+
* MySQL 5.6+
* Node 8.9.0+

### Install

```bash
./build.sh build

//UI
cd ui && yarn 
```

### Config

#### Init config file 

```bash
cd cmd
./subscan install
```
**Attention**
Install configuration requires mysql to connect without password, if mysql has set the root password, we suggest to set the database manually refer to `internal/script/install.sh`.

#### Set

1. Redis  configs/redis.toml

> addr： redis host and port (default: 127.0.0.1:6379)

2. Mysql  configs/mysql.toml

> host: mysql host (default: 127.0.0.1)
> user: mysql user (default: root)
> pass: mysql user passwd (default: "")
> db:   mysql db name (default: "subscan")

3. Http   configs/http.toml

> addr: local http server port (default: 0.0.0.0:4399)

4. Env.go util/env.go

> WSEndPoint: set connected chain (default: GetEnv("CHAIN_WS_ENDPOINT", "ws://127.0.0.1:9944"))
> NetworkNode: set type of network node (default: GetEnv("NETWORK_NODE", "manta"))

### Usage

- Start DB

**Make sure you have started redis and mysql, also you need to clear the redis cache before use it**

- Start 
```bash
cd cmd
./subscan start substrate
```
- End 
```bash
cd cmd
./subscan stop substrate
```

### UI

The ui part is built with [nuxt.js](https://nuxtjs.org/) and [amis](https://github.com/baidu/amis)

Please change proxy target in nuxt.config.js to your server name in development.

```js
proxy: {
   "/api": {
      target: "https://your_server_name.com",
      secure: false,
      changeOrigin: true,
      pathRewrite: {
         "^/api": "/api"
      }
   },
}
```

Please change browserBaseURL in nuxt.config.js to your server name in production.

```js
axios: {
   proxy: process.env.NODE_ENV !== 'production',
    browserBaseURL: process.env.NODE_ENV !== 'production' ? "" : "https://your_server_name.com"
},
```
Start api server
- Api Server
```bash
cd cmd
./subscan
```

Start UI
```
cd ui && yarn dev
```
