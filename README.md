# Assert Call

```go
package service

import "testing"
import "app/app/entity"
import "app/app/service"
import assertcall "github.com/reisraff/assertcall/assert"

type DatabaseWriterMock struct {
}

func (self DatabaseWriterMock) Insert(entity interface{}) {
    assertcall.AddAssert(self, "Insert", []interface{}{entity})
}

func (self DatabaseWriterMock) Save(entity interface{}) {
    assertcall.AddAssert(self, "Save", []interface{}{entity})
}

func (self DatabaseWriterMock) Remove(entity interface{}) {
    assertcall.AddAssert(self, "Save", []interface{}{entity})
}

func GetNotificationService() service.NotificationService {
    notificationService := service.NotificationService{}
    notificationService.SetDatabaseWriter(DatabaseWriterMock{})

    return notificationService
}

func TestCreate(t *testing.T) {
    assertcall.ResetAsserts()

    notification := entity.Notification{}

    notificationService := GetNotificationService()
    notificationService.Create(&notification)

    if notification.GetGuid() == "" {
        t.Fail()
    }

    assertcall.AssertCalledWith(t, DatabaseWriterMock{}, "Insert", []interface{}{&notification})
}
```
