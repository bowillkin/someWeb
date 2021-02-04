FROM registry.cn-hangzhou.aliyuncs.com/will_default/golang-builder:2.0 as builder
ENV GOPROXY="https://goproxy.baidu.com/,direct" GONOSUMDB="github.com/bowillkin/proto" GO111MODULE="on"
WORKDIR /builder/
# 把那些最不容易发生变化的文件的拷贝操作放在较低的镜像层中，这样在重新 build 镜像时就会使用前面 build 产生的缓存。
COPY go.* ./
# can cached
RUN go mod download
COPY . ./
# 使用交叉编译， 这一层除非代码一点没动， 否则不可能使用缓存
RUN CGO_ENABLED=0 go build -o app *.go

FROM registry.cn-hangzhou.aliyuncs.com/will_default/alpine-shanghai:1.0
COPY --from=builder /builder/app .

CMD ["/app"]