This repo does nothing useful. These are just first experiments with both `otto` and `golang`.

TODO: 

* Next step, use Negroni to add some middleware like logging
  * https://github.com/codegangsta/negroni
* serve some static templates
* serve some static assets (css)
* create some structs for models
* persist models to db:
  * mongo w/ mgo: https://github.com/go-mgo/mgo/tree/v2
  * postgres w/ ...
    * https://github.com/go-xorm/xorm
    * https://github.com/jinzhu/gorm
    * https://github.com/go-gorp/gorp
* expose data manipulation via api
  * GETs
    * pull data from db, resulting in struct instances
    * convert struct to map: https://github.com/fatih/structs
    * transform/sanitize map
    * marshall map into json
  * PUTs
    * unmarshall raw json into map
    * transform/sanitize map
    * convert map to struct https://github.com/mitchellh/mapstructure
    * persist structs to db
* validation on maps maybe, or structs: https://github.com/asaskevich/govalidator
* write unit tests w/ help from: https://github.com/stretchr/testify
* integration tests w/ help from: https://golang.org/pkg/net/http/httptest/

* contemplate partial model updates & partial views
  * write a map transformer library, end result would be similar to JMS Serializer Groups, but applying transformation on intermediary map representation, not structs directly
  * call it mapblaster? :)  Or more generic: mapview
  * theoretical example usage (don't know how possible it is yet):
    * mutate map by whitelisting/blacklisting object paths
      * transformedMap, err := maps.IncludePaths(someMap, {"some.path","another.path"})
      * transformedMap, err := maps.ExcludePaths(someMap, {"some.path","another.path"})
    * create and apply configurable views to a map
      * creating a view
        * var ApiView maps.MapView = maps.NewMapView().
          SetExcludedPaths({"some.path"}).
          SetPathHandler("some.path", handlers.MyCustomPathHandler)
        * var SearchIndexView = maps.NewMapView().
          Include(ApiView).
          AddRequiredPath("id").
          AddExcludedPath("different.nested.path").
          SetPathHandler("another.path", handlers.SomeHandler)
      * using a view
        * transformedMap, err := maps.ApplyMapView(someMap, ApiView)
  
