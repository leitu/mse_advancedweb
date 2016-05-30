#需求
请确保你已安装:

 * docker-engine
 * docker-compose
 * curl/postman

##运行命令

取得源码

```bash
$git clone https://github.com/leitu/mse_advancedweb.git mse
or
$tar -xvf mse.tar.gz
```

运行命令

```bash
$cd mse/code
$docker-compose run --service-ports web
```