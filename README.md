# go-macho

v0.1.0 [![Build Status](https://travis-ci.org/wefound/macho-go.png?branch=master)](https://travis-ci.org/wefound/macho-go)


---
The go-macho package is a **Mach-O** parser for golang, as an supplement for official package debug/macho.

It's **under development** now.

这是一个 macho 文件的小工具，补充官方库不支持获取符号表 uuid 信息的功能。依然处于开发阶段，属于玩具，不建议使用到任何环境。

## Get
    go get github.com/jqnote/macho-go

## Do Sth.

```golang
var name string = "path/to/your/macho/file"

macho, _ := NewMacho(name)
uuids := macho.GetUUIDS()
fmt.Printf("UUIDS\n%v\n", uuids)

var uuid string
uuid, _ = macho.GetUUIDByName("arm64")
fmt.Printf("uuid of arm64\n%v\n", uuid)


```
## Other

 - [debug/macho](https://golang.org/pkg/debug/macho/) _official_
 - [Mach-O Executables](https://www.objc.io/issues/6-build-tools/mach-o-executables/)
