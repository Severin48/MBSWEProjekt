Wahlpflichtfach Modellbasierte Softwareentwicklung

Severin Hotz

Betreut durch Prof. Dr. Martin Sulzmann

# Topic 1: Translating methods, interfaces and structural subtyping
Description of the tasks:
https://sulzmann.github.io/ModelBasedSW/projectSoSe23.html#(3)

## Task 1: Compare the run-time performance of RT and DT (e.g. call “sumArea” in a loop and measure which version runs faster)

The first task was to expand the given example to measure the run-time differences between "run-time method look up" (RT) and "dictionary-translation" (DT).
RT and DT are methods of dealing with polymorphism in object oriented programming.
RT describes the process of looking up the code to be executed in response to a method call at runtime by using the values to carry the type information.
The type of the values is queried at run-time to select the appropriate method.
In DT, the interface values are represented as a pair of actual value and the dictionary of method definitions. It tries to improve performance by resolving method calls at compile time.
Both methods rely on Go's interface{}.

In the file Task1.go the existing code from
https://sulzmann.github.io/ModelBasedSW/lec-go-2-types-methods-interfaces.html#(7)
was used. The functions generateShapes(amount int) and Task1() were added to create a large amount of sample shapes and later compare the speed of RT and DT.
In this example, 10000000 shapes are created with random properties with a ratio of about 50% rectangles and 50% squares.
During this creation process, two arrays are filled with objects - shapes[] and shape_values[]. The former is used for run-time method look up while the latter is used for dictionary-translation.
The RT array (shapes) only needs to contain the objects (rectangle or square), while the DT array gets the objects and the area function.

In the function Task1(), the function generateShapes is called and the two arrays of shapes and shape_values are used to measure the times of sumArea_Lookup and sumArea_Dict respectively.

When executed, the resulting times are printed to the console. Typically, the DT approach is slightly slower than the RT method look up by a few milliseconds. This result is unexpected but could possibly be explained with certain compiler optimizations.

## Task 2: Apply the RT and DT approach to one further example of your own choice
In Task 2, the same code was used with the addition of the interface body (instead of shape) as well as the structs cube and sphere. The time measurement results in this experiment are almost identical to the measurements of Task 1 even though the calculations are slightly more complex.

## Task 3: Extend RT and DT to deal with type assertions
The goal of task 3 was to deal with type assertions in both the RT and the DT approach. This included the check whether the current object is a square and adjusting the calculation method as shown in the example when it is. The calculation was to be carried out as before if the shape was not a square. This was done by utilizing Go's run-time type checking capabilities.

Implementing this increased the run-times of RT and DT respectively by around 10 milliseconds each for 10000000 shapes.

## Task 4: Extend RT and DT to deal with type bounds
For the implementation of type bounds into the RT and DT approaches, another example was given with the topic description. In the example, different nodes are being created and shown. Their correct representation is printed to the console, depending on which type of node it is.

In this task, the time is measured as well after having created random nodes. The results of this experiment shows the DT method as slightly faster than RT.