# k8sciagent


k8s CI 代理工具 避免直接调用k8s获得过高权限
 
Install 

```bash
go install github.com/luaxlou/k8sciagent@latest
```

Usage
```bash
k8sciagent deploy -a appName -i imageName -v version -e env
```