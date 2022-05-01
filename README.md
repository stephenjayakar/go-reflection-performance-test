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

# Results

## Summary for strings

Slices only behaved better on small N, and only slightly. When we were at `N = 100`, we already are at a 30% faster using `MapSet`. Looks like for most string cases, using a `map` as the underlying storage makes sense.

## N = 100

```
BenchmarkSliceSet-10       	   51308	     22884 ns/op
BenchmarkSliceSetExp-10    	   54638	     22027 ns/op
BenchmarkMapSet-10         	   79706	     14912 ns/op
```

## N = 1000

```
BenchmarkSliceSet-10       	     741	   1504577 ns/op
BenchmarkSliceSetExp-10    	     793	   1535511 ns/op
BenchmarkMapSet-10         	    6694	    180629 ns/op
```

## N = 10

```
BenchmarkSliceSet-10       	 1110586	      1066 ns/op
BenchmarkSliceSetExp-10    	 1000000	      1052 ns/op
BenchmarkMapSet-10         	  987474	      1222 ns/op
```

## Open questions

Wonder if it's different for ints?
