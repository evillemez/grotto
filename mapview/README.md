# Mapview #

Some ideas...

```go
var ApiView maps.MapView = maps.NewMapView().
  SetExcludedPaths({"some.path"}).
  SetPathHandler("some.path", handlers.MyCustomPathHandler)
```

```go
var SearchIndexView = maps.NewMapView().
  Include(ApiView).
  AddRequiredPath("id").
  AddExcludedPath("different.nested.path").
  SetPathHandler("another.path", handlers.SomeHandler)
```
```go
// creates a MapView from a Struct definition, uses
// reflection to inspect tags and create a MapView
var WhateverView = maps.NewStructView(models.MyStruct)
```