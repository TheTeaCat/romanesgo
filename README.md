Romanesgo

A simple fractal exploration program I wrote when I was 17/18.
Named after the "cooler" version of broccoli.

Do "romanesgo help" for info on how to use it.

Parameters used to create the example images:
- mandelbrot.png
  - romanesgo -w=2600 -h=2000 -ff=mandelbrot -x=-0.65 -z=0.8 -i=1024 -ss=2
- julia.png
  - romanesgo -w=2600 -h=2000 -ff=julia -c=-0.2 -c=0.65 -z=0.9 -i=512 -ss=2
- burningship.png
  - romanesgo -w=2000 -h=2600 -ff=burningship -x=-1.749 -y=0.037 -z=20 -i=256 -ss=2
- mandelbrot2.png
  - romanesgo -w=2600 -h=2000 -ff=mandelbrot -x=-0.82 -y=-0.1905 -z=50 -i=512 -ss=2 -cf=smoothgreyscale
- julia2.png
  - romanesgo -w=2600 -h=2000 -ff=julia -c=-0.2 -c=0.65 -z=5 -i=512 -ss=2 -cf=whackygreyscale
- julia3.png
  - romanesgo -w=2600 -h=2000 -ff=julia -c=-1 -c=-0.25 -z=1.5 -i=512 -ss=2 -cf=zgreyscale
- julia4.png
  - romanesgo -w=2000 -h=2600 -ff=julia -c=0.1 -c=0.7 -z=0.75 -ss=2 -cf=smoothcolour