# wecom
企业微信SDK

拉取取消息留痕，使用三方扩展
```https://github.com/NICEXAI/WeWorkFinanceSDK```
请将.so 文件拷贝进入系统，并进行配置
下面给出实际案例 dockerfile 配置
```azure
FROM golang:1.18

COPY lib/libWeWorkFinanceSdk_C.so /usr/local/lib

ENV LD_LIBRARY_PATH=/usr/local/lib

RUN cp -r -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
	&& sed -i 's/deb.debian.org/mirrors.aliyun.com/g' /etc/apt/sources.list \
	&& sed -i 's/security.debian.org/mirrors.aliyun.com/g' /etc/apt/sources.list \
	&& apt-get update \
	&& apt-get install -y \
		procps \
		net-tools \
		vim \
	&& echo "alias ll='ls -alF'" >> /root/.bashrc \
	&& echo "alias l='ls -CF'" >> /root/.bashrc 

# RUN apt-get install -y ffmpeg

RUN rm -r /var/lib/apt/lists/*

ENV ROCKETMQ_GO_LOG_LEVEL ERROR

RUN go env \
	&& go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w GO111MODULE=on \
    && go env

EXPOSE 10800

WORKDIR /app

```