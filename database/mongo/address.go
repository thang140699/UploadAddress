package mongo

import (
	"WeddingUtilities/model"
	"WeddingUtilities/utilities/provider/mongo"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const AddressMongoCollection = "Address"

type AddressMongoRepository struct {
	provider       *mongo.MongoProvider
	collectionName string
}

func SaveAddress(provider *mongo.MongoProvider, repo *AddressMongoRepository) error {

	collection, close := repo.collection()
	defer close()
	// ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	var docs []interface{}
	var dfg []model.Address
	err := collection.Find(bson.M{"order": 10891}).One(&dfg)
	if err != nil {
		fmt.Println(err)
		err = Insert(docs, collection)
		if err != nil {
			return err
		}
		return nil
	}
	fmt.Println(docs)
	return nil
}

func Insert(docs []interface{}, collection *mgo.Collection) error {
	byteValues, err := ioutil.ReadFile("/home/thang/Wedding-event/Wedding_Utilities/utilities/provider/mongo/new_address_value.json")
	if err != nil {
		fmt.Println("ioutil.ReadFile Error", err)
		return err
	}

	err = json.Unmarshal(byteValues, &docs)

	for i := range docs {
		doc := docs[i]
		// fmt.Println(doc)
		insertErr := collection.Insert(doc)
		if insertErr != nil {
			fmt.Println("InsertOne ERROR:", insertErr)
			return err
		}
	}
	return nil
}
func NewAddressMongoRepository(provider *mongo.MongoProvider) *AddressMongoRepository {
	repo := &AddressMongoRepository{provider, AddressMongoCollection}
	// collection, close := repo.collection()
	// defer close()

	// collection.EnsureIndex(mgo.Index{
	// 	Key: []string{
	// 		"email",
	// 	},
	// 	Unique: true,
	// })

	// collection.EnsureIndex(mgo.Index{
	// 	Key: []string{
	// 		"userID",
	// 	},
	// 	Unique: true,
	// })
	err := SaveAddress(provider, repo)
	if err != nil {
		fmt.Println(err)
		return &AddressMongoRepository{}

	}
	return repo
}

func (repo *AddressMongoRepository) collection() (collection *mgo.Collection, close func()) {
	session := repo.provider.MongoClient().GetCopySession()
	close = session.Close

	return session.DB(repo.provider.MongoClient().Database()).C(repo.collectionName), close
}
func (repo *AddressMongoRepository) All() ([]model.Address, error) {
	collection, close := repo.collection()
	defer close()

	result := make([]model.Address, 0)
	err := collection.Find(nil).All(&result)
	return result, repo.provider.NewError(err)
}

func (repo *AddressMongoRepository) FindByID(Id string) (*model.Address, error) {
	collection, close := repo.collection()
	defer close()
	if !bson.IsObjectIdHex(Id) {
		return nil, fmt.Errorf("invalid Id")
	}

	var address model.Address
	err := collection.Find(bson.ObjectIdHex(Id)).One(&address)
	return &address, repo.provider.NewError(err)
}
func (repo *AddressMongoRepository) FindByCodeName(CodeName string) (*model.Address, error) {
	collection, close := repo.collection()
	defer close()
	var address model.Address
	err := collection.Find(bson.M{"CodeName": CodeName}).One(&address)
	return &address, repo.provider.NewError(err)
}
func (repo *AddressMongoRepository) FindByName(name string) (*model.Address, error) {
	collection, close := repo.collection()
	defer close()
	var address model.Address
	err := collection.Find(bson.M{"Name": name}).One(&address)
	return &address, repo.provider.NewError(err)
}
func (repo *AddressMongoRepository) FindByDivisionType(Divisiontype string) (*model.Address, error) {
	collection, close := repo.collection()
	defer close()
	var address model.Address
	err := collection.Find(bson.M{"DivisionType": Divisiontype}).One(&address)
	return &address, repo.provider.NewError(err)
}
func (repo *AddressMongoRepository) FindByPhoneCode(phoneCode string) (*model.Address, error) {
	collection, close := repo.collection()
	defer close()
	var address model.Address
	err := collection.Find(bson.M{"phoneCode": phoneCode}).One(&address)
	return &address, repo.provider.NewError(err)
}
func (repo *AddressMongoRepository) FindByParentId(ParentId string) (*model.Address, error) {
	collection, close := repo.collection()
	defer close()
	var address model.Address
	err := collection.Find(bson.M{"ParentID": ParentId}).One(&address)
	return &address, repo.provider.NewError(err)
}
func (repo *AddressMongoRepository) FindByLevel(level int) (*model.Address, error) {
	collection, close := repo.collection()
	defer close()
	var address model.Address
	err := collection.Find(bson.M{"Level": level}).One(&address)
	return &address, repo.provider.NewError(err)
}

func (repo *AddressMongoRepository) Save(address model.Address) error {
	collection, close := repo.collection()
	defer close()

	err := collection.Insert(address)
	return repo.provider.NewError(err)

}
