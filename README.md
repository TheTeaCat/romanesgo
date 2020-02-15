# Romanesgo

A simple fractal exploration program I wrote when I was 17/18.
Named after the "cooler" version of broccoli.



## Usage

``` 
$ ./romanesgo help

Do "romanesco help {Fractal Name}" for further info on a particular fractal function.

Fractals:
         burningship
         julia
         mandelbrot

Flags:
  -c value
        constants
  -cf string
        colouring function (default "default")
  -ff string
        fractal (default "none")
  -fn string
        filename (default "temp.png")
  -h int
        image height (default 1000)
  -i int
        maximum iterations (default 128)
  -r int
        goroutines used (default 8)
  -ss int
        supersampling factor (default 1)
  -w int
        image width (default 1000)
  -x float
        central x coord
  -y float
        central y coord
  -z float
        zoom factor (default 1)
```



## Performance

So, here's some usage on an i5-3320m (pretty old lil laptop processor):

```
$ ./romanesgo -ff=burningship -x=-1.748 -y=0.035 -z=20 -ss=4 -w=25000 -h=25000 -fn=bigship.png

	Fractal (ff):		 burningship 
	Constants (c):		  
	Max Iterations (i):	 128 
	Colouring function (cf): default 
	Centre x Coord (x):	 -1.748 
	Centre y Coord (y):	 0.035 
	Zoom factor (z):	 20 
	Image Width (w):	 25000 
	Image Height (h):	 25000 
	Supersampling (ss):	 4 
	Routines (r):		 4 
	Filename (png) (fn):	 bigship.png 

Routine 2 Done.
Routine 0 Done.
Routine 3 Done.
Routine 1 Done.

Time taken: 14m9.110651581s
```

What is that, a 625 megapixel image, in 14 minutes? I guess that's alright for an old laptop. :man_shrugging:



## Example images

| mandelbrot.png                                               | julia.png                                                    |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| -w=2600 -h=2000 -ff=mandelbrot -x=-0.65 -z=0.8 -i=1024 -ss=2 | -w=2600 -h=2000 -ff=julia -c=-0.2 -c=0.65 -z=0.9 -i=512 -ss=2 |
| ![mandelbrot.png](./example%20images/mandelbrot.png)         | ![julia.png](./example%20images/julia.png)                   |

| burningship.png                                              | julia4.png                                                   |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| -w=2000 -h=2600 -ff=burningship -x=-1.749 -y=0.037 -z=20 -i=256 -ss=2 | -w=2000 -h=2600 -ff=julia -c=0.1 -c=0.7 -z=0.75 -ss=2 -cf=smoothcolour |
| ![burningship.png](./example%20images/burningship.png)       | ![julia4.png](./example%20images/julia4.png)                 |


| mandelbrot2.png                                              | julia2.png                                                   |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| -w=2600 -h=2000 -ff=mandelbrot -x=-0.82 -y=-0.1905 -z=50 -i=512 -ss=2 -cf=smoothgreyscale | -w=2600 -h=2000 -ff=julia -c=-0.2 -c=0.65 -z=5 -i=512 -ss=2 -cf=whackygreyscale |
| ![mandelbrot2.png](./example%20images/mandelbrot2.png)       | ![julia2.png](./example%20images/julia2.png)                 |

| julia3.png                                                   |
| ------------------------------------------------------------ |
| -w=2600 -h=2000 -ff=julia -c=-1 -c=-0.25 -z=1.5 -i=512 -ss=2 -cf=zgreyscale |
| ![julia3.png](./example%20images/julia3.png)                 |



  
