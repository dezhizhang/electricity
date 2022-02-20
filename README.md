# 电商系统

### 创建用户

```go
for i := 0; i < 10; i++ {
    birthDay := time.Unix(int64(1606832965), 0)
    user := model.User{
    NickName: fmt.Sprintf("刘德华%d", i),
    Phone: fmt.Sprintf("1599247844%d", i),
    Gender: fmt.Sprintf("男"),
    Birthday:&birthDay,
    Password: utils.Md5String(string("123456")),
    
    }
    user.BaseModel.UpdateAt = birthDay
    driver.DB.Save(&user)
}
```