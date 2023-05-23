### 运送系统

## Getting Started

```shell
# init data and get database prepared.
go run init/init.go > init.sql
sqlite3 order.db
.read init.sql

# valid range of id is [1,1000]
go run --user-id=id
```

## Testing
我自己添加了两个测试用例，在test package下的delivery_test.go中,如果想要测试更多可以添加tests切片即可

```shell
go test test/delivery_test.go
```

## explanation of weight algorithm in init.go
1.首先生成1到100, 基于 1/w 权重的分布切片 weightIntegers\
2.使用19305作为因子进行计算时，得到的weightIntegers数组长度为100099，近似于100000\
3.因为重量计算价格时，在(w-1,w]区间的重量都按照w来进行计算，故将weightIntegers打乱重排，遍历后随机生成(w-1,w]的随机数
