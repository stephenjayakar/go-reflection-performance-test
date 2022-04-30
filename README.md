# Thesis

The accepted paradigm for doing set contains (or set operations in general) pre Go 1.18 (generics) is to use a `map[type]struct{}`. I believe that using this for small datasets is going to be slower than 
* cache coherency

Make sure to allow arbitrary elements inside the set?
* Can I even do this for the map example?

Set contains
* Using slices.Contains
* Using map
* Using interface{} + reflect for a 

Set equals
* Using map
* Using sort + slices.Equal
* using reflect.DeepEqual
