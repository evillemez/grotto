# Mapview #

Some ideas...

Create and use a generic map view.

```go
var ApiView maps.MapView = maps.NewMapView().
  SetExcludedPaths({"some.path"}).
  SetPathHandler("some.path", handlers.MyCustomPathHandler)
  
mapview.ApplyMapView(ApiView, someMap)
```

Customize a map view.

```go
var SearchIndexView = maps.NewMapView().
  Include(ApiView).
  AddRequiredPath("id").
  AddExcludedPath("different.nested.path").
  SetPathHandler("another.path", handlers.SomeHandler).
  SetTypeHandler(typ, handlers.DateHandler).
  SetMaxDepth(3)
```

Create a map view from a structure definition.

```go
// creates a MapView from a Struct definition, uses
// reflection to inspect tags and create a MapView
var WhateverView = maps.NewStructView(models.MyStruct)
```

Change behavior of view at apply time:

```go
var WhateverView = maps.NewStructView(models.MyStruct)
mapped := somelib.StructToMap(myStructInstance)
context := mapview.ViewContext{tags: {"protected","public"}}
mapview.ApplyMapViewContext(WhateverView, mapped, context)
```