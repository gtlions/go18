# 🌿gos10i

gos10i, golang十八般武艺(shibabanwuyi-s10i).

![](sources/sj.jpg)


> 《水浒传》言说：史进每日求王教头点拨十八般武艺，一一从头指教。那十八般武艺？矛锤弓弩铳，鞭锏剑链挝，斧钺并戈戟，牌棒与枪杈。

# conver.go

- X2String 任意对象转换成字符串

# encoding.go

- GbkToUtf8 字符串编码gbk->utf-8
- Utf8ToGbk 字符串编码utf-8->gbk

# random.go

- XRandomInt 生成指定长度的随机数字字符串
- XRandomString 生成指定长度的随机字符串
- XRandomIntRange 生成指定区间的整数随机数

# slice.go
- XIsInSlice 判断元素item是否在切片内

# time.go

- XMonthDayList 当前月份日期列表
- XMonthDayCnt 当前月份天数
- XOrderNoFromNow 基于当前时间的字符串，可作为订单号使用，格式：20060102150405000000000
- XChinaMonth 中国月份，「一」~「十二」月
- XChinaWeekday 中国星期几，周「一」～「日」或星期「一」～「日」
- XWeekdayInt 日期的本周的第『n』天，1-7，每周的第一天为周一
- XDayLast235959 返回天的最后时间23:59:59.999999999
- XWeekFirst 周的第一天的时间
- XWeekLast 周的最后一天的时间
- XWeekLast235959 返回周的最后一天的时间23:59:59.999999999
- XMonthFirst 月的第一天的时间
- XMonthLast 月的最后一天的时间
- XMonthLast235959 返回月的最后一天的时间23:59:59.999999999
- XYearFirst 返回年的第一天的时间
- XYearLast 返回年的最后一天的时间
- XYearLast235959 返回年的最后一天的时间23:59:59.999999999
- XZeroTime 0点时间

其它的基于时间戳生成可读唯一订单号等

# aliyunoss

> 阿里云OSS存储相关

# aliyunsms

> 阿里云SMS短信服务相关

# ~~comstring~~

> discard

# ebi

> 电银支付相关

- UnifiedOrder 支付统一下单
    > ParseNotify 解析支付异步通知的参数
- QueryOrder 支付订单查询
- RefundOrder 订单退款
    > ParseRefuncNotify 解析退款异步通知的参数

TODO
- ~~统一下单~~
    - ~~WX_NATIVE~~
    - ~~AL_NATIVE~~
    - ~~WX_JSAPI~~
    - ...
- ~~订单查询~~
- ~~订单退款~~

# tenxunloc

> 腾讯位置服务


# getui

> 个推服务
   
   Android/IOS App的消息通知推送
