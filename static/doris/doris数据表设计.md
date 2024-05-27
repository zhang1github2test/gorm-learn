### 一、分区及分桶

#### 1、分区partition

分区时的注意事项：

```text 
1、partition列可以指定一列或者多列，分区列必须为key列
2、不论分区列是什么类型，在写分区值时，都需要加双引号
3、分区数量理论上没有上限
4、当不使用 Partition 建表时，系统会自动生成一个和表名同名的，全值范围的 Partition。该Partition 对用户不可见，并且不可删改
5、创建分区时不可添加范围重叠的分区。
```

##### 1.1、range分区

 1）、分区列通常为时间列，以方便的管理新旧数据。

 2）、Range 分区支持的列类型：[DATE,DATETIME,TINYINT,SMALLINT,INT,BIGINT,LARGEINT]

3）、Partition 支持通过 `VALUES LESS THAN (...)` 仅指定上界，系统会将前一个分区的上界作为该分区的下界，生成一个左闭右开的区间。也支持通过 `VALUES [...)` 指定上下界，生成一个左闭右开的区间



* 通过"VALUES[... )"创建

如下面的建表语句，将会创建6个分区的

```txt
CREATE TABLE IF NOT EXISTS example_db.example_range1 (
		`timestamp` DATETIME NOT NULL COMMENT "日志时间",
		`type` INT NOT NULL COMMENT "日志类型",
		`error_code` INT COMMENT "错误码",
		`error_msg` VARCHAR ( 1024 ) COMMENT "错误详细信息",
		`op_id` BIGINT COMMENT "负责人id",
	`op_time` DATETIME COMMENT "处理时间" 
	) DUPLICATE KEY ( `timestamp`, `type`, `error_code` ) 
	PARTITION BY RANGE (`timestamp`) (
	   PARTITION `p2022` VALUES [("1970-01-01 00:00:00"),("2023-01-01 00:00:00")),
		 PARTITION `p2023` VALUES [("2023-01-01 00:00:00"),("2024-01-01 00:00:00")),
		 PARTITION `p2024` VALUES [("2024-01-01 00:00:00"),("2025-01-01 00:00:00")),
		 PARTITION `p2025` VALUES [("2025-01-01 00:00:00"),("2026-01-01 00:00:00")),
		 PARTITION `p2026` VALUES [("2026-01-01 00:00:00"),("2027-01-01 00:00:00")),
		 PARTITION `p2027` VALUES [("2027-01-01 00:00:00"), ("2099-01-01 00:00:00"))
	)
	
	DISTRIBUTED BY HASH ( `type` ) BUCKETS 1 
	PROPERTIES ( "replication_allocation" = "tag.location.default: 1" );
```

检查具体的分区表如下所示：

SHOW PARTITIONS FROM example_db.example_range1;

![image-20240131154554186](C:\Users\admin\AppData\Roaming\Typora\typora-user-images\image-20240131154554186.png)

* 通过" VALUES LESS THAN (...)" 创建 Ranage 分区

  建表示例：

  ```sql
  CREATE TABLE IF NOT EXISTS example_db.example_range_less_than (
  		`timestamp` DATETIME NOT NULL COMMENT "日志时间",
  		`type` INT NOT NULL COMMENT "日志类型",
  		`error_code` INT COMMENT "错误码",
  		`error_msg` VARCHAR ( 1024 ) COMMENT "错误详细信息",
  		`op_id` BIGINT COMMENT "负责人id",
  	`op_time` DATETIME COMMENT "处理时间" 
  	) DUPLICATE KEY ( `timestamp`, `type`, `error_code` ) 
  	PARTITION BY RANGE (`timestamp`) (
  	   PARTITION `p2022` VALUES LESS THAN ("2023-01-01 00:00:00"),
  		 PARTITION `p2023` VALUES LESS THAN ("2024-01-01 00:00:00"),
  		 PARTITION `p2024` VALUES LESS THAN ("2025-01-01 00:00:00"),
  		 PARTITION `p2025` VALUES LESS THAN ("2026-01-01 00:00:00"),
  		 PARTITION `p2026` VALUES LESS THAN ("2027-01-01 00:00:00"),
  		 PARTITION `p2027` VALUES LESS THAN  ("2099-01-01 00:00:00")
  	)
  	
  	DISTRIBUTED BY HASH ( `type` ) BUCKETS 1 
  	PROPERTIES ( "replication_allocation" = "tag.location.default: 1" );
  ```

  查看分区情况：

  SHOW PARTITIONS FROM example_db.example_range_less_than;

  ![image-20240131155018553](C:\Users\admin\AppData\Roaming\Typora\typora-user-images\image-20240131155018553.png)

* 通过" FROM(...) TO (...) INTERVAL ..." 创建 Ranage 分区

  建表示例：

  ```sql
  CREATE TABLE IF NOT EXISTS example_db.example_range_from_to (
  		`timestamp` DATETIME NOT NULL COMMENT "日志时间",
  		`type` INT NOT NULL COMMENT "日志类型",
  		`error_code` INT COMMENT "错误码",
  		`error_msg` VARCHAR ( 1024 ) COMMENT "错误详细信息",
  		`op_id` BIGINT COMMENT "负责人id",
  	`op_time` DATETIME COMMENT "处理时间" 
  	) DUPLICATE KEY ( `timestamp`, `type`, `error_code` ) 
  	PARTITION BY RANGE (`timestamp`) (
  	   FROM ("2000-01-01 00:00:00") TO ("2099-01-01 00:00:00") INTERVAL 1 YEAR
  	)
  	
  	DISTRIBUTED BY HASH ( `type` ) BUCKETS 1 
  	PROPERTIES ( "replication_allocation" = "tag.location.default: 1" );
  ```

  查看对应表的分区信息，上面的人建表语句将创建99个分区

  ​	SHOW PARTITIONS FROM example_db.example_range_from_to;

![image-20240131155425775](C:\Users\admin\AppData\Roaming\Typora\typora-user-images\image-20240131155425775.png)

* 分区的增加和删除

  官方参考文档：https://doris.apache.org/zh-CN/docs/sql-manual/sql-reference/Data-Definition-Statements/Alter/ALTER-TABLE-PARTITION

  新增加的分区不能和原来存在的分区重合，否则新增的分区将会失败。

  只支持VALUES LESS THAN （...）及 VALUES [("value1", ...), ("value1", ...))两种写法

  这里以上面的表example_range_from_to演示增加分区及删除分区

​      目前example_range_from_to表的基于timestamp字段进行分区，每年一个区时间从[2000-01-01 00:00:00，2099-01-01 00:00:00），那么就是说我们新增的分区不能落入到这个区间里面。

​      尝试建立新增一个在此区间的分区，预期将会报错。

建立一个分区小于2098-06-01 05:00:00的分区且名称为p_new

ALTER TABLE example_db.example_range_from_to ADD PARTITION p_new VALUES LESS THAN ("2098-06-01 05:00:00")，执行完该语句后，看到如下报错信息。这里由于新建的分区与已经存在分区存在交集了，所以不允许新建

```log
ALTER TABLE example_db.example_range_from_to ADD PARTITION p_new VALUES LESS THAN ("2098-06-01 05:00:00")
> 1105 - errCode = 2, detailMessage = Range [types: [DATETIME]; keys: [2098-01-01 00:00:00]; ..types: [DATETIME]; keys: [2098-06-01 05:00:00]; ) is intersected with range: [types: [DATETIME]; keys: [2098-01-01 00:00:00]; ..types: [DATETIME]; keys: [2099-01-01 00:00:00]; )
```

然后我们尝试创建一个没有交集的分区

* 新增分区

```sql
ALTER TABLE example_db.example_range_from_to ADD PARTITION p_2099 VALUES LESS THAN ("2100-01-01 05:00:00")
> OK
> 时间: 0.01s
```

发现能够正常创建

* 删除分区

  ```sql
  ALTER TABLE example_db.example_range_from_to DROP PARTITION IF EXISTS p_2099
  > OK
  > 时间: 0.008s
  ```

##### 1.1、list分区

- 分区列支持 `BOOLEAN, TINYINT, SMALLINT, INT, BIGINT, LARGEINT, DATE, DATETIME, CHAR, VARCHAR` 数据类型，分区值为枚举值。只有当数据为目标分区枚举值其中之一时，才可以命中分区。
- Partition 支持通过 `VALUES IN (...)` 来指定每个分区包含的枚举值



建表语句

```sql
CREATE TABLE IF NOT EXISTS example_db.example_list_tbl
(
    `user_id` LARGEINT NOT NULL COMMENT "用户id",
    `date` DATE NOT NULL COMMENT "数据灌入日期时间",
    `timestamp` DATETIME NOT NULL COMMENT "数据灌入的时间戳",
    `city` VARCHAR(20) NOT NULL COMMENT "用户所在城市",
    `age` SMALLINT COMMENT "用户年龄",
    `sex` TINYINT COMMENT "用户性别",
    `last_visit_date` DATETIME REPLACE DEFAULT "1970-01-01 00:00:00" COMMENT "用户最后一次访问时间",
    `cost` BIGINT SUM DEFAULT "0" COMMENT "用户总消费",
    `max_dwell_time` INT MAX DEFAULT "0" COMMENT "用户最大停留时间",
    `min_dwell_time` INT MIN DEFAULT "99999" COMMENT "用户最小停留时间"
)
ENGINE=olap
AGGREGATE KEY(`user_id`, `date`, `timestamp`, `city`, `age`, `sex`)
PARTITION BY LIST(`city`)
(
    PARTITION `p_cn` VALUES IN ("Beijing", "Shanghai", "Hong Kong"),
    PARTITION `p_usa` VALUES IN ("New York", "San Francisco"),
    PARTITION `p_jp` VALUES IN ("Tokyo")
)
DISTRIBUTED BY HASH(`user_id`) BUCKETS 1
PROPERTIES
(
    "replication_num" = "1",
    "storage_cooldown_time" = "2018-01-01 12:00:00"
);
```

查看刚刚新建好表分区情况：

​	SHOW PARTITIONS FROM example_db.example_list_tbl;

![image-20240131171147143](C:\Users\admin\AppData\Roaming\Typora\typora-user-images\image-20240131171147143.png)





#### 2、分桶Bucket

- 如果使用了 Partition，则 `DISTRIBUTED ...` 语句描述的是数据在**各个分区内**的划分规则。如果不使用 Partition，则描述的是对整个表的数据的划分规则。

- 分桶列可以是多列，Aggregate 和 Unique 模型必须为 Key 列，Duplicate 模型可以是 key 列和 value 列。分桶列可以和 Partition 列相同或不同。

- 分桶列的选择，是在查询吞吐和查询并发之间的一种权衡：

  1. 如果选择多个分桶列，则数据分布更均匀。如果一个查询条件不包含所有分桶列的等值条件，那么该查询会触发所有分桶同时扫描，这样查询的吞吐会增加，单个查询的延迟随之降低。这个方式适合大吞吐低并发的查询场景。
  2. 如果仅选择一个或少数分桶列，则对应的点查询可以仅触发一个分桶扫描。此时，当多个点查询并发时，这些查询有较大的概率分别触发不同的分桶扫描，各个查询之间的IO影响较小（尤其当不同桶分布在不同磁盘上时），所以这种方式适合高并发的点查询场景。

- AutoBucket: 根据数据量，计算分桶数。 对于分区表，可以根据历史分区的数据量、机器数、盘数，确定一个分桶。

- 分桶的数量理论上没有上限。

  **关于 Partition 和 Bucket 的数量和数据量的建议**

  

- 一个表的 Tablet 总数量等于 (Partition num * Bucket num)。
- 一个表的 Tablet 数量，在不考虑扩容的情况下，推荐略多于整个集群的磁盘数量。
- 单个 Tablet 的数据量理论上没有上下界，但建议在 1G - 10G 的范围内。如果单个 Tablet 数据量过小，则数据的聚合效果不佳，且元数据管理压力大。如果数据量过大，则不利于副本的迁移、补齐，且会增加 Schema Change 或者 Rollup 操作失败重试的代价（这些操作失败重试的粒度是 Tablet）。
- 当 Tablet 的数据量原则和数量原则冲突时，建议优先考虑数据量原则。
- 在建表时，每个分区的 Bucket 数量统一指定。但是在动态增加分区时（`ADD PARTITION`），可以单独指定新分区的 Bucket 数量。可以利用这个功能方便的应对数据缩小或膨胀。
- 一个 Partition 的 Bucket 数量一旦指定，不可更改。所以在确定 Bucket 数量时，需要预先考虑集群扩容的情况。比如当前只有 3 台 host，每台 host 有 1 块盘。如果 Bucket 的数量只设置为 3 或更小，那么后期即使再增加机器，也不能提高并发度。
- 举一些例子：假设在有10台BE，每台BE一块磁盘的情况下。如果一个表总大小为 500MB，则可以考虑4-8个分片。5GB：8-16个分片。50GB：32个分片。500GB：建议分区，每个分区大小在 50GB 左右，每个分区16-32个分片。5TB：建议分区，每个分区大小在 50GB 左右，每个分区16-32个分片。

> 注：表的数据量可以通过 [`SHOW DATA`](https://doris.apache.org/zh-CN/docs/sql-manual/sql-reference/Show-Statements/SHOW-DATA) 命令查看，结果除以副本数，即表的数据量。

### 二、索引

目前 Doris 主要支持两类索引：

1. 内建的智能索引，包括前缀索引和 ZoneMap 索引。
2. 用户手动创建的二级索引，包括 [倒排索引](https://doris.apache.org/zh-CN/docs/data-table/index/inverted-index)、 [bloomfilter索引](https://doris.apache.org/zh-CN/docs/data-table/index/bloomfilter)、 [ngram bloomfilter索引](https://doris.apache.org/zh-CN/docs/data-table/index/ngram-bloomfilter-index) 和[bitmap索引](https://doris.apache.org/zh-CN/docs/data-table/index/bitmap-index)。

#### 前缀索引

不同于传统的数据库设计，Doris 不支持在任意列上创建索引。Doris 这类 MPP 架构的 OLAP 数据库，通常都是通过提高并发，来处理大量数据的。