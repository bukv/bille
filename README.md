# Simple billiards in Golang
A small game using the [fogleman/gg](http://github.com/fogleman/gg "github.com/fogleman/gg") library.
The current version communicates via command-line interface.

***
### How to play
1. In the terminal, open the directory where this project is located. Run this project:
> $ go run bille
2. Now you can see the starting position of the ball. Open the directory `out` and open the image `out.png`:
![start_position.png](https://raw.githubusercontent.com/bukv/bille/develop/images/start_position.png "Start position")
3. Return to your terminal. The game asks you to specify the force of impact (from 1 to 255) and then the direction (from 0 to 360 degrees). Indicate the values that you think are appropriate.
![degrees.png](https://raw.githubusercontent.com/bukv/bille/develop/images/degrees.png "This is how the degrees are calculated")
*This is how the degrees are calculated.*
4. Now return to the image `out.png` from `out` directory. You can see the trajectory of the ball:
![final_1.png](https://raw.githubusercontent.com/bukv/bille/develop/images/final_1.png "Final_1")
If you get in the hole, the effect is as follows:
![final_2.png](https://raw.githubusercontent.com/bukv/bille/develop/images/final_2.png "Final_2")