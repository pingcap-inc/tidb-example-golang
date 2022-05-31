# Expected-Output

## gorm

```
/Library/Developer/CommandLineTools/usr/bin/make -C gorm
make build run
go build -o bin/gorm-example
./bin/gorm-example

2022/05/19 01:26:17 /Users/cheese/go/pkg/mod/gorm.io/driver/mysql@v1.3.3/migrator.go:228
[0.248ms] [rows:-] SELECT DATABASE()

2022/05/19 01:26:17 /Users/cheese/go/pkg/mod/gorm.io/driver/mysql@v1.3.3/migrator.go:231
[0.784ms] [rows:1] SELECT SCHEMA_NAME from Information_schema.SCHEMATA where SCHEMA_NAME LIKE 'test%' ORDER BY SCHEMA_NAME='test' DESC limit 1

2022/05/19 01:26:17 /Users/cheese/GolandProjects/tidb-example-golang/gorm/gorm.go:47
[5.459ms] [rows:-] SELECT count(*) FROM information_schema.tables WHERE table_schema = 'test' AND table_name = 'player' AND table_type = 'BASE TABLE'

2022/05/19 01:26:17 /Users/cheese/go/pkg/mod/gorm.io/driver/mysql@v1.3.3/migrator.go:228
[0.323ms] [rows:-] SELECT DATABASE()

2022/05/19 01:26:17 /Users/cheese/go/pkg/mod/gorm.io/driver/mysql@v1.3.3/migrator.go:231
[1.049ms] [rows:1] SELECT SCHEMA_NAME from Information_schema.SCHEMATA where SCHEMA_NAME LIKE 'test%' ORDER BY SCHEMA_NAME='test' DESC limit 1

2022/05/19 01:26:17 /Users/cheese/go/pkg/mod/gorm.io/driver/mysql@v1.3.3/migrator.go:148
[3.036ms] [rows:-] SELECT * FROM `player` LIMIT 1

2022/05/19 01:26:17 /Users/cheese/go/pkg/mod/gorm.io/driver/mysql@v1.3.3/migrator.go:166
[1.923ms] [rows:-] SELECT column_name, column_default, is_nullable = 'YES', data_type, character_maximum_length, column_type, column_key, extra, column_comment, numeric_precision, numeric_scale , datetime_precision FROM information_schema.columns WHERE table_schema = 'test' AND table_name = 'player' ORDER BY ORDINAL_POSITION

2022/05/19 01:26:18 /Users/cheese/GolandProjects/tidb-example-golang/gorm/gorm.go:53
[3.935ms] [rows:0] INSERT INTO `player` (`id`,`coins`,`goods`) VALUES ('test',1,1) ON DUPLICATE KEY UPDATE `coins`=VALUES(`coins`),`goods`=VALUES(`goods`)

2022/05/19 01:26:18 /Users/cheese/GolandProjects/tidb-example-golang/gorm/gorm.go:59
[1.421ms] [rows:1] SELECT * FROM `player` WHERE id = 'test'
getPlayer: {ID:test Coins:1 Goods:1}

2022/05/19 01:26:18 /Users/cheese/GolandProjects/tidb-example-golang/gorm/gorm.go:80
[8.693ms] [rows:1] SELECT count(*) FROM `player`
countPlayers: 3841

2022/05/19 01:26:18 /Users/cheese/GolandProjects/tidb-example-golang/gorm/gorm.go:85
[1.537ms] [rows:3] SELECT * FROM `player` LIMIT 3
print 1 player: {ID:test Coins:1 Goods:1}
print 2 player: {ID:2d1a4cc2-a49e-4ad1-aea0-3b97b06ad400 Coins:8081 Goods:7887}
print 3 player: {ID:bfe25359-fc21-40c7-9049-b024fff02e36 Coins:1847 Goods:4059}

2022/05/19 01:26:18 /Users/cheese/GolandProjects/tidb-example-golang/gorm/gorm.go:98
[10.434ms] [rows:2] INSERT INTO `player` (`id`,`coins`,`goods`) VALUES ('1',100,0) ON DUPLICATE KEY UPDATE `coins`=VALUES(`coins`),`goods`=VALUES(`goods`)

2022/05/19 01:26:18 /Users/cheese/GolandProjects/tidb-example-golang/gorm/gorm.go:99
[3.990ms] [rows:2] INSERT INTO `player` (`id`,`coins`,`goods`) VALUES ('2',114514,20) ON DUPLICATE KEY UPDATE `coins`=VALUES(`coins`),`goods`=VALUES(`goods`)

buyGoods:
    => this trade will fail

2022/05/19 01:26:18 /Users/cheese/GolandProjects/tidb-example-golang/gorm/gorm.go:119
[1.526ms] [rows:1] SELECT * FROM `player` WHERE id = '2' FOR UPDATE

2022/05/19 01:26:18 /Users/cheese/GolandProjects/tidb-example-golang/gorm/gorm.go:128
[1.445ms] [rows:1] SELECT * FROM `player` WHERE id = '1' FOR UPDATE

buyGoods:
    => this trade will success

2022/05/19 01:26:18 /Users/cheese/GolandProjects/tidb-example-golang/gorm/gorm.go:119
[1.669ms] [rows:1] SELECT * FROM `player` WHERE id = '2' FOR UPDATE

2022/05/19 01:26:18 /Users/cheese/GolandProjects/tidb-example-golang/gorm/gorm.go:128
[1.412ms] [rows:1] SELECT * FROM `player` WHERE id = '1' FOR UPDATE

2022/05/19 01:26:18 /Users/cheese/GolandProjects/tidb-example-golang/gorm/gorm.go:137
[0.603ms] [rows:1] UPDATE player set goods = goods + -2, coins = coins + 100 WHERE id = '2'

2022/05/19 01:26:18 /Users/cheese/GolandProjects/tidb-example-golang/gorm/gorm.go:141
[0.475ms] [rows:1] UPDATE player set goods = goods + 2, coins = coins + -100 WHERE id = '1'

[buyGoods]:
    'trade success'
```

## sqldriver

```
/Library/Developer/CommandLineTools/usr/bin/make -C sqldriver
make mysql build run
mycli --host 127.0.0.1 --port 4000 -u root --no-warn < sql/dbinit.sql
go build -o bin/sql-driver-example
./bin/sql-driver-example
getPlayer: {ID:test Coins:1 Goods:1}
countPlayers: 1920
print 1 player: {ID:test Coins:1 Goods:1}
print 2 player: {ID:87f59982-d3ff-49eb-b720-c3a74fbab215 Coins:8081 Goods:7887}
print 3 player: {ID:990817b2-074d-4f80-bd03-d28bfc7a35a0 Coins:1847 Goods:4059}

buyGoods:
    => this trade will fail

buyGoods:
    => this trade will success

[buyGoods]:
    'trade success'
```

## http service

```
make build run
go build -o bin/http
./bin/http

2022/05/19 02:55:47 /Users/cheese/go/pkg/mod/gorm.io/driver/mysql@v1.3.3/migrator.go:228
[0.217ms] [rows:-] SELECT DATABASE()

2022/05/19 02:55:47 /Users/cheese/go/pkg/mod/gorm.io/driver/mysql@v1.3.3/migrator.go:231
[0.778ms] [rows:1] SELECT SCHEMA_NAME from Information_schema.SCHEMATA where SCHEMA_NAME LIKE 'test%' ORDER BY SCHEMA_NAME='test' DESC limit 1

2022/05/19 02:55:47 /Users/cheese/GolandProjects/tidb-example-golang/http/service.go:35
[8.220ms] [rows:-] SELECT count(*) FROM information_schema.tables WHERE table_schema = 'test' AND table_name = 'player' AND table_type = 'BASE TABLE'

2022/05/19 02:55:47 /Users/cheese/go/pkg/mod/gorm.io/driver/mysql@v1.3.3/migrator.go:228
[0.352ms] [rows:-] SELECT DATABASE()

2022/05/19 02:55:47 /Users/cheese/go/pkg/mod/gorm.io/driver/mysql@v1.3.3/migrator.go:231
[1.022ms] [rows:1] SELECT SCHEMA_NAME from Information_schema.SCHEMATA where SCHEMA_NAME LIKE 'test%' ORDER BY SCHEMA_NAME='test' DESC limit 1

2022/05/19 02:55:47 /Users/cheese/go/pkg/mod/gorm.io/driver/mysql@v1.3.3/migrator.go:148
[1.453ms] [rows:-] SELECT * FROM `player` LIMIT 1

2022/05/19 02:55:47 /Users/cheese/go/pkg/mod/gorm.io/driver/mysql@v1.3.3/migrator.go:166
[1.850ms] [rows:-] SELECT column_name, column_default, is_nullable = 'YES', data_type, character_maximum_length, column_type, column_key, extra, column_comment, numeric_precision, numeric_scale , datetime_precision FROM information_schema.columns WHERE table_schema = 'test' AND table_name = 'player' ORDER BY ORDINAL_POSITION

2022/05/19 02:56:12 /Users/cheese/GolandProjects/tidb-example-golang/http/service.go:45
[1.016ms] [rows:1] INSERT INTO `player` (`id`,`coins`,`goods`) VALUES ('4db45278-c594-4574-9928-08a10f60f317',100,20)

2022/05/19 02:56:12 /Users/cheese/GolandProjects/tidb-example-golang/http/service.go:45
[1.018ms] [rows:1] INSERT INTO `player` (`id`,`coins`,`goods`) VALUES ('6b3a3392-bbc6-4a50-bc00-a1a15b3b27cf',100,20)

2022/05/19 02:56:12 /Users/cheese/GolandProjects/tidb-example-golang/http/service.go:45
[1.021ms] [rows:1] INSERT INTO `player` (`id`,`coins`,`goods`) VALUES ('84237628-22be-4931-9aff-22b6289f2954',100,20)

2022/05/19 02:56:12 /Users/cheese/GolandProjects/tidb-example-golang/http/service.go:45
[1.044ms] [rows:1] INSERT INTO `player` (`id`,`coins`,`goods`) VALUES ('70794d37-f86a-41dc-ae96-4ae0a65d7011',100,20)

2022/05/19 02:56:12 /Users/cheese/GolandProjects/tidb-example-golang/http/service.go:45
[1.021ms] [rows:1] INSERT INTO `player` (`id`,`coins`,`goods`) VALUES ('ce36d6dd-4507-4643-b4f3-b42911593697',100,20)

2022/05/19 02:56:12 /Users/cheese/GolandProjects/tidb-example-golang/http/service.go:45
[0.969ms] [rows:1] INSERT INTO `player` (`id`,`coins`,`goods`) VALUES ('8ada052f-6c9e-4998-b232-869359d2fe8e',100,20)

2022/05/19 02:56:12 /Users/cheese/GolandProjects/tidb-example-golang/http/service.go:45
[0.969ms] [rows:1] INSERT INTO `player` (`id`,`coins`,`goods`) VALUES ('fc5c0b2a-a6bc-4136-9062-193df0f8b910',100,20)

2022/05/19 02:56:12 /Users/cheese/GolandProjects/tidb-example-golang/http/service.go:45
[1.000ms] [rows:1] INSERT INTO `player` (`id`,`coins`,`goods`) VALUES ('a595379b-def9-49af-8a9f-5eba31bc8692',100,20)

2022/05/19 02:56:12 /Users/cheese/GolandProjects/tidb-example-golang/http/service.go:45
[1.052ms] [rows:1] INSERT INTO `player` (`id`,`coins`,`goods`) VALUES ('84fba798-9114-43fc-86f1-9ac992a3e4c2',100,20)

2022/05/19 02:56:12 /Users/cheese/GolandProjects/tidb-example-golang/http/service.go:45
[1.256ms] [rows:1] INSERT INTO `player` (`id`,`coins`,`goods`) VALUES ('16e300cc-a170-4929-9c39-9b833330baa5',100,20)

2022/05/19 02:56:12 /Users/cheese/GolandProjects/tidb-example-golang/http/service.go:51
[1.531ms] [rows:1] SELECT * FROM `player` WHERE id = '1'

2022/05/19 02:56:12 /Users/cheese/GolandProjects/tidb-example-golang/http/service.go:58
[1.483ms] [rows:3] SELECT * FROM `player` LIMIT 3

2022/05/19 02:56:12 /Users/cheese/GolandProjects/tidb-example-golang/http/service.go:51
[1.174ms] [rows:0] SELECT * FROM `player` WHERE id = 'page'

2022/05/19 02:56:12 /Users/cheese/GolandProjects/tidb-example-golang/http/service.go:65
[2.994ms] [rows:1] SELECT count(*) FROM `player`

2022/05/19 02:56:12 /Users/cheese/GolandProjects/tidb-example-golang/http/service.go:73
[1.828ms] [rows:1] SELECT * FROM `player` WHERE id = '1' FOR UPDATE
```

## http request

```
./request.sh
loop to create 10 players:
1111111111

get player 1:
{"ID":"1","Coins":0,"Goods":2}

get players by limit 3:
[{"ID":"test","Coins":1,"Goods":1},{"ID":"011603d4-1b76-4871-95e5-cfc0e6a9872f","Coins":8181,"Goods":7877},{"ID":"52bbc281-81bd-46a6-8f3b-42138e052285","Coins":1747,"Goods":4069}]

get first players:
{"ID":"","Coins":0,"Goods":0}

get players count:
1934

trade by two players:
false%
```