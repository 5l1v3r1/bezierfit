# Abstract

With CSS3 animations, you can write Bezier curves that act as animation timing functions. Bezier curves are generally nice and easy to work with, but it is sometimes difficult to create a Bezier curve that looks just like you want it.

Enter **bezierfit**. With **bezierfit**, you can produce a Bezier curve animation which matches a set of key points *as best as possible* (in a mathematical least-squares sense). For instance, suppose you want an animation to go through completion percentages 30%, 50%, 80%, and 90% at time values 0.1, 0.2, 0.5, and 0.8 respectively. With this information, **bezierfit** can search for the Bezier animation which *best matches* the given parameters.

# How it works

Solving a least-squares problem like this is an optimization problem (but not necessarily a convex one). There are several numerical tricks behind **bezierfit**:

 * Bezier animations are evaluated at a given X value using bisection search.
 * Gradients of the loss function with respect to Bezier curve parameters are computed using two-point numerical differentiation.
 * Locally optimal bezier animations are discovered through an active set optimization technique based on gradient descent (with a fixed step size).
 * Good local minima are discovered by solving multiple optimization problems from random starting points.

This non-convex optimization problem is relatively easy for a number of reasons. First, there are only four parameters, so even inefficient optimization techniques are relatively quick. Second, all parameters will tend to be around the same order of magnitude (i.e. bezier parameters are all relatively close to the range [0,1]). Third, not too much precision is needed, so using tons of numerical approximations is acceptable.
