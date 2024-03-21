# gorm实战学习

## 一、gorm入门

### 1、安装gorm

这里我们使用GoLand编辑器，然后golang的版本为1.22.0，项目的依赖通过Go mod来进行管理。

```shell
 go get -u gorm.io/gorm@v1.25.8
 go get -u gorm.io/driver/mysql@v1.5.5
```

### 2、测试gorm链接mysql数据库

为了链接mysql数据库，通常需要导入mysql驱动以及gorm依赖，示例如下：

```go
import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
```

导入完成后，需要mysql地址信息，整个链接格式如下：

```
[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
```

各个参数说明如下：

username：mysql数据的账号名称

password：mysql数据的账号密码

protocol：协议

address：host+port

dbname: 数据库名

param1： 连接的参数1

value1： 连接的参数1对应的值

连接地址的举例： root:123456@tcp(192.168.188.155:3306)/szkfpt?charset=utf8mb4&parseTime=True&loc=Local

完整代码如下：

```go
package chapter01

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
  *  演示通过gorm来链接mysql库
     前提条件：mysql数据库已经安装好了
*/

func GetMysqlDb(account string, password string, host string, port int, dbname string) (*gorm.DB, error) {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	// 链接格式： [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", account, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}
```

测试代码如下：

```go
package test

import (
	"fmt"
	"go-orm-learn/chapter01"
	"testing"
)

func TestMysqlConnection(t *testing.T) {
	db, err := chapter01.GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")
	if err != nil {
		t.Errorf("err is not nil")
	}
	sqldb, _ := db.DB()

	fmt.Println("已经建立的链接数", sqldb.Stats().OpenConnections)
}
```

运行测试代码，结果如下：

```txt
=== RUN   TestMysqlConnection
root:123456@tcp(192.168.188.155:3306)/szkfpt?charset=utf8mb4&parseTime=True&loc=Local
已经建立的链接数 1
--- PASS: TestMysqlConnection (0.01s)
PASS
```

从上面的测试结果来看，目前已经能够正确链接到mysql数据库中了！

### 3、插入数据

##### 准备数据模型

```go
// chapter01/gorm_user.go
package chapter01

import (
	"database/sql"
	"time"
)

type User struct {
	ID           uint           // Standard field for the primary key
	Name         string         // A regular string field
	Email        *string        // A pointer to a string, allowing for null values
	Age          uint8          // An unsigned 8-bit integer
	Birthday     *time.Time     // A pointer to time.Time, can be null
	MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
	CreatedAt    time.Time      // Automatically managed by GORM for creation time
	UpdatedAt    time.Time      // Automatically managed by GORM for update time
}
```



##### 执行插入操作

```go
package chapter01

import (
	"fmt"
	"time"
)

func CreateSingle() {
	db, err := GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")
	if err != nil {

	}
	birthTime := time.Now()
	user := User{Name: "Jinzhu", Age: 18, Birthday: &birthTime}

	// 如果表不存在，则进行自动创建
	db.AutoMigrate(&User{})
	result := db.Create(&user) // 通过数据的指针来创建
	fmt.Println("影响的数据行数为:", result.RowsAffected)
}
```

测试代码：

```go
func TestCreate(t *testing.T) {
	chapter01.CreateSingle()
}
```



##### 插入的数据确认

运行的结果如下：

![image-20240321165520471](E:\go\go-orm-learn\gorm教程.assets\image-20240321165520471.png)

