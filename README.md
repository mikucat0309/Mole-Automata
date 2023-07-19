# 摩爾：自動人形 (Mole: Automata)

## 功能

- 餐廳
    - 製作新的料理
    - 收取完成的料理
    - 倒掉過期的料理

## 使用方式

重新命名 `template.env` 為 `.env` 並修改對應資訊

| field | description |
| ------ | --- |
| `USER` | 帳號 |
| `PASSWORD` | 密碼 |
| `DISH_ID` | 料理 ID (請參照下方對照表) |

輸入指令運行本程式

```go
go run .\cmd\main.go
```

## 逆向工程

請見 [reverse/README.md](reverse/README.md)

## 料理 ID 對照表

| 料理 ID | 料理名稱 | 料理 ID | 料理名稱 | 料理 ID | 料理名稱 |
| ------- | -------- | ------- | -------- | ------- | -------- |
| 1340001 | 清炒毛毛豆 | 1340002 | 醬爆雪頂菇 | 1340003 | 果凍花咖喱飯 |
| 1340004 | 咕嚕果奶酥 | 1340005 | 咕唧蛋撻 | 1340006 | 藍莓彩羽湯 |
| 1340007 | 陽光酥油肉鬆 | 1340008 | 開心果蔬餅 | 1340009 | 傑克南瓜酥 |
| 1340010 | 串烤海琴花 | 1340011 | 蜜汁卡蘭花 | 1340012 | 尤尤什錦百燴 |
| 1340013 | 摩雅霜降魚排 | 1340014 | 四季百花羹 | 1340015 | 驚奇松塔餅 |
| 1340016 | 彩虹漿果燴蝶魚 | 1340017 | 拉姆小饅頭 | 1340018 | 神秘湖蟹蓉泡芙 |
| 1340019 | 金槍魚彩菇披薩 | 1340020 | 梅森鮮果湯 | 1340021 | 玫瑰香蒸塔塔酥 |
| 1340022 | 雪丁丁 | 1340023 | 荷香爽飲 | 1340024 | 豆香南瓜飯 |
| 1340025 | 葡萄冰沙 | 1340026 | 七彩草莓霜淇淋 | 1340027 | 十二星座餅乾 |
| 1340028 | 酸甜冰泥 | 1340029 | 雲朵松糕 | 1340030 | 愛心蛋蛋堡 |
| 1340031 | 蜂蜜水果熱飲 | 1340032 | 黑胡椒沙朗牛排 | 1340033 | 胡蘿蔔蓋飯 |
| 1340034 | 拉姆七彩沙拉 | 1340035 | 拉姆曲奇餅 | 1340036 | 拉姆鴨蛋麵 |
| 1340037 | 麻辣小龍蝦 | 1340039 | 葡萄石榴派 | 1340040 | 七寶飯 |
| 1340041 | 清爽薯片 | 1340042 | 濃情南瓜盅 | 1340043 | 果蔬蛋捲 |
| 1340044 | 雙層巧克力蛋糕 | 1340045 | 水果飯糰 | 1340046 | 冰糖葫蘆 |
| 1340048 | 小惡魔披薩 | 1340049 | 音符巧克力 | 1340050 | 月亮船優酪乳雪糕 |
| 1340051 | 焦糖拉姆布丁 | 1340052 | 多彩粽子 | 1340053 | 繽紛土豆泥 |
| 1340054 | 檸檬蛋糕 | 1340055 | 草莓冰爽碎碎冰 | 1340056 | 彩虹冰爽 |
| 1340057 | 甜橙布丁冷飲 | 1340058 | 鮮蝦時蔬拼盤 | 1340059 | 黑糊糊粥 |
| 1340060 | 多彩塔塔酥 | 1340061 | 摩摩蛋黃月餅 | 1340062 | 清爽布丁 |
| 1340063 | 秘製小煎餅 | 1340064 | 水果豐收蛋糕 | 1340065 | 水蛋堡 |
| 1340066 | 南瓜布丁 | 1340067 | 香濃薑湯 | 1340068 | 水果冰晶粽 |
| 1340069 | 夏日清爽涼麵 | 1340070 | 燒烤拉姆 | 1340071 | 夏威夷冰飲 |
| 1340072 | 摩摩蛋包飯 | 1340073 | 櫻花冰淇淋蛋糕 | 1340074 | 漿果糯米糍 |

