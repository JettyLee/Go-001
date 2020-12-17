## 作业
我们在数据库操作的时候，比如 `dao` 层中当遇到一个 `sql.ErrNoRows` 的时候，是否应该 `Wrap` 这个 `error`，抛给上层。为什么？应该怎么做请写出代码

## 个人理解
[个人总结](http://jettylee.cn/2020/12/02/go训练营-golang-错误处理/)

## 代码实现：
```go
// DAO 层

type Dao struct { ...... }

var (
	ErrRecordNotFound = errors.New("record not found")
	......
)

func New(......) *Dao {
	return &Dao{ ...... }
}

func (d *Dao) GetById(id int64) (*models.Advertisement, error) {
	item := new(models.Advertisement)
	o := orm.NewOrm()
	err := o.QueryTable(item).Filter("id", id).One(item)
	return item, err
}
```

```go
// Service 层

type Service struct { ...... }

......

func (s *Service) FindAdvertisementByID(id int64) (*model.Advertisement, error) {
    return dao.GetById(id)
}
```

```go
// 接口层

func (c *AdvertisementController) GetAdvertisementById() {
	id, err := c.GetInt64("id")
	if err != nil {
		logs.Error(err)
		c.Data["json"] = utils.Result{StatusCode: utils.Fail, StatusMsg: err.Error()}
		c.ServeJSON()
		return
	}
	advertisementService := services.NewAdvertisementService()

	item, err := advertisementService.FindAdvertisementByID(id)
	if errors.Is(err, dao.ErrRecordNotFound) {
    		item = advertisementService.GetDefaultAd()
    }
	c.Data["json"] = utils.Result{StatusCode: utils.Success, Data: item}
	c.ServeJSON()
	return
}
```