package objectstore

import (
	"fmt"
	"github.com/banzaicloud/banzai-types/components"
	"github.com/banzaicloud/banzai-types/constants"
	"github.com/banzaicloud/pipeline/config"
	"github.com/banzaicloud/pipeline/secret"
	"github.com/sirupsen/logrus"
	"github.com/banzaicloud/pipeline/auth"
	"errors"
	"reflect"
	"github.com/banzaicloud/pipeline/model"
)

var logger *logrus.Logger

func init() {
	logger = config.Logger()
}

var ManagedBucketNotFoundError 			= errors.New("Managed bucket not found")
var MultipleManagedBucketsFoundError = errors.New("Multiple managed buckets found")

type CommonObjectStore interface {
	CreateBucket(string) error
	ListBuckets() error
	DeleteBucket(string) error
}

type ManagedBucketsStore interface {
	GetManagedBuckets(string) (interface{}, error)
}

func ListCommonObjectStoreBuckets(s *secret.SecretsItemResponse) (CommonObjectStore, error) {
	switch s.SecretType {
	case constants.Amazon:
		return nil, nil
	case constants.Google:
		return nil, nil
	case constants.Azure:
		return nil, nil
	default:
		return nil, fmt.Errorf("listing a bucket is not supported for %s", s.SecretType)
	}
}

func CreateCommonObjectStoreBuckets(createBucketRequest components.CreateBucketRequest, s *secret.SecretsItemResponse, user *auth.User) (CommonObjectStore, error) {
	switch s.SecretType {
	case constants.Amazon:
		return &AmazonObjectStore{
				region: createBucketRequest.Properties.CreateAmazonObjectStoreBucketProperties.Location,
				secret: s,
				user: user,
			}, nil
	case constants.Google:
		return &GoogleObjectStore{
				location:       createBucketRequest.Properties.CreateGoogleObjectStoreBucketProperties.Location,
				serviceAccount: NewGoogleServiceAccount(s),
				user:           user,
			}, nil
	case constants.Azure:
		return &AzureObjectStore{
				storageAccount: createBucketRequest.Properties.CreateAzureObjectStoreBucketProperties.StorageAccount,
				resourceGroup:  createBucketRequest.Properties.CreateAzureObjectStoreBucketProperties.ResourceGroup,
				location:       createBucketRequest.Properties.CreateAzureObjectStoreBucketProperties.Location,
				secret:         s,
				user:           user,
			}, nil
	default:
		return nil, fmt.Errorf("creating a bucket is not supported for %s", s)
	}
}

func NewGoogleObjectStore(s *secret.SecretsItemResponse, user *auth.User) (CommonObjectStore, error) {
	return &GoogleObjectStore{
		serviceAccount: NewGoogleServiceAccount(s),
		user: user,
	}, nil
}

func NewAmazonObjectStore(s *secret.SecretsItemResponse, user *auth.User, region string) (CommonObjectStore, error) {
	return &AmazonObjectStore{
		secret: s,
		region: region,
		user: 	user,
	}, nil
}

func NewAzureObjectStore(s *secret.SecretsItemResponse, user *auth.User, resourceGroup, storageAccount string) (CommonObjectStore, error) {
	return &AzureObjectStore{
		storageAccount: storageAccount,
		resourceGroup:  resourceGroup,
		secret:         s,
		user: 					user,
	}, nil
}



func GetValidatedManagedBucket(bucketName string, managedBucketStore ManagedBucketsStore) (interface{}, error) {
	managedBuckets, err := managedBucketStore.GetManagedBuckets(bucketName)
	if err != nil {
		return nil, err
	}

	managedBucketsTyped := reflect.ValueOf(managedBuckets)

	if managedBucketsTyped.Len() == 0 {
		return nil, ManagedBucketNotFoundError
	}

	if managedBucketsTyped.Len() > 1 {
		return nil, MultipleManagedBucketsFoundError
	}

	return managedBucketsTyped.Index(0), nil
}

func persistToDb(m interface{}) error {
	log := logger.WithFields(logrus.Fields{"tag": "persistToDb"})
	log.Info("Persisting Bucket Description to Db")
	db := model.GetDB()
	return db.Save(m).Error
}

func deleteFromDbByPK(deleteCriteria interface{}) error {
	log := logger.WithFields(logrus.Fields{"tag": "deleteFromDbByPK"})
	log.Info("Deleting from DB...")
	db := model.GetDB()
	return db.Delete(deleteCriteria).Error
}

func deleteFromDb(deleteCriteria interface{}) error {
	log := logger.WithFields(logrus.Fields{"tag": "deleteFromDb"})
	log.Info("Deleting from DB...")
	db := model.GetDB()
	return db.Delete(deleteCriteria, deleteCriteria).Error
}
