# 逆向工程

## 簡介

Flash 程式是由 ActionScript 撰寫而成，swf 檔案是封裝檔，類似 Java 的 jar 封裝檔。

## 準備環境

### Flash Player

Adobe 已經移除下載連結，但摩爾莊園官網有提供

[http://asupdate.917play.com.tw/download/flashplayer_32_sa.zip](http://asupdate.917play.com.tw/download/flashplayer_32_sa.zip)

### 摩爾莊園程式連結

- 主程式 [http://mole.61.com.tw/Client.swf](http://mole.61.com.tw/Client.swf)
- 共同函式庫 [http://mole.61.com.tw/ClientCommonDLL.swf](http://mole.61.com.tw/ClientCommonDLL.swf)
- 通訊函式庫 [http://mole.61.com.tw/ClientSocketDLL.swf](http://mole.61.com.tw/ClientSocketDLL.swf)
- 應用函式庫 [http://mole.61.com.tw/ClientAppDLL.swf](http://mole.61.com.tw/ClientAppDLL.swf)

### 反編譯工具

- [ffdec](https://github.com/jindrapetrik/jpexs-decompiler)

Flash 反組譯工具

- [decode-amf3.py](decode-amf3.py)

swf 封裝許多 amf3 格式的二進位資料，轉成 JSON 方便閱讀與處理

### 流量分析工具

- [WireShark](https://www.wireshark.org/download.html)

- [摩爾莊園協定 dissector lua plugin for WireShark](dissector.lua)
