# Assert Call

```go
package service

import "testing"
import "errors"
import "app/app/entity"
import "app/app/service"
import mockings "github.com/reisraff/mockings/mockings"
import assert "github.com/stretchr/testify/assert"

type DatabaseWriterMock struct {
}

func (self DatabaseWriterMock) Insert(entity interface{}) {
    mockings.AddCall(&self, "Insert", []interface{}{entity})
}

func (self DatabaseWriterMock) Save(entity interface{}) {
    mockings.AddCall(&self, "Save", []interface{}{entity})
}

func (self DatabaseWriterMock) Remove(entity interface{}) {
    mockings.AddCall(&self, "Remove", []interface{}{entity})

}
type UploadProviderMock struct {
    upload.UploadProvider
}

func (self UploadProviderMock) Validate(v *multipart.FileHeader) error {
    result := mockings.AddCall(&self, "Validate", []interface{}{v})
    return mockings.ErrorOrNil(result[0])
}

func (self UploadProviderMock) GetExtension(v *multipart.FileHeader) string {
    result := mockings.AddCall(&self, "GetExtension", []interface{}{v})
    return result[0].(string)
}

func (self UploadProviderMock) Move(v *multipart.FileHeader, v2 string) error {
    result := mockings.AddCall(&self, "Move", []interface{}{v, v2})
    return mockings.ErrorOrNil(result[0])
}

func (self UploadProviderMock) Remove(v string) {
    mockings.AddCall(&self, "Remove", []interface{}{v})
}

func (self UploadProviderMock) GetUploadDir() string {
    result := mockings.AddCall(&self, "GetUploadDir", []interface{}{})
    return result[0].(string)
}

func GetNotificationService() service.NotificationService {
    notificationService := service.NotificationService{}

    return notificationService
}

func TestCreate(t *testing.T) {
    mockings.ResetAsserts()

    notification := entity.Notification{}

    notificationService := GetNotificationService()
    notificationService.SetDatabaseWriter(&databaseWriterMock)
    notificationService.Create(&notification)

    if notification.GetGuid() == "" {
        t.Fail()
    }

    mockings.AssertCalledWith(t, DatabaseWriterMock{}, "Insert", []interface{}{&notification})
}

func TestUpdateErrorMovingImage(t *testing.T) {
    mockings.Reset()

    mime := make(map[string][]string, 0)

    fileheader := multipart.FileHeader{}
    fileheader.Filename = "bla"
    fileheader.Header = mime
    fileheader.Size = 1

    notification := entity.Notification{}
    notification.SetId(1)
    notification.SetUploadedImage(&fileheader)

    notificationService := GetNotificationService()

    databaseWriterMock := DatabaseWriterMock{}
    notificationService.SetDatabaseWriter(&databaseWriterMock)

    uploadProviderMock := UploadProviderMock{}
    mockings.Mock(&uploadProviderMock, "Validate", []interface{}{&fileheader}, []interface{}{nil})
    mockings.Mock(&uploadProviderMock, "GetExtension", []interface{}{&fileheader}, []interface{}{"jpg"})
    mockings.Mock(&uploadProviderMock, "Move", mockings.ANY, []interface{}{errors.New("Bla")})
    notificationService.SetUploadProvider(&uploadProviderMock)

    assert.Error(t, notificationService.Update(&notification))
}
```
