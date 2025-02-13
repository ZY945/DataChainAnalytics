# 爬虫获取数据

1.获取html然后解析  
curl html -o gold.html

# 黄金价格采集器

## 黄金服务提供api
### alltick
https://alltick.co/faqs  
文档:https://docs.google.com/spreadsheets/d/1avkeR1heZSj6gXIkDeBt8X3nv4EzJetw4yFuKjSDYtA/edit?gid=665777415#gid=665777415
请求code文档:
https://github.com/alltick/alltick-realtime-forex-crypto-stock-tick-finance-websocket-api/blob/main/product_code_list_commodities_gold_cn.md
## 内嵌页面
https://s.tradingview.com/goldprice/widgetembed/?hideideas=1&overrides=%7B%7D&enabled_features=%5B%5D&disabled_features=%5B%5D&locale=en#%7B%22symbol%22%3A%22TVC%3AGOLD%22%2C%22frameElementId%22%3A%22tradingview_a2026%22%2C%22interval%22%3A%22D%22%2C%22hide_side_toolbar%22%3A%220%22%2C%22allow_symbol_change%22%3A%221%22%2C%22save_image%22%3A%221%22%2C%22watchlist%22%3A%22TVC%3AGOLD%5Cu001fTVC%3ASILVER%5Cu001fTVC%3APLATINUM%5Cu001fTVC%3APALLADIUM%5Cu001fTVC%3AGOLDSILVER%5Cu001fTVC%3AUSOIL%5Cu001fOANDA%3AEURUSD%5Cu001fFX_IDC%3AUSDJPY%5Cu001fINDEX%3AHUI%5Cu001fINDEX%3AXAU%5Cu001fCOINBASE%3ABTCUSD%22%2C%22details%22%3A%221%22%2C%22studies%22%3A%22%5B%5D%22%2C%22theme%22%3A%22White%22%2C%22style%22%3A%221%22%2C%22timezone%22%3A%22America%2FNew_York%22%2C%22hideideasbutton%22%3A%221%22%2C%22withdateranges%22%3A%221%22%2C%22show_popup_button%22%3A%221%22%2C%22studies_overrides%22%3A%22%7B%7D%22%2C%22utm_source%22%3A%22goldprice.org%22%2C%22utm_medium%22%3A%22widget%22%2C%22utm_campaign%22%3A%22chart%22%2C%22utm_term%22%3A%22TVC%3AGOLD%22%2C%22page-uri%22%3A%22goldprice.org%2F%22%7D