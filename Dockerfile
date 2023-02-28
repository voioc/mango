FROM public-env-mirror-service-registry.cn-beijing.cr.aliyuncs.com/dist/golang:1.16

WORKDIR /tmp/go-sample-app
# 准备工作
#RUN export 
COPY go.mod .
COPY go.sum .
COPY . .

#加入git访问权限
COPY .netrc /root/.netrc

# 编译
RUN go env -w GOPRIVATE=codeup.aliyun.com
RUN go env
RUN GOPROXY="https://goproxy.cn" GO111MODULE=on go build -o ./out/go-sample-app .
#RUN go build -o ./out/go-sample-app .
RUN chmod +x ./out/go-sample-app

ARG envType=test
COPY conf/env/${envType}.toml conf/env/env.toml
# 执行编译生成的二进制文件
CMD ["./out/go-sample-app","-c","conf/env/env.toml"]
# 暴露端口
EXPOSE 80