# erero
简单的errors包，和菠萝菠萝蜜的相同，erero，就是个简单的错误包，当发生错误的时候自动打印日志，假如名字叫errors就有点烂大街啦，还得解决包名的冲突问题，比如和标准errors或者github.com/pkg/errors的冲突。因此随便起个名字吧，假如叫ero就过短啦，还容易和变量名冲突。因此就用erero吧。

因为经常会遇到说
```
if err != nil {
    showLogMessage(err)
    return err
}
```
的情况。

更有甚者会遇到
```
if err != nil {
    showLogMessage(errors.WithMessage(err, "wrong"))
    return errors.WithMessage(err, "wrong")
}
```
的情况

这时候即使是合并变量
```
if err != nil {
    err = errors.WithMessage(err, "wrong")
    showLogMessage(err)
    return err
}
```
也依然觉得很别扭啊。

因此有了这个包，这个包平替 github.com/pkg/errors 的全部功能，当你返回错误需要打印时，就用它，而当你不需要打印日志时，就用别的 errors 包就行。

当然由于我的日志可能不是你想要的，因此你可以自己实现自己的 errors 包，配合其它 errors 交替使用即可。
