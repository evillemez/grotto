This repo does nothing useful. These are just first experiments with both `otto` and `golang`.

TODO: 

* Next step, use Negroni to add some middleware like logging
* serve some static templates
* serve some static assets (css)
* create some models
* expose via api
  * GET
    * convert struct to map: https://github.com/fatih/structs
    * transform/sanitize map
    * marshall map into json
  * PUT
    * unmarshall raw json into map
    * transform/sanitize map
    * convert map to struct https://github.com/mitchellh/mapstructure
* persist models to mongo w/ mgo
* contemplate partial model updates & partial views