# DatabaseSqlNullString

## Properties

 Name       | Type                  | Description                         | Notes      
------------|-----------------------|-------------------------------------|------------
 **String** | Pointer to **string** |                                     | [optional] 
 **Valid**  | Pointer to **bool**   | Valid is true if String is not NULL | [optional] 

## Methods

### NewDatabaseSqlNullString

`func NewDatabaseSqlNullString() *DatabaseSqlNullString`

NewDatabaseSqlNullString instantiates a new DatabaseSqlNullString object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseSqlNullStringWithDefaults

`func NewDatabaseSqlNullStringWithDefaults() *DatabaseSqlNullString`

NewDatabaseSqlNullStringWithDefaults instantiates a new DatabaseSqlNullString object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetString

`func (o *DatabaseSqlNullString) GetString() string`

GetString returns the String field if non-nil, zero value otherwise.

### GetStringOk

`func (o *DatabaseSqlNullString) GetStringOk() (*string, bool)`

GetStringOk returns a tuple with the String field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetString

`func (o *DatabaseSqlNullString) SetString(v string)`

SetString sets String field to given value.

### HasString

`func (o *DatabaseSqlNullString) HasString() bool`

HasString returns a boolean if a field has been set.

### GetValid

`func (o *DatabaseSqlNullString) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *DatabaseSqlNullString) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *DatabaseSqlNullString) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *DatabaseSqlNullString) HasValid() bool`

HasValid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


