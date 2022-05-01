# Thesis

The accepted paradigm for doing set contains (or set operations in general) pre Go 1.18 (generics) is to use a `map[type]struct{}`. I believe that using this for small datasets is going to be slower than using slices. I also just don't like the way `map[string]struct{}{}` looks. Disgusting.

# String Results

## Summary

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

# Int results

## Summary

Ints perform better as slices for longer. At `N = 100`, we still have it at about 54% faster. However, the gap really widens once we get to `N = 1000`.

## N = 100
```
BenchmarkIntMapSet-10      	  197214	      5917 ns/op
BenchmarkIntSliceSet-10    	  435561	      2709 ns/op
```

## N = 10

```
BenchmarkIntMapSet-10      	 3872775	       311.5 ns/op
BenchmarkIntSliceSet-10    	 8202166	       144.7 ns/op
```

## N = 1000

```
BenchmarkIntMapSet-10      	   20090	     60114 ns/op
BenchmarkIntSliceSet-10    	    6854	    175596 ns/op
```
