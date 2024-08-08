# erero
简单的errors包，读音和菠萝菠萝蜜的音节节奏相同，erero，就是个简单的错误包，当发生错误是记录错误的位置，同时打印日志，通过日志显示错误位置。

erero 的包名不和标准 errors 或者 github.com/pkg/errors 的有冲突，也不和 "github.com/go-kratos/kratos/v2/errors" 有冲突。

# 背景
经常会遇到这种的代码：
```
if err != nil {
    log.Error(err)
    return err
}
```

或者会遇到这样的代码：
```
if err != nil {
    log.Error(errors.WithMessage(err, "wrong"))
    return errors.WithMessage(err, "wrong")
}
```

这时候即使是合并变量:
```
if err != nil {
    err = errors.WithMessage(err, "wrong")
    log.Error(err)
    return err
}
```
也依然觉得这很别扭啊。

这个包封装 "github.com/pkg/errors" 和 "go.uber.org/zap" 日志的逻辑，当出错且需要打印日志时，就用它，而当不需要打印日志时，就用别的 errors，在项目中交替使用。
