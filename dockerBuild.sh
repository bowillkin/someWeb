GOPROJECT="some-web"
current=`date "+%Y-%m-%d %H:%M:%S"`
timeStamp=`date -d "$current" +%s`
#将current转换为时间戳，精确到毫秒
currentTimeStamp=$((timeStamp*1000+`date "+%N"`/1000000))
echo '2905058'|sudo docker ps
sudo docker login --username=444944454@qq.com registry.cn-hangzhou.aliyuncs.com
sudo -S docker build -t registry.cn-hangzhou.aliyuncs.com/will_default/$GOPROJECT:$currentTimeStamp .
sudo -S docker push registry.cn-hangzhou.aliyuncs.com/will_default/$GOPROJECT:$currentTimeStamp